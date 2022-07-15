// package handler

// import (
// 	"bytes"
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math/big"
// 	"net/http"
// 	"strings"

// 	"github.com/Shuixingchen/go-dapp/contract/artificial/erc721"
// 	"github.com/Shuixingchen/go-dapp/contract/client"
// 	"github.com/Shuixingchen/go-dapp/contract/models"

// 	"github.com/Shuixingchen/go-dapp/contract/artificial/erc20"

// 	log "github.com/sirupsen/logrus"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// )

// var (
// 	transferEventSig  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" //Transfer(address,address,uint256)
// 	transferSingleSig = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
// 	transferBatchSig  = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
// 	ordersMatchedSig  = "0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9"
// 	approvalSig       = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
// 	approvalForAllSig = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
// )
// var (
// 	InterfaceIdErc165           = [4]byte{1, 255, 201, 167}  // 0x01ffc9a7
// 	InterfaceIdErc721           = [4]byte{128, 172, 88, 205} // 0x80ac58cd
// 	InterfaceIdErc721Metadata   = [4]byte{91, 94, 19, 159}   // 0x5b5e139f
// 	InterfaceIdErc721Enumerable = [4]byte{120, 14, 157, 99}  // 0x780e9d63
// 	InterfaceIdErc1155          = [4]byte{217, 182, 122, 38} // 0xd9b67a26
// )

// // nft 转移
// type TransferEvent struct {
// 	From    common.Address
// 	To      common.Address
// 	TokenId *big.Int
// }
// type TransferOldEvent struct {
// 	From      string
// 	To        string
// 	TokenId   *big.Int
// 	TokenAddr string
// 	Value     *big.Int //erc20
// 	IsNFT     bool
// 	TokenType int
// }

// // erc1155
// type TransferSingleEvent struct {
// 	Operator common.Address
// 	From     common.Address
// 	To       common.Address
// 	Id       *big.Int
// 	Value    *big.Int
// }
// type TransferBatchEvent struct {
// 	Operator common.Address
// 	From     common.Address
// 	To       common.Address
// 	Ids      []*big.Int
// 	Values   []*big.Int
// }

// // opensea 交易
// type OrdersMatchedEvent struct {
// 	BuyHash  [32]byte
// 	SellHash [32]byte
// 	Maker    common.Address
// 	Taker    common.Address
// 	Price    *big.Int
// 	Metadata [32]byte
// }

// // approve
// type ApprovalEvent struct {
// 	Owner    common.Address
// 	Approved common.Address
// 	TokenId  *big.Int
// }

// // ApproveForAll
// type ApprovalForAllEvent struct {
// 	Owner    common.Address
// 	Operator common.Address
// 	Approved bool
// }

// // function input atomicMatch_ params
// type ParamsAtomicMatch_ struct {
// 	addrs                          [14]common.Address
// 	uints                          [18]*big.Int
// 	feeMethodsSidesKindsHowToCalls [8]uint8
// 	calldataBuy                    []byte
// 	calldataSell                   []byte
// 	replacementPatternBuy          []byte
// 	replacementPatternSell         []byte
// 	staticExtradataBuy             []byte
// 	staticExtradataSell            []byte
// 	vs                             [2]uint8
// 	rssMetadata                    [2][32]byte
// }

// func AccountInfo() {
// 	account := common.HexToAddress("0xe725D38CC421dF145fEFf6eB9Ec31602f95D8097")
// 	balance, err := ec.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(balance) // 25893180161173005034
// }

