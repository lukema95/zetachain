package staking

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/zeta-chain/node/precompiles/bank"
	precompiletypes "github.com/zeta-chain/node/precompiles/types"
	fungibletypes "github.com/zeta-chain/node/x/fungible/types"
)

// function distribute(address zrc20, uint256 amount) external returns (bool success)
func (c *Contract) distribute(
	ctx sdk.Context,
	evm *vm.EVM,
	contract *vm.Contract,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	if len(args) != 2 {
		return nil, &precompiletypes.ErrInvalidNumberOfArgs{
			Got:    len(args),
			Expect: 2,
		}
	}

	// Unpack arguments and check if they are valid.
	zrc20Addr, amount, err := unpackDistributeArgs(args)
	if err != nil {
		return nil, err
	}

	// Get the original caller address. Necessary for LockZRC20 to work.
	caller, err := precompiletypes.GetEVMCallerAddress(evm, contract)
	if err != nil {
		return nil, err
	}

	// Create the coinSet in advance, if this step fails do not lock ZRC20.
	coinSet, err := precompiletypes.CreateZRC20CoinSet(zrc20Addr, amount)
	if err != nil {
		return nil, err
	}

	// LockZRC20 locks the ZRC20 under the locker address.
	// It performs all the necessary checks such as allowance in order to execute a transferFrom.
	// - spender is the staking contract address (c.Address()).
	// - owner is the caller address.
	// - locker is the bank address. Assets are locked under this address to prevent liquidity fragmentation.
	if err := c.fungibleKeeper.LockZRC20(ctx, zrc20Addr, c.Address(), caller, bank.ContractAddress, amount); err != nil {
		return nil, &precompiletypes.ErrUnexpected{
			When: "LockZRC20InBank",
			Got:  err.Error(),
		}
	}

	// With the ZRC20 locked, proceed to mint the cosmos coins.
	if err := c.bankKeeper.MintCoins(ctx, fungibletypes.ModuleName, coinSet); err != nil {
		return nil, &precompiletypes.ErrUnexpected{
			When: "MintCoins",
			Got:  err.Error(),
		}
	}

	// Send the coins to the FeePool.
	if err := c.bankKeeper.SendCoinsFromModuleToModule(ctx, fungibletypes.ModuleName, authtypes.FeeCollectorName, coinSet); err != nil {
		return nil, &precompiletypes.ErrUnexpected{
			When: "SendCoinsFromModuleToModule",
			Got:  err.Error(),
		}
	}

	// Log similar message as in abci DistributeValidatorRewards function.
	ctx.Logger().Info(
		"Distributing ZRC20 Validator Rewards",
		"Total", amount.String(),
		"Fee_collector", authtypes.FeeCollectorName,
		"Denom", precompiletypes.ZRC20ToCosmosDenom(zrc20Addr),
	)

	if err := c.addDistributeLog(ctx, evm.StateDB, caller, zrc20Addr, amount); err != nil {
		return nil, &precompiletypes.ErrUnexpected{
			When: "AddDistributeLog",
			Got:  err.Error(),
		}
	}

	return method.Outputs.Pack(true)
}

func unpackDistributeArgs(args []interface{}) (zrc20Addr common.Address, amount *big.Int, err error) {
	zrc20Addr, ok := args[0].(common.Address)
	if !ok {
		return common.Address{}, nil, &precompiletypes.ErrInvalidAddr{
			Got: zrc20Addr.String(),
		}
	}

	amount, ok = args[1].(*big.Int)
	if !ok || amount == nil || amount.Sign() <= 0 {
		return common.Address{}, nil, &precompiletypes.ErrInvalidAmount{
			Got: amount.String(),
		}
	}

	return zrc20Addr, amount, nil
}
