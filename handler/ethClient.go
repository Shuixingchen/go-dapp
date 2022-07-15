package handler

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ec  *ethclient.Client
	ecw *ethclient.Client
)

func init() {
	var err error
	ec, err = ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85")
	if err != nil {
		log.Fatal(err)
	}
	ecw, err = ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85")
	if err != nil {
		log.Fatal(err)
	}

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