// // 与erc20合约交互
// func QueryERC20(addr, holder string) {
// 	contractAddr := common.HexToAddress(addr)
// 	var token = new(models.Token)
// 	var err error
// 	// 加载智能合约
// 	tc, err := erc20.NewErc20(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	token.Addr = addr
// 	token.Name, err = tc.Name(nil)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the erc 20 token name")
// 	}
// 	token.Symbol, err = tc.Symbol(nil)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the symbol of the erc20 token")
// 	}
// 	decimals, err := tc.Decimals(nil)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the decimals of the erc 20 token")
// 	}
// 	token.Decimals = int64(decimals)
// 	supply, err := tc.TotalSupply(nil)
// 	if err != nil || supply == nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the total supply of the erc 20 token")
// 	}
// 	if supply == nil {
// 		supply = new(big.Int)
// 	}
// 	token.InitTotalSupply = supply
// 	b, err := tc.BalanceOf(nil, common.HexToAddress(holder))
// 	if err != nil {
// 		log.WithFields(log.Fields{"tokenaddr": addr, "holder": holder}).Error(err)
// 	}
// 	fmt.Println(token, b)
// }

// // 与erc721合约交互
// func QueryERC721(addr string) {
// 	contractAddr := common.HexToAddress(addr)
// 	var token = new(models.Token)
// 	var err error
// 	// 加载智能合约
// 	tc, err := erc721.NewErc721(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	token.Symbol, err = tc.Symbol(nil)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the symbol of the erc721 token")
// 	}
// 	token.Name, err = tc.Name(nil)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get the erc721 token name")
// 	}
// 	// 查询NFT
// 	tokenId := big.NewInt(1)
// 	ownerAddr, err := tc.OwnerOf(nil, tokenId)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get tokenId ownerof")
// 	}
// 	fmt.Println("NFT id=1 owner is:" + ownerAddr.Hex())
// 	tokenURI, err := tc.TokenURI(nil, tokenId)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable to get tokenId tokenURI")
// 	}
// 	fmt.Println("NFT id=2 uri is:" + tokenURI)
// 	// 查询某个地址持有nft数量erc721
// 	owner := common.HexToAddress("0x2e02B22B72F5f2334fDE5aC62B817E8854576e9D")
// 	balance, err := tc.BalanceOf(nil, owner)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("unable get balance")
// 	}
// 	fmt.Println("owner:" + owner.String() + " balance:" + balance.String())

// 	fmt.Println(owner)
// }
// func QueryTokenURI(addr string, num string) {
// 	tokenId := new(big.Int)
// 	tokenId, ok := tokenId.SetString(num, 10)
// 	if !ok {
// 		log.Fatal("setstring fail")
// 	}
// 	log.Info(tokenId)
// 	contractAddr := common.HexToAddress(addr)
// 	// 加载智能合约
// 	tc, err := erc721.NewErc721(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tokenURI, err := tc.TokenURI(nil, tokenId)
// 	if err != nil {
// 		log.WithField("tokenURI", tokenId).Error(err)
// 	}
// 	fmt.Println(tokenURI)
// 	tokenURI, err = tc.Uri(nil, tokenId)
// 	if err != nil {
// 		log.WithField("tokenURI", tokenId).Error(err)
// 	}
// 	fmt.Println(tokenURI)
// }
// func QueryERC721Balance(addr string, num string) {
// 	tokenId := new(big.Int)
// 	tokenId, ok := tokenId.SetString(num, 10)
// 	if !ok {
// 		log.Fatal("setstring fail")
// 	}
// 	log.Info(tokenId)
// 	contractAddr := common.HexToAddress(addr)
// 	// 加载智能合约
// 	tc, err := erc721.NewErc721(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	IsNFTByClient(addr)
// 	owner, err := tc.OwnerOf(nil, tokenId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(owner)
// }
// func QueryERC1155Balance(addr, owner, num string) {
// 	tokenId := new(big.Int)
// 	tokenId, ok := tokenId.SetString(num, 10)
// 	if !ok {
// 		log.Fatal("setstring fail")
// 	}
// 	contractAddr := common.HexToAddress(addr)
// 	// 加载智能合约
// 	tc, err := erc721.NewErc721(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	o, err := tc.OwnerOf(nil, tokenId)
// 	b, err := tc.BalanceOf(nil, common.HexToAddress(owner))
// 	ids := []*big.Int{tokenId}
// 	owners := []common.Address{common.HexToAddress(owner)}
// 	balances, err := tc.BalanceOfBatch(nil, owners, ids)
// 	fmt.Println(o, b, balances)
// }

