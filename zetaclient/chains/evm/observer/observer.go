// Package observer implements the EVM chain observer
package observer

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/zeta-chain/protocol-contracts/pkg/erc20custody.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/gatewayevm.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/zeta.non-eth.sol"
	zetaconnectoreth "github.com/zeta-chain/protocol-contracts/pkg/zetaconnector.eth.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/zetaconnector.non-eth.sol"

	"github.com/zeta-chain/node/zetaclient/chains/base"
	"github.com/zeta-chain/node/zetaclient/chains/evm/client"
	"github.com/zeta-chain/node/zetaclient/chains/evm/common"
	"github.com/zeta-chain/node/zetaclient/chains/interfaces"
	"github.com/zeta-chain/node/zetaclient/metrics"
)

// Observer is the observer for evm chains
type Observer struct {
	// base.Observer implements the base chain observer
	*base.Observer

	priorityFeeConfig

	// evmClient is the EVM client for the observed chain
	evmClient interfaces.EVMRPCClient

	// outboundConfirmedReceipts is the map to index confirmed receipts by hash
	outboundConfirmedReceipts map[string]*ethtypes.Receipt

	// outboundConfirmedTransactions is the map to index confirmed transactions by hash
	outboundConfirmedTransactions map[string]*ethtypes.Transaction
}

// priorityFeeConfig is the configuration for priority fee
type priorityFeeConfig struct {
	// checked indicates whether the observer checked
	// this EVM chain for EIP-1559 (further checks are cached)
	checked bool

	// supported indicates whether this EVM chain supports EIP-1559
	supported bool
}

// New Observer constructor
func New(baseObserver *base.Observer, client interfaces.EVMRPCClient) (*Observer, error) {
	// create evm observer
	ob := &Observer{
		Observer:                      baseObserver,
		evmClient:                     client,
		outboundConfirmedReceipts:     make(map[string]*ethtypes.Receipt),
		outboundConfirmedTransactions: make(map[string]*ethtypes.Transaction),
		priorityFeeConfig:             priorityFeeConfig{},
	}

	// load last block scanned
	if err := ob.LoadLastBlockScanned(context.Background()); err != nil {
		return nil, errors.Wrap(err, "unable to load last block scanned")
	}

	return ob, nil
}

// GetConnectorContract returns the non-Eth connector address and binder
func (ob *Observer) GetConnectorContract() (ethcommon.Address, *zetaconnector.ZetaConnectorNonEth, error) {
	addr := ethcommon.HexToAddress(ob.ChainParams().ConnectorContractAddress)
	contract, err := zetaconnector.NewZetaConnectorNonEth(addr, ob.evmClient)
	return addr, contract, err
}

// GetConnectorContractEth returns the Eth connector address and binder
func (ob *Observer) GetConnectorContractEth() (ethcommon.Address, *zetaconnectoreth.ZetaConnectorEth, error) {
	addr := ethcommon.HexToAddress(ob.ChainParams().ConnectorContractAddress)
	contract, err := FetchConnectorContractEth(addr, ob.evmClient)
	return addr, contract, err
}

// GetERC20CustodyContract returns ERC20Custody contract address and binder
func (ob *Observer) GetERC20CustodyContract() (ethcommon.Address, *erc20custody.ERC20Custody, error) {
	addr := ethcommon.HexToAddress(ob.ChainParams().Erc20CustodyContractAddress)
	contract, err := erc20custody.NewERC20Custody(addr, ob.evmClient)
	return addr, contract, err
}

// GetERC20CustodyV2Contract returns ERC20Custody contract address and binder
// NOTE: we use the same address as gateway v1
// this simplify the migration process v1 will be completely removed in the future
// currently the ABI for withdraw is identical, therefore both contract instances can be used
func (ob *Observer) GetERC20CustodyV2Contract() (ethcommon.Address, *erc20custody.ERC20Custody, error) {
	addr := ethcommon.HexToAddress(ob.ChainParams().Erc20CustodyContractAddress)
	contract, err := erc20custody.NewERC20Custody(addr, ob.evmClient)
	return addr, contract, err
}

