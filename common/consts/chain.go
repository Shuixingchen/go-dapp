package consts

// BlockTxListSupportParam 区块交易列表排序参数
var BlockTxListSupportParam = map[string]string{
	"tx_block_id":    "block_idx",
	"tx_hash":        "hash",
	"fee":            "fee",
	"sigops":         "sigops",
	"size":           "size",
	"weight":         "weight",
	"inputs_count":   "inputs_count",
	"inputs_volume":  "inputs_value",
	"outputs_count":  "outputs_count",
	"outputs_volume": "outputs_value",
	"witness_hash":   "witness_hash",
}

// AddressTxListSupportParam 地址交易列表排序参数
var AddressTxListSupportParam = map[string]string{
	"time":     "addr_tx_idx",
	"received": "received",
	"sent":     "sent",
}

// 排序参数
var (
	TxHash    = "tx_hash"
	ASC       = "asc"
	DESC      = "desc"
	BlockIdx  = "block_idx"
	AddrTxIdx = "addr_tx_idx"
)

var AddressType = map[string]string{
	"01": "P2PKH",
	"02": "P2SH",
	"03": "P2WPKH_V0",
	"04": "P2WSH_V0",
	"05": "WITNESS_UNKNOWN",
	"06": "WITNESS_V1_TAPROOT",
}

// FlatBufferAddressType Address Display types
var FlatBufferAddressType = map[int32]string{
	BitcoinAddrDisplayTypeNONSTANDARD:      "NONSTANDARD",
	BitcoinAddrDisplayTypeNullData:         "NULL_DATA",
	BitcoinAddrDisplayTypeP2PKH:            "P2PKH",
	BitcoinAddrDisplayTypePubkey:           "P2PKH_PUBKEY",
	BitcoinAddrDisplayTypeP2PKHMultiSig:    "P2PKH_MULTISIG",
	BitcoinAddrDisplayTypeP2SH:             "P2SH",
	BitcoinAddrDisplayTypeP2SHP2WPKH:       "P2SH_P2WPKH",
	BitcoinAddrDisplayTypeP2SHP2WSH:        "P2SH_P2WSH",
	BitcoinAddrDisplayTypeP2WPKHV0:         "P2WPKH_V0",
	BitcoinAddrDisplayTypeP2WSHV0:          "P2WSH_V0",
	BitcoinAddrDisplayTypeWitnessV1Taproot: "WITNESS_V1_TAPROOT",
	BitcoinAddrDisplayTypeWitnessUnknown:   "WITNESS_UNKNOWN",
	AddrDisplayTypeWitnessMWEBPegin:        "WITNESS_MWEB_PEGIN",
	AddrDisplayTypeWitnessMWEBHogaddr:      "WITNESS_MWEB_HOGADDR",
}

// Chain Config AppName
var (
	ExternalChainBTC = "external-chain-v4-btc"
	ExternalChainBCH = "external-chain-v4-bch"
	ExternalChainBSV = "external-chain-v4-bsv"
	ExternalChainLTC = "external-chain-v4-ltc"
)

// Address Display Type Value
var (
	BitcoinAddrDisplayTypeNONSTANDARD int32 = 0
	// 实际上，现在返回的 NULL_DATA 的 type 值是 0
	BitcoinAddrDisplayTypeNullData      int32 = 16
	BitcoinAddrDisplayTypeP2PKH         int32 = 32
	BitcoinAddrDisplayTypePubkey        int32 = 33
	BitcoinAddrDisplayTypeP2PKHMultiSig int32 = 34
	BitcoinAddrDisplayTypeP2SH          int32 = 48
	BitcoinAddrDisplayTypeP2SHP2WPKH    int32 = 49
	BitcoinAddrDisplayTypeP2SHP2WSH     int32 = 50
	BitcoinAddrDisplayTypeP2WPKHV0      int32 = 64
	BitcoinAddrDisplayTypeP2WSHV0       int32 = 65
	// P2TR
	BitcoinAddrDisplayTypeWitnessV1Taproot int32 = 80
	// 实际上，现在返回的 WITNESS_UNKNOWN 的 type 值是 0
	BitcoinAddrDisplayTypeWitnessUnknown int32 = -1
	// MWEB
	AddrDisplayTypeWitnessMWEBPegin   int32 = 96
	AddrDisplayTypeWitnessMWEBHogaddr int32 = 97
)

// some stuff
const (
	CoinBasePrevTxHash       = "0000000000000000000000000000000000000000000000000000000000000000"
	LatestBlockNextBlockHash = "0000000000000000000000000000000000000000000000000000000000000000"
	Null                     = "null"
)

// block type
const (
	BlockHeight            = "height"
	BlockHash              = "hash"
	Latest                 = "latest"
	UnConfirmedTransaction = "Unconfirmed"
)

// CashAddrPrefix cash address prefix
var CashAddrPrefix = "bitcoincash:"
