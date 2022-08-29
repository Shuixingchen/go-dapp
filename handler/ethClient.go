package handler

import (
	"crypto/ecdsa"
	"log"

	"github.com/Shuixingchen/go-dapp/plugins/clients"
	"github.com/Shuixingchen/go-dapp/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ec  *ethclient.Client
	ecw *ethclient.Client
	evm *clients.EvmClient
)

func init() {
	initConfig()
	var err error
	ec, err = ethclient.Dial("http://54.255.92.123:8545")
	if err != nil {
		log.Fatal(err)
	}
	ecw, err = ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85")
	if err != nil {
		log.Fatal(err)
	}
	evm = clients.NewEvmClient(utils.Config.Nodes["eth"])
}

func initConfig() {
	utils.InitConfig("./configs/config.yaml")
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  common.Address
}

func InitWallet(privateHexKeys string) *Wallet {
	if privateHexKeys == "" {
		return nil
	}
	privateKey, err := crypto.HexToECDSA(privateHexKeys)
	if err != nil {
		return nil
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  crypto.PubkeyToAddress(privateKey.PublicKey),
	}
}