// GetGatewayContract returns the gateway contract address and binder
func (ob *Observer) GetGatewayContract() (ethcommon.Address, *gatewayevm.GatewayEVM, error) {
	addr := ethcommon.HexToAddress(ob.ChainParams().GatewayAddress)
	contract, err := gatewayevm.NewGatewayEVM(addr, ob.evmClient)
	return addr, contract, err
}

// FetchConnectorContractEth returns the Eth connector address and binder
// TODO(revamp): move this to a contract package
func FetchConnectorContractEth(
	addr ethcommon.Address,
	client interfaces.EVMRPCClient,
) (*zetaconnectoreth.ZetaConnectorEth, error) {
	return zetaconnectoreth.NewZetaConnectorEth(addr, client)
}

// FetchZetaTokenContract returns the non-Eth ZETA token binder
// TODO(revamp): move this to a contract package
func FetchZetaTokenContract(
	addr ethcommon.Address,
	client interfaces.EVMRPCClient,
) (*zeta.ZetaNonEth, error) {
	return zeta.NewZetaNonEth(addr, client)
}

// SetTxNReceipt sets the receipt and transaction in memory
func (ob *Observer) SetTxNReceipt(nonce uint64, receipt *ethtypes.Receipt, transaction *ethtypes.Transaction) {
	ob.Mu().Lock()
	defer ob.Mu().Unlock()
	ob.outboundConfirmedReceipts[ob.OutboundID(nonce)] = receipt
	ob.outboundConfirmedTransactions[ob.OutboundID(nonce)] = transaction
}

// GetTxNReceipt gets the receipt and transaction from memory
func (ob *Observer) GetTxNReceipt(nonce uint64) (*ethtypes.Receipt, *ethtypes.Transaction) {
	ob.Mu().Lock()
	defer ob.Mu().Unlock()
	receipt := ob.outboundConfirmedReceipts[ob.OutboundID(nonce)]
	transaction := ob.outboundConfirmedTransactions[ob.OutboundID(nonce)]
	return receipt, transaction
}

// IsTxConfirmed returns true if there is a confirmed tx for 'nonce'
func (ob *Observer) IsTxConfirmed(nonce uint64) bool {
	ob.Mu().Lock()
	defer ob.Mu().Unlock()
	id := ob.OutboundID(nonce)

	return ob.outboundConfirmedReceipts[id] != nil && ob.outboundConfirmedTransactions[id] != nil
}

// CheckTxInclusion returns nil only if tx is included at the position indicated by the receipt ([block, index])
func (ob *Observer) CheckTxInclusion(ctx context.Context, tx *ethtypes.Transaction, receipt *ethtypes.Receipt) error {
	block, err := ob.GetBlockByNumberCached(ctx, receipt.BlockNumber.Uint64())
	if err != nil {
		return errors.Wrapf(err, "GetBlockByNumberCached error for block %d txHash %s nonce %d",
			receipt.BlockNumber.Uint64(), tx.Hash(), tx.Nonce())
	}

	// #nosec G115 non negative value
	if receipt.TransactionIndex >= uint(len(block.Transactions)) {
		return fmt.Errorf("transaction index %d out of range [0, %d), txHash %s nonce %d block %d",
			receipt.TransactionIndex, len(block.Transactions), tx.Hash(), tx.Nonce(), receipt.BlockNumber.Uint64())
	}

	txAtIndex := block.Transactions[receipt.TransactionIndex]
	if !strings.EqualFold(txAtIndex.Hash, tx.Hash().Hex()) {
		ob.RemoveCachedBlock(receipt.BlockNumber.Uint64()) // clean stale block from cache
		return fmt.Errorf("transaction at index %d has different hash %s, txHash %s nonce %d block %d",
			receipt.TransactionIndex, txAtIndex.Hash, tx.Hash(), tx.Nonce(), receipt.BlockNumber.Uint64())
	}

	return nil
}

