package main

import (
	"flag"
	"fmt"

	"github.com/Shuixingchen/go-dapp/handler"
	"github.com/golang/glog"
)

func init() {
	glog.MaxSize = uint64(1 * 1024 * 1024)
	consoleLevel := fmt.Sprintf("-stderrthreshold=%s", "info")
	infoLevel := fmt.Sprintf("-v=%d", 1)
	_ = flag.CommandLine.Parse([]string{consoleLevel})
	_ = flag.CommandLine.Parse([]string{infoLevel})
}
func main() {
	BTCAddr()
}
func BTCAddr() {
	// handler.BTCAddr()
	handler.ParserAddrFromScript()
}

func ETHContract() {
	// handler.AccountInfo()
	// 反编译合约
	handler.CheckContractFunction("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	// handler.CallData()
	// handler.QueryERC1155Balance("0x7aAd8FdEBa6a9655E37fE3AEb908B67dAD83b935", "0x787B093C62F0f6bd4e599dA1435f9FCdFb87B9E7", "1")
	// handler.PubkeyToAddress()
	// handler.KeyShow()
	// handler.PenddingTx()
}