// // 执行合约方法
// func ExecERC20(addr string, to string) {
// 	contractAddr := common.HexToAddress(addr)
// 	receiver := common.HexToAddress(to)
// 	// 加载智能合约
// 	tc, err := erc20.NewErc20(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	txOpts := GetTxOpts()
// 	value := big.NewInt(10000000)
// 	tx, err := tc.Transfer(txOpts, receiver, value)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("transfer")
// 	}
// 	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
// }
// func ExecERC721(addr, to, tokenUrl string) {
// 	contractAddr := common.HexToAddress(addr)
// 	receiver := common.HexToAddress(to)
// 	tc, err := erc721.NewErc721(contractAddr, ec)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	txOpts := GetTxOpts()
// 	// 铸造NFT, 不能立即知道铸造的nft的id,只能监听合约的铸造事件获取函数执行结果
// 	tx, err := tc.AwardItem(txOpts, receiver, tokenUrl)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("transfer")
// 	}
// 	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
// }

// // 构造交易选项
// func GetTxOpts() *bind.TransactOpts {
// 	// 私钥转为ECDSA的privateKey
// 	privateKey, err := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	// 获取当前地址的noce,当前的gasPrice
// 	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	gasPrice, err := ec.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// 新建一个keyed transactor
// 	auth := bind.NewKeyedTransactor(privateKey)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // in wei
// 	auth.GasLimit = uint64(300000) // in units
// 	auth.GasPrice = gasPrice
// 	return auth
// }

// // 订阅事件,客户端需要ws连接
// func SubEvent() {
// 	contractAddress := common.HexToAddress("0xfF06b40b853b2700Afa5019aBE084469F10b63a5")
// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{contractAddress},
// 	}
// 	logs := make(chan types.Log)
// 	// 订阅指定的log
// 	sub, err := ecw.SubscribeFilterLogs(context.Background(), query, logs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ec.FilterLogs(context.Background(), query)
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:
// 			fmt.Println(vLog) // pointer to event log
// 		}
// 	}
// }

// // 查询event log
// func QueryEventLog(addr string, start, end uint64) {
// 	// var txSingleLog TransferSingleEvent
// 	transfer := common.HexToHash(transferEventSig)
// 	transferSingle := common.HexToHash(transferSingleSig)
// 	transferBatch := common.HexToHash(transferBatchSig)
// 	approval := common.HexToHash(approvalSig)
// 	approvalForAll := common.HexToHash(approvalForAllSig)
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(int64(start)),
// 		ToBlock:   big.NewInt(int64(end)),
// 		Topics:    [][]common.Hash{{transfer, transferSingle, transferBatch, approval, approvalForAll}},
// 	}
// 	logs, err := ec.FilterLogs(context.Background(), query)
// 	fmt.Println(len(logs))
// 	if err != nil {
// 		log.WithFields(log.Fields{"method": "FilterLogs"}).Error(err)
// 	}
// 	abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
// 	if err != nil {
// 		log.WithFields(log.Fields{"method": "abi.JSON"}).Error(err)
// 	}
// 	for _, log := range logs {
// 		switch log.Topics[0].Hex() {
// 		case transferEventSig:
// 			fmt.Printf("Log Name: Transfer\n")
// 			var transferEvent TransferOldEvent
// 			if log.Address.Hex() == "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d" {
// 				err = abi.UnpackIntoInterface(&transferEvent, "TransferOld", log.Data) //data只包含未index的参数
// 			} else {
// 				err = abi.UnpackIntoInterface(&transferEvent, "Transfer", log.Data) //data只包含未index的参数
// 			}
// 			if err != nil {
// 				fmt.Println("abi", err)
// 			}

