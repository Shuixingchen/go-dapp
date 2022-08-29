package handler

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Shuixingchen/go-dapp/models"
	log "github.com/sirupsen/logrus"
)

func Calculate() {
	number, err := evm.BlockNumber()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error(err)
	}
	block, _ := evm.BlockByNumber(number)
	mainTTD, _ := models.StringToHexBig("58750000000000000000000")
	res := big.NewInt(0)
	num := big.NewInt(0)
	res.Sub(mainTTD.Int, block.TotalDifficulty.Int)
	num.Div(res, block.Difficulty.Int)
	totalsecond := num.Int64() * 15
	t := time.Now().Add(time.Duration(totalsecond) * time.Second)
	fmt.Println(t)
}
