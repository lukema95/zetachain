package mocks

import (
	"context"
	"encoding/hex"
	"sync"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/rpc/client/mock"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/stretchr/testify/require"
)

type CometBFTClient struct {
	mock.Client

	t      *testing.T
	err    error
	code   uint32
	txHash bytes.HexBytes

	subscribers     map[string]chan<- coretypes.ResultEvent
	subscribersLock sync.Mutex
}

func (c *CometBFTClient) BroadcastTxCommit(
	_ context.Context,
	_ tmtypes.Tx,
) (*coretypes.ResultBroadcastTxCommit, error) {
	return nil, c.err
}

func (c *CometBFTClient) BroadcastTxAsync(_ context.Context, _ tmtypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func (c *CometBFTClient) BroadcastTxSync(_ context.Context, _ tmtypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	log := ""
	if c.err != nil {
		log = c.err.Error()
	}
	return &coretypes.ResultBroadcastTx{
		Code:      c.code,
		Data:      bytes.HexBytes{},
		Log:       log,
		Codespace: "",
		Hash:      c.txHash,
	}, c.err
}

func (c *CometBFTClient) Tx(_ context.Context, _ []byte, _ bool) (*coretypes.ResultTx, error) {
	return &coretypes.ResultTx{
		Hash:   bytes.HexBytes{},
		Height: 0,
		Index:  0,
		TxResult: abci.ExecTxResult{
			Log: "",
		},
		Tx:    []byte{},
		Proof: tmtypes.TxProof{},
	}, c.err
}

func (c *CometBFTClient) Block(_ context.Context, _ *int64) (*coretypes.ResultBlock, error) {
	return &coretypes.ResultBlock{Block: &tmtypes.Block{
		Header:   tmtypes.Header{},
		Data:     tmtypes.Data{},
		Evidence: tmtypes.EvidenceData{},
	}}, c.err
}

func (c *CometBFTClient) SetBroadcastTxHash(hash string) *CometBFTClient {
	b, err := hex.DecodeString(hash)
	require.NoError(c.t, err)

	c.txHash = b

	return c
}

func (c *CometBFTClient) SetError(err error) *CometBFTClient {
	c.err = err
	return c
}

// PublishToSubscribers will publish an event to all subscribers (mock only)
func (c *CometBFTClient) PublishToSubscribers(event coretypes.ResultEvent) {
	c.subscribersLock.Lock()
	defer c.subscribersLock.Unlock()
	for _, ch := range c.subscribers {
		ch <- event
	}
}

func (c *CometBFTClient) Subscribe(
	_ context.Context,
	subscriber, _ string,
	_ ...int,
) (out <-chan coretypes.ResultEvent, err error) {
	outChan := make(chan coretypes.ResultEvent)
	c.subscribers[subscriber] = outChan
	return outChan, nil
}

func NewSDKClientWithErr(t *testing.T, err error, code uint32) *CometBFTClient {
	return &CometBFTClient{
		t:           t,
		Client:      mock.Client{},
		err:         err,
		code:        code,
		subscribers: make(map[string]chan<- coretypes.ResultEvent),
	}
}