// 			fmt.Println(transferEvent)
// 		case transferSingleSig:
// 			fmt.Printf("Log Name: TransferSingle\n")
// 			var transferEvent TransferSingleEvent
// 			err := abi.UnpackIntoInterface(&transferEvent, "TransferSingle", log.Data)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(transferEvent)
// 		case transferBatchSig:
// 			fmt.Printf("Log Name: TransferBatchSig\n")
// 		case ordersMatchedSig:
// 			fmt.Printf("Log Name: OrdersMatched\n")
// 			var transferEvent OrdersMatchedEvent
// 			err := abi.UnpackIntoInterface(&transferEvent, "OrdersMatched", log.Data)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			transferEvent.Maker = common.HexToAddress(log.Topics[1].String())
// 			transferEvent.Taker = common.HexToAddress(log.Topics[2].String())
// 		}
// 	}
// }

// // 解析NFT交易，openSea
// func ParseNFTTx(txHash string) {
// 	hash := common.HexToHash(txHash)
// 	_, isPending, err := ec.TransactionByHash(context.Background(), hash)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("TransactionByHash")
// 	}
// 	if isPending {
// 		log.WithFields(log.Fields{"isPending": isPending}).Error("TransactionByHash")
// 	}
// 	txReceipt, err := ec.TransactionReceipt(context.Background(), hash)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("TransactionReceipt")
// 	}
// 	for _, el := range txReceipt.Logs {
// 		if len(el.Topics) < 4 {
// 			continue
// 		}
// 		// fmt.Printf("log.data: %s\n",  el.Data)
// 		abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
// 		if err != nil {
// 			log.WithFields(log.Fields{"method": "abi.JSON"}).Error(err)
// 		}
// 		switch el.Topics[0].Hex() {
// 		case transferEventSig:
// 			fmt.Printf("Log Name: Transfer\n")
// 			var transferEvent TransferEvent
// 			transferEvent.From = common.HexToAddress(el.Topics[1].String())
// 			transferEvent.To = common.HexToAddress(el.Topics[2].String())
// 			transferEvent.TokenId = el.Topics[3].Big()
// 			fmt.Println(transferEvent)
// 		case transferSingleSig:
// 			fmt.Printf("Log Name: TransferSingle\n")
// 			var transferEvent TransferSingleEvent
// 			err := abi.UnpackIntoInterface(&transferEvent, "TransferSingle", el.Data)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			transferEvent.Operator = common.HexToAddress(el.Topics[1].String())
// 			transferEvent.From = common.HexToAddress(el.Topics[2].String())
// 			transferEvent.To = common.HexToAddress(el.Topics[3].String())
// 			fmt.Println(transferEvent)
// 		case transferBatchSig:
// 			fmt.Printf("Log Name: TransferBatchSig\n")
// 			var transferEvent TransferBatchEvent
// 			err := abi.UnpackIntoInterface(&transferEvent, "TransferBatch", el.Data)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			transferEvent.Operator = common.HexToAddress(el.Topics[1].String())
// 			transferEvent.From = common.HexToAddress(el.Topics[2].String())
// 			transferEvent.To = common.HexToAddress(el.Topics[3].String())
// 			fmt.Println(transferEvent)
// 		case ordersMatchedSig:
// 			fmt.Printf("Log Name: OrdersMatched\n")
// 			var transferEvent OrdersMatchedEvent
// 			err := abi.UnpackIntoInterface(&transferEvent, "OrdersMatched", el.Data)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			transferEvent.Maker = common.HexToAddress(el.Topics[1].String())
// 			transferEvent.Taker = common.HexToAddress(el.Topics[2].String())
// 			transferEvent.Metadata = el.Topics[3]
// 			fmt.Println(transferEvent)
// 		case approvalSig:
// 			fmt.Printf("Log Name: approval\n")
// 			var transferEvent ApprovalEvent
// 			transferEvent.Owner = common.HexToAddress(el.Topics[1].String())
// 			transferEvent.Approved = common.HexToAddress(el.Topics[2].String())
// 			transferEvent.TokenId = el.Topics[3].Big()
// 			fmt.Println(transferEvent)
// 		case approvalForAllSig:
// 			fmt.Printf("Log Name: approval\n")
// 		}
// 	}
// }

