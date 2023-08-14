package consts

// BTCCOMMinerAddress btc.com pool miner address
var BTCCOMMinerAddress = map[string][]string{
	CoinTypeBitcoin:         BTCCOMBTCMinerAddress,
	CoinTypeBitcoinCash:     BTCCOMBCHMinerAddress,
	CoinTypeLitecoin:        BTCCOMLTCMinerAddress,
	CoinTypeDash:            BTCCOMDashMinerAddress,
	CoinTypeEthereum:        BTCCOMETHMinerAddress,
	CoinTypeEthereumClassic: BTCCOMETCMinerAddress,
	CoinTypeEthereumFair:    BTCCOMETHFMinerAddress,
	CoinTypeZilliqa:         BTCCOMZILMinerAddress,
}

// PoolETHMinerAddress pool eth miner address
var PoolETHMinerAddress = map[string]string{
	"btc.com": "0xeea5b82b61424df8020f5fedd81767f2d0d25bfb",
	"huobi":   "0x1b50e8779040b182cc2322079144d225b0361b75",
	"binance": "0xe72f79190bc8f92067c6a62008656c6a9077f6aa",
}

var (
	// BTCCOMPool btc.com pool name
	BTCCOMPool             = "BTC.com"
	BTCCOMBTCMinerAddress  = []string{"1Bf9sZvBHPFGVPX71WX2njhd1NXKv5y7v5"}
	BTCCOMBCHMinerAddress  = []string{"bitcoincash:qpv5y82t8z7n6w80fpm64afah7ntptxue59h5cdsn2"}
	BTCCOMLTCMinerAddress  = []string{"LMs7eqZhREmAP4xpmXi6QQxVaqTYqPFTFK"}
	BTCCOMDashMinerAddress = []string{"DsaF6Va2L4kAn2uDJgqWeWVf9bYZrM7xyq6"}
	BTCCOMETHMinerAddress  = []string{"0xeea5b82b61424df8020f5fedd81767f2d0d25bfb", "0x5b310960a7922092fdcb9295ece336012f9cf87e"}
	BTCCOMETCMinerAddress  = []string{"0xb35c9BA58293D4542b974dBC2c93CA6A9EdD4D17", "0x8B167bee2A5135b23c52550098EF39A47a0E3DE9"}
	BTCCOMETHFMinerAddress = []string{"0x0d80AF2ADd1df1268EFC7982e7c87fb9822B7722"}
	BTCCOMZILMinerAddress  = []string{"zil1ma88rhdjcawx2905t47pghcup9lwry57rq7m3p", "zil1pzs9jk7u2g2r5hwwux6u79pwvlp6aaa8quae74", "zil1azqyxnc9e45nptzlctz7r0ryq8fq00w872kfvg"}
)
