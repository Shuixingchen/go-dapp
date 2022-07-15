package handler

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Shuixingchen/go-dapp/contract/artificial/erc721"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	TokenAddr  = "0xCB5C3e56DAAbBAFF5252F7600CA8f005A33D61B7"                       // 地址
	privateHex = "19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107" // 地址对应的私钥
)

// 使用go构造一个tx, 交易触发合约的setA(uint256)方法
//tx.Data 包含字符串，

type transferFromParam struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
}

func SendTx() {
	wallet := InitWallet(privateHex)

	// 创建签名交易
	sigTransaction, err := singTx(wallet)

	// 四、发送交易
	err = ec.SendTransaction(context.Background(), sigTransaction)
	if err != nil {
		fmt.Println("ethClient.SendTransaction failed: ", err.Error())
		return
	}
	fmt.Println("send transaction success,tx: ", sigTransaction.Hash().Hex())
}

func singTx(wallet *Wallet) (*types.Transaction, error) {
	// 一、ABI编码请求参数
	params := transferFromParam{
		From:    common.HexToAddress("0xe725D38CC421dF145fEFf6eB9Ec31602f95D8097"),
		To:      common.HexToAddress("0xD9478B7cf6C4ACD11e90701Aa6C335B93a2C2368"),
		TokenID: big.NewInt(0),
	}
	abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
	callData, err := abi.Pack("transferFrom", params)
	if err != nil {
		fmt.Println("abi.Pack", err)
	}

	// 二、构造交易对象
	nonce, _ := ec.NonceAt(context.Background(), wallet.PublicKey, nil)
	gasPrice, _ := ec.SuggestGasPrice(context.Background())
	value := big.NewInt(0)
	gasLimit := uint64(3000000)
	rawTx := types.NewTransaction(nonce, common.HexToAddress(TokenAddr), value, gasLimit, gasPrice, callData)

	// 三、交易签名
	chainID, err := ec.NetworkID(context.Background())
	sigTransaction, err := types.SignTx(rawTx, types.NewEIP155Signer(chainID), wallet.PrivateKey)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
		return nil, err
	}
	return sigTransaction, nil
}

// 验证交易，比较签名恢复得到的地址与tx.sender
func ParseAddressFromSigTx() {
	wallet := InitWallet(privateHex)
	sigTransaction, err := singTx(wallet)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
	}
	// 使用相同的签名器，eip155后修改了签名器
	signer := types.NewEIP155Signer(big.NewInt(1))
	senderAddr, err := signer.Sender(sigTransaction)

	fmt.Println("sender: ", senderAddr.Hex())
	fmt.Println("addressHex: ", wallet.PublicKey)
}