// TransactionByHash query transaction by hash via JSON-RPC
func (ob *Observer) TransactionByHash(ctx context.Context, txHash string) (*client.Transaction, bool, error) {
	tx, err := ob.evmClient.TransactionByHashCustom(ctx, txHash)
	if err != nil {
		return nil, false, err
	}
	err = common.ValidateEvmTransaction(tx)
	if err != nil {
		return nil, false, err
	}
	return tx, tx.BlockNumber == nil, nil
}

func (ob *Observer) TransactionReceipt(ctx context.Context, hash ethcommon.Hash) (*ethtypes.Receipt, error) {
	return ob.evmClient.TransactionReceipt(ctx, hash)
}

// GetBlockByNumberCached get block by number from cache
// returns block, ethrpc.Block, isFallback, isSkip, error
func (ob *Observer) GetBlockByNumberCached(ctx context.Context, blockNumber uint64) (*client.Block, error) {
	if result, ok := ob.BlockCache().Get(blockNumber); ok {
		if block, ok := result.(*client.Block); ok {
			return block, nil
		}
		return nil, errors.New("cached value is not of type *client.Block")
	}
	if blockNumber > math.MaxInt32 {
		return nil, fmt.Errorf("block number %d is too large", blockNumber)
	}
	// #nosec G115 always in range, checked above
	block, err := ob.BlockByNumber(ctx, int(blockNumber))
	if err != nil {
		return nil, err
	}
	ob.BlockCache().Add(blockNumber, block)
	return block, nil
}

// RemoveCachedBlock remove block from cache
func (ob *Observer) RemoveCachedBlock(blockNumber uint64) {
	ob.BlockCache().Remove(blockNumber)
}

// BlockByNumber query block by number via JSON-RPC
func (ob *Observer) BlockByNumber(ctx context.Context, blockNumber int) (*client.Block, error) {
	block, err := ob.evmClient.BlockByNumberCustom(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block not found: %d", blockNumber)
	}
	for i := range block.Transactions {
		err := common.ValidateEvmTransaction(&block.Transactions[i])
		if err != nil {
			return nil, err
		}
	}
	return block, nil
}

// LoadLastBlockScanned loads the last scanned block from the database
// TODO(revamp): move to a db file
func (ob *Observer) LoadLastBlockScanned(ctx context.Context) error {
	err := ob.Observer.LoadLastBlockScanned(ob.Logger().Chain)
	if err != nil {
		return errors.Wrapf(err, "error LoadLastBlockScanned for chain %d", ob.Chain().ChainId)
	}

	// observer will scan from the last block when 'lastBlockScanned == 0', this happens when:
	// 1. environment variable is set explicitly to "latest"
	// 2. environment variable is empty and last scanned block is not found in DB
	if ob.LastBlockScanned() == 0 {
		blockNumber, err := ob.evmClient.BlockNumber(ctx)
		if err != nil {
			return errors.Wrapf(err, "error BlockNumber for chain %d", ob.Chain().ChainId)
		}
		ob.WithLastBlockScanned(blockNumber)
	}
	ob.Logger().Chain.Info().Uint64("last_block_scanned", ob.LastBlockScanned()).Msg("LoadLastBlockScanned succeed")

	return nil
}

func (ob *Observer) CheckRPCStatus(ctx context.Context) error {
	blockTime, err := ob.evmClient.HealthCheck(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to check rpc health")
	}

	ob.ReportBlockLatency(blockTime)
	ob.Logger().Chain.Info().Msg("CheckRPCStatus succeed")

	return nil
}

// updateLastBlock is a helper function to update the last block number.
// Note: keep last block up-to-date helps to avoid inaccurate confirmation.
func (ob *Observer) updateLastBlock(ctx context.Context) error {
	blockNumber, err := ob.evmClient.BlockNumber(ctx)
	switch {
	case err != nil:
		return errors.Wrap(err, "error getting block number")
	case blockNumber < ob.LastBlock():
		return fmt.Errorf("block number should not decrease: current %d last %d", blockNumber, ob.LastBlock())
	default:
		ob.WithLastBlock(blockNumber)
	}

	// increment prom counter
	metrics.GetBlockNumberPerChain.WithLabelValues(ob.Chain().Name).Inc()

	return nil
}