// // 解析交易tx.DATA
// func ParserTxInput(txHash string) {
// 	hash := common.HexToHash(txHash)
// 	tx, isPending, err := ec.TransactionByHash(context.Background(), hash)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("TransactionByHash")
// 	}
// 	if isPending {
// 		log.WithFields(log.Fields{"isPending": isPending}).Error("TransactionByHash")
// 	}
// 	abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
// 	if err != nil {
// 		log.WithFields(log.Fields{"method": "abi.JSON"}).Error(err)
// 	}
// 	receivedMap := make(map[string]interface{})
// 	// res, err := DecodeTxParams(abi, receivedMap, tx.Data(), "transferFrom")
// 	res, err := DecodeTxParams(abi, receivedMap, tx.Data(), "atomicMatch_")
// 	// fmt.Println(res["addrs"])
// 	// fmt.Println(res)
// 	// fmt.Println(res["feeMethodsSidesKindsHowToCalls"])
// 	fmt.Println(res["calldataBuy"])
// 	// fmt.Println(res["calldataSell"])
// 	// fmt.Println(res["replacementPatternBuy"])
// 	// fmt.Println(res["replacementPatternSell"])
// 	// fmt.Println(res["staticExtradataBuy"])
// 	// fmt.Println(res["staticExtradataSell"])
// 	// fmt.Println(res["vs"])
// 	// fmt.Println(res["rssMetadata"])
// }

// func DecodeTxParams(abi abi.ABI, v map[string]interface{}, data []byte, name string) (map[string]interface{}, error) {
// 	m, ok := abi.Methods[name]
// 	if !ok {
// 		return nil, fmt.Errorf("error transferFrom method")
// 	}
// 	// tx.Data只包含input部分
// 	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
// 		return nil, err
// 	}
// 	return v, nil
// }

// // 函数签名 对函数名Keccak-256(SHA-3)hash后取前4个字节,转十六进制字符
// // 事件签名 对函数名Keccak-256(SHA-3)hash
// func Signature(method string, isEvent bool) {
// 	hash := crypto.Keccak256Hash([]byte(method))
// 	hashHex := hash.Hex()
// 	if isEvent {
// 		fmt.Println("event " + method + ":" + hashHex)
// 		return
// 	}
// 	fmt.Println(method + ":" + hashHex[0:10]) //"transfer(address,uint256):0xa9059cbb"
// }

// func ShowERC20MethodSignature(erc20Addr string) {
// 	erc20Methers := []string{"0xa9059cbb", "0x70a08231", "0x60806040"}
// 	kec := client.CreateEthClient()
// 	code, err := kec.GetCodeLatest(context.Background(), erc20Addr)
// 	if err != nil {
// 		log.WithFields(log.Fields{"err": err}).Error("GetCodeLatest")
// 	}
// 	for _, ms := range erc20Methers {
// 		b := strings.Contains(code, ms)
// 		if !b {
// 			log.WithFields(log.Fields{"method": ms}).Error("Contains")
// 		}
// 	}
// }

// type Receipt struct {
// 	BlockHash         string `json:"blockHash"`
// 	BlockNumber       string `json:"blockNumber"`
// 	ContractAddress   string `json:"contractAddress"`
// 	CumulativeGasUsed string `json:"cumulativeGasUsed"`
// 	EffectiveGasPrice string `json:"effectiveGasPrice,omitempty"`
// 	GasUsed           string `json:"gasUsed"`
// 	Status            string `json:"status"`
// 	TransactionHash   string `json:"transactionHash"`
// 	TransactionIndex  string `json:"transactionIndex"`
// 	Logs              []Log  `json:"logs"`
// }

// type Block struct {
// 	Transactions []string `json:"transactions"`
// }

