package handler

import (
	"encoding/hex"

	"github.com/Shuixingchen/go-dapp/common/btc"
	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/Shuixingchen/go-dapp/common/tools/blockchain"
	"github.com/Shuixingchen/go-dapp/common/tools/currency"
	"github.com/golang/glog"
)

// btc addr 转address_id
func BTCAddr() {
	var address, coin string
	keyword := "bc1qxhmdufsvnuaaaer4ynz88fspdsxq2h9e9cetdj"
	// 正则匹配检查
	if !blockchain.InputCheck(keyword) {
		glog.Error("keyword=%s is invalid", keyword)
		return
	}
	// 判断是地址还是txid
	if blockchain.IsBlockHashOrTransactionHash(keyword) {
		glog.Errorf("keyword=%s is tx_id", keyword)
	}
	if blockchain.IsAddress(keyword) {
		glog.Errorf("keyword=%s is address", keyword)
	}
	address = keyword
	if blockchain.IsLTCAddress(address) {
		coin = consts.CoinTypeLitecoin
	}

	if blockchain.IsCashAddr(address) && coin != consts.CoinTypeBitcoinCash {
		coin = consts.CoinTypeBitcoinCash
	}

	if blockchain.IsBC1Address(address) && coin != consts.CoinTypeBitcoin {
		coin = consts.CoinTypeBitcoin
	}
	// convert address to address ID
	// check the field
	addressType := currency.GuessType(address)
	if addressType == "" {
		glog.Errorf("addressType=%s", addressType)
		return
	}
	addressID, err := currency.GetAddressID(coin, address, addressType)
	if err != nil {
		glog.Error(err)
	}
	glog.Infof("addressID=%s, addressType=%s \n", addressID, addressType)
}

func ParserAddrFromScript() {
	pkScript, _ := hex.DecodeString("a9144b09d828dfc8baaba5d04ee77397e04b1050cc7387")
	res := btc.ParsePkScript(pkScript)
	glog.Infof("%+v", res)
}
