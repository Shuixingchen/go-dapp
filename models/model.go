package models

import "math/big"

type Token struct {
	Addr            string   `json:"tokenAddress"`
	Name            string   `json:"name"`
	Symbol          string   `json:"sym"`
	Decimals        int64    `json:"decimals"`
	InitTotalSupply *big.Int `json:"initTotalSupply"` // the maximum supply of a ERC20 token can be 2^256 - 1
	Type            string   `json:"tokenType"`
}

type Block struct {
	Number           HexUint64      `json:"number"`
	Hash             string         `json:"hash"`
	TimeStamp        HexUint64      `json:"timestamp"`
	ParentHash       string         `json:"parentHash"`
	Sha3Uncles       string         `json:"sha3Uncles"`
	TransactionsRoot string         `json:"transactionsRoot"`
	StateRoot        string         `json:"stateRoot"`
	ReceiptsRoot     string         `json:"receiptsRoot"`
	Miner            string         `json:"miner"`
	Difficulty       HexBig         `json:"difficulty"`
	TotalDifficulty  HexBig         `json:"totalDifficulty"`
	Size             HexUint64      `json:"size"`
	GasUsed          HexUint64      `json:"gasUsed"`
	GasLimit         HexUint64      `json:"gasLimit"`
	AvgGasPrice      *HexBig        `json:"gasPrice,omitempty"`
	Nonce            string         `json:"nonce"`
	BasicReward      *HexBig        `json:"blockBasicReward,omitempty"`
	TxFee            *HexBig        `json:"blockTxFee,omitempty"`
	UncleRefReward   *HexBig        `json:"blockUncleRefReward,omitempty"`
	TotalReward      *HexBig        `json:"blockReward,omitempty"`
	UnclesReward     *HexBig        `json:"unclesReward,omitempty"`
	ExtraData        string         `json:"extraData"`
	Uncles           []string       `json:"uncles"`
	Transactions     []*Transaction `json:"transactions"`
}

type Transaction struct {
	Hash                 string    `json:"hash"`
	Nonce                HexUint64 `json:"nonce"` // the number of transactions made by the sender prior to this one.
	TimeStamp            HexUint64 `json:"timestamp,omitempty"`
	BlockHash            string    `json:"blockHash"`
	BlockNumber          HexUint64 `json:"blockNumber"` // for pending transactions, the number might be null.
	Status               string    `json:"status,omitempty"`
	Sender               string    `json:"from"`
	Receiver             string    `json:"to"`
	Value                *HexBig   `json:"value"`
	GasLimit             HexUint64 `json:"gas"`
	GasUsed              HexUint64 `json:"gasUsed,omitempty"`
	MaxFeePerGas         *HexBig   `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *HexBig   `json:"maxPriorityFeePerGas,omitempty"`
	BaseFeePerGas        *HexBig   `json:"baseFeePerGas,omitempty"`
	GasPrice             *HexBig   `json:"gasPrice"`
	ActualCost           *HexBig   `json:"actualCost,omitempty"`
	InputData            string    `json:"input"`
	TxIndex              HexUint64 `json:"index"`
}

type TransactionReceipt struct {
	TransactionHash   string    `json:"transactionHash"`
	TransactionIndex  HexUint64 `json:"transactionIndex"`
	BlockHash         string    `json:"blockHash"`
	BlockNumber       HexUint64 `json:"blockNumber"`
	CumulativeGasUsed HexUint64 `json:"cumulativeGasUsed"`
	GasUsed           HexUint64 `json:"gasUsed"`
	ContractAddress   string    `json:"contractAddress,omitempty"`
	Root              string    `json:"root"`
	Status            string    `json:"status"`
	EffectiveGasPrice HexUint64 `json:"effectiveGasPrice,omitempty"` // eip1559
	From              string    `json:"from"`
	To                string    `json:"to"`
}