// type Log struct {
// 	Address     string   `json:"address"`
// 	BlockHash   string   `json:"blockHash"`
// 	BlockNumber string   `json:"blockNumber"`
// 	Data        string   `json:"data"`
// 	LogIndex    string   `json:"logIndex"`
// 	Removed     bool     `json:"removed"`
// 	Topics      []string `json:"topics"`
// 	TxHash      string   `json:"transactionHash"`
// 	TxIndex     string   `json:"transactionIndex"`
// 	// custom fields
// 	TimeStamp string `json:"timestamp,omitempty"`
// }

// type rpcResponse struct {
// 	ID      int             `json:"id"`
// 	JSONRPC string          `json:"jsonrpc"`
// 	Result  json.RawMessage `json:"result"`
// 	Error   interface{}     `json:"error"`
// }

// // func ParseTx(txHash string) ([]*gokit.ERC20Tx, []*gokit.NFTTransaction, []*gokit.ApprovalLog, []*gokit.ApprovalForAllLog) {
// // 	jsonStr := []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":1}`)
// // 	txJson := []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["` + txHash + `"],"id":1}`)
// // 	result := Fetch(jsonStr)
// // 	var receipt Receipt
// // 	if err := json.Unmarshal([]byte(result), &receipt); err != nil {
// // 		log.WithFields(log.Fields{"method": "GetTransactionReceipt", "params": result}).Panic(err)
// // 	}
// // 	txRes := Fetch(txJson)
// // 	var tx Transaction
// // 	if err := json.Unmarshal([]byte(txRes), &tx); err != nil {
// // 		log.WithFields(log.Fields{"method": "GetTransactionReceipt", "params": result}).Panic(err)
// // 	}
// // 	fmt.Println("txHash:", txHash)
// // 	// erc20TxList, nfttxs, approvals, approvalForAlls := gec.ParseLog(tx.To, tx.Input, TurnToTypeLogs(receipt.Logs))
// // 	// return erc20TxList, nfttxs, approvals, approvalForAlls
// // }

// func Fetch(jsonStr []byte) json.RawMessage {
// 	res, err := http.Post("https://mainnet.infura.io/v3/40b043c639b44d72966d3535d523a4b3", "application/json", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var response rpcResponse
// 	if err := json.Unmarshal(body, &response); err != nil {
// 		fmt.Println(err)
// 	}
// 	return response.Result
// }

// func ListenEvent() {
// 	contractAddress := common.HexToAddress("0xf4272c09B933a2d4Db1d916E282cE1394fc2cd60")
// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{contractAddress},
// 	}
// 	logs := make(chan types.Log)
// 	sub, err := ecw.SubscribeFilterLogs(context.Background(), query, logs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:
// 			fmt.Println(vLog) // pointer to event log
// 		}
// 	}
// }

// // func GetBlockData(blockNumber int64) {
// // 	erc20TxList := make([]*gokit.ERC20Tx, 0)
// // 	nfttxs := make([]*gokit.NFTTransaction, 0)
// // 	approvals := make([]*gokit.ApprovalLog, 0)
// // 	approvalfoAll := make([]*gokit.ApprovalForAllLog, 0)
// // 	number := strconv.FormatInt(blockNumber, 16)
// // 	blockStr := []byte(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x` + number + `",false],"id":1}`)
// // 	result := Fetch(blockStr)
// // 	var block Block
// // 	if err := json.Unmarshal([]byte(result), &block); err != nil {
// // 		log.WithFields(log.Fields{"method": "GetTransactionReceipt", "params": result}).Panic(err)
// // 	}
// // 	for _, txHash := range block.Transactions {
// // 		aa, bb, cc, dd := ParseTx(txHash)
// // 		erc20TxList = append(erc20TxList, aa...)
// // 		nfttxs = append(nfttxs, bb...)
// // 		approvals = append(approvals, cc...)
// // 		approvalfoAll = append(approvalfoAll, dd...)
// // 	}
// // 	fmt.Printf("erc20:%d nfttx: %d appproval:%d approvalfoAll:%d", len(erc20TxList), len(nfttxs), len(approvals), len(approvalfoAll))
// // }
