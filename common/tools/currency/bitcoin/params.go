package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// MainNetParams 各个币种主网参数设置
var MainNetParams = map[string]chaincfg.Params{
	"btc":  BTCMainNetParams,
	"btcv": BTCVMainNetParams,
	"ltc":  LTCMainNetParams,
}

// BTCMainNetParams defines the network parameters for the main Bitcoin network.
var BTCMainNetParams = chaincfg.Params{
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "bc", // always bc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x00, // starts with 1
	ScriptHashAddrID:        0x05, // starts with 3
	PrivateKeyID:            0x80, // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh
}

// BTCVMainNetParams defines the network parameters for the main Bitcoin vault network.
var BTCVMainNetParams = chaincfg.Params{
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "royale", // always royale for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x4e, // starts with Y
	ScriptHashAddrID:        0x3c, // starts with R
	PrivateKeyID:            0x80, // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh
}

// LTCMainNetParams defines the network parameters for the main Bitcoin network.
var LTCMainNetParams = chaincfg.Params{
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "ltc", // always ltc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x30, // starts with L
	ScriptHashAddrID:        0x32, // starts with M
	PrivateKeyID:            0xB0, // starts with 6 (uncompressed) or T (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh
}
