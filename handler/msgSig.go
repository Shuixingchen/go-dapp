package handler

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	/*
	 * bytes4(keccak256('supportsInterface(bytes4)')) == 0x01ffc9a7
	 */
	_INTERFACE_ID_ERC165 = "0x01ffc9a7"

	// Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
	// which can be also obtained as `IERC721Receiver(0).onERC721Received.selector`
	_ERC721_RECEIVED = "0x150b7a02"

	/*
	 *     bytes4(keccak256('balanceOf(address)')) == 0x70a08231
	 *     bytes4(keccak256('ownerOf(uint256)')) == 0x6352211e
	 *     bytes4(keccak256('approve(address,uint256)')) == 0x095ea7b3
	 *     bytes4(keccak256('getApproved(uint256)')) == 0x081812fc
	 *     bytes4(keccak256('setApprovalForAll(address,bool)')) == 0xa22cb465
	 *     bytes4(keccak256('isApprovedForAll(address,address)')) == 0xe985e9c5
	 *     bytes4(keccak256('transferFrom(address,address,uint256)')) == 0x23b872dd
	 *     bytes4(keccak256('safeTransferFrom(address,address,uint256)')) == 0x42842e0e
	 *     bytes4(keccak256('safeTransferFrom(address,address,uint256,bytes)')) == 0xb88d4fde
	 *     => 0x70a08231 ^ 0x6352211e ^ 0x095ea7b3 ^ 0x081812fc ^
	 *        0xa22cb465 ^ 0xe985e9c5 ^ 0x23b872dd ^ 0x42842e0e ^ 0xb88d4fde == 0x80ac58cd
	 */
	_INTERFACE_ID_ERC721 = "0x80ac58cd"

	/*
	 *     bytes4(keccak256('name()')) == 0x06fdde03
	 *     bytes4(keccak256('symbol()')) == 0x95d89b41
	 *     bytes4(keccak256('tokenURI(uint256)')) == 0xc87b56dd
	 *     => 0x06fdde03 ^ 0x95d89b41 ^ 0xc87b56dd == 0x5b5e139f
	 */
	_INTERFACE_ID_ERC721_METADATA = "0x5b5e139f"

	/*
	 *     bytes4(keccak256('totalSupply()')) == 0x18160ddd
	 *     bytes4(keccak256('tokenOfOwnerByIndex(address,uint256)')) == 0x2f745c59
	 *     bytes4(keccak256('tokenByIndex(uint256)')) == 0x4f6ccce7
	 *     => 0x18160ddd ^ 0x2f745c59 ^ 0x4f6ccce7 == 0x780e9d63
	 */
	_INTERFACE_ID_ERC721_ENUMERABLE = "0x780e9d63"
)

// ECDSA验证签名
// 前提条件：签名，原始数据以及签名者的公钥/地址
func VerifySig(msg, signature string) bool {
	// 1.准备公钥
	privateKey, err := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// 2.原始数据的hash, hash算法要与签名时的一致
	msgHash := crypto.Keccak256Hash([]byte(msg))

	// 3. 调用Ecrecover（椭圆曲线签名恢复）来检索签名者的公钥
	sig, _ := hex.DecodeString(signature)
	sigPublicKey, err := crypto.Ecrecover(msgHash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sigPublicKey:", hex.EncodeToString(sigPublicKey))
	// fmt.Println("publicKeyBytes:", hex.EncodeToString(publicKeyBytes))
	// 4.比较公钥是否一致
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println("VerifySig", matches)
	return matches
}

// ECDSA验证签名
// 前提条件，签名，原始数据，签名者地址
func VerifySig2(msg, signature, addr string) bool {

	// 1. 调用Ecrecover（椭圆曲线签名恢复）来检索签名者的公钥
	data := []byte(msg)
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	msgHash := crypto.Keccak256Hash([]byte(msg))
	sig, _ := hex.DecodeString(signature)
	sigPublicKey, err := crypto.SigToPub(msgHash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	// 2.公钥得到地址
	//使用Keccak256函数手动完成地址解析
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(sigPublicKey[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

	signAddr := crypto.PubkeyToAddress(*sigPublicKey).Hex()

	// 3.对比两个地址是否一致
	// matches := strings.EqualFold(addr, signAddr)
	fmt.Println("addr", addr, "signAddr", signAddr)
	// fmt.Println("VerifySig2", matches)
	return false
}

// ECDSA签名
func SignMessage(msg string) string {
	// 1.准备私钥
	privateKey, err := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if err != nil {
		log.Fatal(err)
	}
	// 2.先对数据进行hash,再用私钥签名,会得到65字节数据，
	// 3.大部分钱包签名的时候会加个前缀
	data := []byte(msg)
	msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	msgHash := crypto.Keccak256Hash([]byte(msg))
	sig, err := crypto.Sign(msgHash[:], privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sig length:", len(sig))
	fmt.Println("sig hex:", hex.EncodeToString(sig))
	return hex.EncodeToString(sig)
}

func SignAccount(msg string) string {
	// 1.准备私钥
	privateKey, err := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if err != nil {
		log.Fatal(err)
	}
	msgHash, _ := accounts.TextAndHash([]byte(msg))
	sig, err := crypto.Sign(msgHash[:], privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sig length:", len(sig))
	fmt.Println("sig hex:", hex.EncodeToString(sig))
	return hex.EncodeToString(sig)
}

func VerifyHandler() {
	message := "Hello World!"
	addr := "0xe725D38CC421dF145fEFf6eB9Ec31602f95D8097"
	signature := SignAccount(message)
	VerifySig2(message, signature, addr)
	// VerifySig(message, signature)
}

// pubkey to address
func PubkeyToAddress() {
	pub, err := hex.DecodeString("04a83030e54912f617b4ff35a85cc7ffde9ba2de3f8aa32ddc9c9a39f1df5714fc624958470bc53cf65b179e045bfbb3d5a784b982d064db2a2d4480d5d22c6ddb")
	if err != nil {
		fmt.Println(err, "hex.DecodeString")
	}
	escaPub, err := crypto.UnmarshalPubkey(pub)
	if err != nil {
		fmt.Println(err, "UnmarshalPubkey")
	}
	addr := crypto.PubkeyToAddress(*escaPub)
	fmt.Println(addr.Hex())
}

func KeyShow() {
	privateKey, _ := crypto.HexToECDSA("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")

	fmt.Println("public key no 0x \n", hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)))
}
