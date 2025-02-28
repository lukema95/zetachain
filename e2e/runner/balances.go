package runner

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
)

// AccountBalances is a struct that contains the balances of the accounts used in the E2E test
type AccountBalances struct {
	ZetaETH   *big.Int
	ZetaZETA  *big.Int
	ZetaWZETA *big.Int
	ZetaERC20 *big.Int
	ZetaBTC   *big.Int
	ZetaSOL   *big.Int
	ZetaSPL   *big.Int
	EvmETH    *big.Int
	EvmZETA   *big.Int
	EvmERC20  *big.Int
	BtcBTC    string
	SolSOL    *big.Int
	SolSPL    *big.Int
}

// AccountBalancesDiff is a struct that contains the difference in the balances of the accounts used in the E2E test
type AccountBalancesDiff struct {
	ETH   *big.Int
	ZETA  *big.Int
	ERC20 *big.Int
}

type ERC20BalanceOf interface {
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

func (r *E2ERunner) getERC20BalanceSafe(z ERC20BalanceOf, name string) *big.Int {
	// have to use reflect to check nil interface because go'ism
	if z == nil || reflect.ValueOf(z).IsNil() {
		r.Logger.Print("❓ balance of %s: nil", name)
		return new(big.Int)
	}
	res, err := z.BalanceOf(&bind.CallOpts{}, r.EVMAddress())
	if err != nil {
		r.Logger.Print("❓ balance of %s: %v", name, err)
		return new(big.Int)
	}
	return res
}

// GetAccountBalances returns the account balances of the accounts used in the E2E test
func (r *E2ERunner) GetAccountBalances(skipBTC bool) (AccountBalances, error) {
	// zevm
	zetaZeta, err := r.ZEVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
	if err != nil {
		return AccountBalances{}, fmt.Errorf("get zeta balance: %w", err)
	}
	zetaWZeta, err := r.WZeta.BalanceOf(&bind.CallOpts{}, r.EVMAddress())
	if err != nil {
		return AccountBalances{}, fmt.Errorf("get wzeta balance: %w", err)
	}
	zetaEth := r.getERC20BalanceSafe(r.ETHZRC20, "eth zrc20")
	zetaErc20 := r.getERC20BalanceSafe(r.ERC20ZRC20, "erc20 zrc20")
	zetaBtc := r.getERC20BalanceSafe(r.BTCZRC20, "btc zrc20")
	zetaSol := r.getERC20BalanceSafe(r.SOLZRC20, "sol zrc20")
	zetaSPL := r.getERC20BalanceSafe(r.SPLZRC20, "spl zrc20")

	// evm
	evmEth, err := r.EVMClient.BalanceAt(r.Ctx, r.EVMAddress(), nil)
	if err != nil {
		return AccountBalances{}, fmt.Errorf("get eth balance: %w", err)
	}
	evmZeta := r.getERC20BalanceSafe(r.ZetaEth, "zeta eth")
	evmErc20 := r.getERC20BalanceSafe(r.ERC20, "eth erc20")

	// bitcoin
	var BtcBTC string
	if !skipBTC {
		if BtcBTC, err = r.GetBitcoinBalance(); err != nil {
			return AccountBalances{}, err
		}
	}

	// solana
	var solSOL *big.Int
	var solSPL *big.Int
	if r.Account.SolanaAddress != "" && r.Account.SolanaPrivateKey != "" && r.SolanaClient != nil {
		solanaAddr := solana.MustPublicKeyFromBase58(r.Account.SolanaAddress.String())
		privateKey := solana.MustPrivateKeyFromBase58(r.Account.SolanaPrivateKey.String())
		solSOLBalance, err := r.SolanaClient.GetBalance(
			r.Ctx,
			solanaAddr,
			rpc.CommitmentConfirmed,
		)
		if err != nil {
			return AccountBalances{}, fmt.Errorf("get sol balance: %w", err)
		}

		// #nosec G115 always in range
		solSOL = big.NewInt(int64(solSOLBalance.Value))

		if r.SPLAddr != (solana.PublicKey{}) {
			ata := r.ResolveSolanaATA(
				privateKey,
				solanaAddr,
				r.SPLAddr,
			)
			splBalance, err := r.SolanaClient.GetTokenAccountBalance(r.Ctx, ata, rpc.CommitmentConfirmed)
			if err != nil {
				return AccountBalances{}, fmt.Errorf("get spl balance: %w", err)
			}

			solSPLParsed, ok := new(big.Int).SetString(splBalance.Value.Amount, 10)
			if !ok {
				return AccountBalances{}, errors.New("can't parse spl balance")
			}

			solSPL = solSPLParsed
		}
	}

	return AccountBalances{
		ZetaETH:   zetaEth,
		ZetaZETA:  zetaZeta,
		ZetaWZETA: zetaWZeta,
		ZetaERC20: zetaErc20,
		ZetaBTC:   zetaBtc,
		ZetaSOL:   zetaSol,
		ZetaSPL:   zetaSPL,
		EvmETH:    evmEth,
		EvmZETA:   evmZeta,
		EvmERC20:  evmErc20,
		BtcBTC:    BtcBTC,
		SolSOL:    solSOL,
		SolSPL:    solSPL,
	}, nil
}

// GetBitcoinBalance returns the spendable BTC balance of the BTC address
func (r *E2ERunner) GetBitcoinBalance() (string, error) {
	address, _ := r.GetBtcAddress()
	total, err := r.GetBitcoinBalanceByAddress(address)
	if err != nil {
		return "", err
	}

	return total.String(), nil
}

// GetBitcoinBalanceByAddress get btc balance by address.
func (r *E2ERunner) GetBitcoinBalanceByAddress(address btcutil.Address) (btcutil.Amount, error) {
	unspentList, err := r.BtcRPCClient.ListUnspentMinMaxAddresses(r.Ctx, 1, 9999999, []btcutil.Address{address})
	if err != nil {
		return 0, errors.Wrap(err, "failed to list unspent")
	}

	var total btcutil.Amount
	for _, unspent := range unspentList {
		if unspent.Spendable {
			total += btcutil.Amount(unspent.Amount * 1e8)
		}
	}

	return total, nil
}

// PrintAccountBalances shows the account balances of the accounts used in the E2E test
// Note: USDT is mentioned as erc20 here because we want to show the balance of any erc20 contract
func (r *E2ERunner) PrintAccountBalances(balances AccountBalances) {
	r.Logger.Print(" ---💰 Account info %s ---", r.EVMAddress().Hex())

	// zevm
	r.Logger.Print("ZetaChain:")
	r.Logger.Print("* ZETA balance:  %s", balances.ZetaZETA.String())
	r.Logger.Print("* WZETA balance: %s", balances.ZetaWZETA.String())
	r.Logger.Print("* ETH balance:   %s", balances.ZetaETH.String())
	r.Logger.Print("* ERC20 balance: %s", balances.ZetaERC20.String())
	r.Logger.Print("* BTC balance:   %s", balances.ZetaBTC.String())
	r.Logger.Print("* SOL balance: %s", balances.ZetaSOL.String())
	r.Logger.Print("* SPL balance: %s", balances.ZetaSPL.String())

	// evm
	r.Logger.Print("EVM:")
	r.Logger.Print("* ZETA balance:  %s", balances.EvmZETA.String())
	r.Logger.Print("* ETH balance:   %s", balances.EvmETH.String())
	r.Logger.Print("* ERC20 balance: %s", balances.EvmERC20.String())

	// bitcoin
	r.Logger.Print("Bitcoin:")
	r.Logger.Print("* BTC balance: %s", balances.BtcBTC)

	// solana
	r.Logger.Print("Solana:")
	if balances.SolSOL != nil {
		r.Logger.Print("* SOL balance: %s", balances.SolSOL.String())
	}
	if balances.SolSPL != nil {
		r.Logger.Print("* SPL balance: %s", balances.SolSPL.String())
	}
}

// PrintTotalDiff shows the difference in the account balances of the accounts used in the e2e test from two balances structs
func (r *E2ERunner) PrintTotalDiff(accoutBalancesDiff AccountBalancesDiff) {
	r.Logger.Print(" ---💰 Total gas spent ---")

	// show the value only if it is not zero
	if accoutBalancesDiff.ZETA.Cmp(big.NewInt(0)) != 0 {
		r.Logger.Print("* ZETA spent:  %s", accoutBalancesDiff.ZETA.String())
	}
	if accoutBalancesDiff.ETH.Cmp(big.NewInt(0)) != 0 {
		r.Logger.Print("* ETH spent:   %s", accoutBalancesDiff.ETH.String())
	}
	if accoutBalancesDiff.ERC20.Cmp(big.NewInt(0)) != 0 {
		r.Logger.Print("* ERC20 spent: %s", accoutBalancesDiff.ERC20.String())
	}
}

// GetAccountBalancesDiff returns the difference in the account balances of the accounts used in the E2E test
func GetAccountBalancesDiff(balancesBefore, balancesAfter AccountBalances) AccountBalancesDiff {
	balancesBeforeZeta := big.NewInt(0).Add(balancesBefore.ZetaZETA, balancesBefore.EvmZETA)
	balancesBeforeEth := big.NewInt(0).Add(balancesBefore.ZetaETH, balancesBefore.EvmETH)
	balancesBeforeErc20 := big.NewInt(0).Add(balancesBefore.ZetaERC20, balancesBefore.EvmERC20)

	balancesAfterZeta := big.NewInt(0).Add(balancesAfter.ZetaZETA, balancesAfter.EvmZETA)
	balancesAfterEth := big.NewInt(0).Add(balancesAfter.ZetaETH, balancesAfter.EvmETH)
	balancesAfterErc20 := big.NewInt(0).Add(balancesAfter.ZetaERC20, balancesAfter.EvmERC20)

	diffZeta := big.NewInt(0).Sub(balancesBeforeZeta, balancesAfterZeta)
	diffEth := big.NewInt(0).Sub(balancesBeforeEth, balancesAfterEth)
	diffErc20 := big.NewInt(0).Sub(balancesBeforeErc20, balancesAfterErc20)

	return AccountBalancesDiff{
		ETH:   diffEth,
		ZETA:  diffZeta,
		ERC20: diffErc20,
	}
}

// formatBalances formats the AccountBalancesDiff into a one-liner string
func formatBalances(balances AccountBalancesDiff) string {
	parts := []string{}
	if balances.ETH != nil && balances.ETH.Cmp(big.NewInt(0)) > 0 {
		parts = append(parts, fmt.Sprintf("ETH:%s", balances.ETH.String()))
	}
	if balances.ZETA != nil && balances.ZETA.Cmp(big.NewInt(0)) > 0 {
		parts = append(parts, fmt.Sprintf("ZETA:%s", balances.ZETA.String()))
	}
	if balances.ERC20 != nil && balances.ERC20.Cmp(big.NewInt(0)) > 0 {
		parts = append(parts, fmt.Sprintf("ERC20:%s", balances.ERC20.String()))
	}
	return strings.Join(parts, ",")
}
