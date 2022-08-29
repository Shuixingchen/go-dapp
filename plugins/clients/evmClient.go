package clients

import (
	"context"
	"math/big"

	"github.com/Shuixingchen/go-dapp/abis/artificial/erc20"
	"github.com/Shuixingchen/go-dapp/models"
	"github.com/Shuixingchen/go-dapp/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

type EvmClient struct {
	RPCClients []*rpc.Client
	Clients    []*ethclient.Client
}

func NewEvmClient(nodes []utils.Node) *EvmClient {
	clients := make([]*ethclient.Client, 0)
	rpcClients := make([]*rpc.Client, 0)
	for _, node := range nodes {
		c, err := rpc.DialContext(context.Background(), node.Addr)
		if err != nil {
			log.WithFields(log.Fields{"method:": "rpc.DialContext"}).Panic(err)
		}
		ec := ethclient.NewClient(c)
		clients = append(clients, ec)
		rpcClients = append(rpcClients, c)
	}
	return &EvmClient{
		Clients:    clients,
		RPCClients: rpcClients,
	}
}

func (ec *EvmClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (sub ethereum.Subscription, err error) {
	for _, c := range ec.Clients {
		sub, err = c.SubscribeFilterLogs(ctx, q, ch)
		if err == nil {
			return
		}
	}
	return
}

func (ec *EvmClient) BlockNumber() (blockNumber uint64, err error) {
	for _, c := range ec.Clients {
		blockNumber, err = c.BlockNumber(context.Background())
		if err != nil {
			continue
		}
		return
	}
	return
}

func (ec *EvmClient) FilterLogs(q ethereum.FilterQuery) (logs []types.Log, err error) {
	for _, c := range ec.Clients {
		logs, err = c.FilterLogs(context.Background(), q)
		if err != nil {
			continue
		}
		return
	}
	return
}

func (ec *EvmClient) BlockByNumber(number uint64) (block *models.Block, err error) {
	b := hexutil.EncodeUint64(number)
	for _, c := range ec.RPCClients {
		err = c.CallContext(context.Background(), &block, "eth_getBlockByNumber", b, true)
		if err != nil {
			log.WithFields(log.Fields{"method": "BlockByNumber"}).Error(err)
			continue
		}
		return
	}
	return
}

// TraceBlock returns traces created at given block.
func (ec *EvmClient) TraceBlock(number uint64) (r []*models.RPCTrace, err error) {
	block := hexutil.EncodeUint64(number)
	for _, c := range ec.RPCClients {
		var r []*models.RPCTrace
		err = c.CallContext(context.Background(), &r, "trace_block", block)
		if err != nil {
			continue
		}
		return r, nil
	}
	return
}

// TraceBlock returns traces created at given block.
func (ec *EvmClient) TraceTransaction(txHash string) (r []*models.RPCTrace, err error) {
	for _, c := range ec.RPCClients {
		var r []*models.RPCTrace
		err = c.CallContext(context.Background(), &r, "trace_transaction", txHash)
		if err != nil {
			continue
		}
		return r, nil
	}
	return
}

func (ec *EvmClient) GetTransactionReceiptsByBlock(number uint64) (r []*types.Receipt, err error) {
	block := hexutil.EncodeUint64(number)
	for _, c := range ec.RPCClients {
		var r []*types.Receipt
		err = c.CallContext(context.Background(), &r, "eth_getTransactionReceiptsByBlock", block)
		if err != nil {
			continue
		}
		return r, nil
	}
	return
}

func (ec *EvmClient) GetTransactionReceipt(ctx context.Context, hash string) (r *models.TransactionReceipt, err error) {
	for _, c := range ec.RPCClients {
		err := c.CallContext(ctx, &r, "eth_getTransactionReceipt", hash)
		if err != nil {
			continue
		}
		return r, nil
	}
	return r, err
}

func (ec *EvmClient) GetTokenCaller(tokenAddr string) (*erc20.Erc20Caller, error) {
	contractAddress := common.HexToAddress(tokenAddr)
	return erc20.NewErc20Caller(contractAddress, ec.Clients[0])
}

func (ec *EvmClient) GetTokenBalance(tokenAddr string, holderAddr string) (*big.Int, error) {
	tc, err := ec.GetTokenCaller(tokenAddr)
	if err != nil {
		log.WithField("method", "GetTokenCaller").Error(err)
		return big.NewInt(0), err
	}
	return tc.BalanceOf(nil, common.HexToAddress(holderAddr))
}
