package consts

const (
	CoinTypeBitcoin         = "btc"
	CoinTypeBitcoinCash     = "bch"
	CoinTypeBitcoinCashABC  = "bcha"
	CoinTypeBitcoinSV       = "bsv"
	CoinTypeLitecoin        = "ltc"
	CoinTypeZcash           = "zec"
	CoinTypeDash            = "dash"
	CoinTypeEthereum        = "eth"
	CoinTypeEthereumClassic = "etc"
	CoinTypeBitcoinVault    = "btcv"
	CoinTypeEthereumFair    = "ethf"
	CoinTypeDoge            = "doge"
	CoinTypeZilliqa         = "zil"
)

const (
	BitcoinAddressTypeP2PKH            = "P2PKH"
	BitcoinAddressTypeP2SH             = "P2SH"
	BitcoinAddressTypeP2WPKHV0         = "P2WPKH_V0"
	BitcoinAddressTypeP2WSHV0          = "P2WSH_V0"
	BitcoinAddressTypePubkey           = "P2PKH_PUBKEY"
	BitcoinAddressTypeBECH32           = "BECH32"
	BitcoinAddressTypeWitnessV1Taproot = "WITNESS_V1_TAPROOT"
	BitcoinAddressTypeWitnessUnknown   = "WITNESS_UNKNOWN"
)

var BitcoinAddressTypeSegWitMap = map[string]string{
	"P2WSH_V0":            "segwit",
	"P2WPKH_V0":           "segwit",
	"P2SH_P2WSH_P2PKH":    "segwit",
	"P2SH_P2WSH_MULTISIG": "segwit",
	"P2SH_P2WSH":          "segwit",
	"P2PKH_PUBKEY":        "no-segwit",
	"P2SH_P2PKH":          "no-segwit",
	"NONSTANDARD":         "NONSTANDARD",
	"NULL_DATA":           "NULL_DATA",
	"P2PKH":               "no-segwit",
	"P2SH":                "no-segwit",
	"P2SH_P2WPKH":         "segwit",
	"P2SH_MULTISIG":       "no-segwit",
	"P2PKH_MULTISIG":      "no-segwit",
}

var DefaultCoinMap = map[string]string{
	CoinTypeBitcoin:         "btc",
	CoinTypeBitcoinCash:     "bch",
	CoinTypeBitcoinCashABC:  "bcha",
	CoinTypeBitcoinSV:       "bsv",
	CoinTypeLitecoin:        "ltc",
	CoinTypeZcash:           "zec",
	CoinTypeDash:            "dash",
	CoinTypeEthereum:        "eth",
	CoinTypeEthereumClassic: "etc",
	CoinTypeBitcoinVault:    "btcv",
}

var CoinBaseUnitName = map[string]string{
	CoinTypeBitcoin:         "Satoshi",
	CoinTypeBitcoinCash:     "Satoshi",
	CoinTypeBitcoinCashABC:  "Satoshi",
	CoinTypeBitcoinSV:       "Satoshi",
	CoinTypeLitecoin:        "Litoshi",
	CoinTypeZcash:           "Zatoshi",
	CoinTypeDash:            "Duffs",
	CoinTypeEthereum:        "Wei",
	CoinTypeEthereumClassic: "Wei",
	CoinTypeBitcoinVault:    "Satoshi",
}
