package currency

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin/bech32m"

	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin"
	"github.com/btcsuite/btcd/btcjson"
	log "github.com/sirupsen/logrus"
)

const (
	PublicKeyHashAddressPrefix       = "01"
	ScriptHashAddressPrefix          = "02"
	WitnessV0KeyHashAddressPrefix    = "03"
	WitnessV0ScriptHashAddressPrefix = "04"
	WitnessUnknownAddressPrefix      = "05"
	WitnessV1TaprootAddressPrefix    = "06"
)

var TaprootV1AddressBeforeActive = []string{
	"bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9",
	"bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqsyjer9e",
	"bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh",
	"bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqs3wf0qm",
	"bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5ss52r5n8",
}

var DecodeTxAddressType = map[string]string{
	"pubkeyhash":            "P2PKH",
	"scripthash":            "P2SH",
	"witness_v0_keyhash":    "P2WPKH_V0",
	"witness_v0_scripthash": "P2WSH_V0",
	"pubkey":                "P2PKH_PUBKEY",
	"nonstandard":           "NONSTANDARD",
	"nulldata":              "NULL_DATA",
	"multisig":              "P2PKH_MULTISIG",
	"witness_v1_taproot":    "WITNESS_V1_TAPROOT",
	"witness_unknown":       "WITNESS_UNKNOWN",
}

func AddressTypeConvert(str string) string {
	addressType, ok := DecodeTxAddressType[str]
	if !ok {
		addressType = "NULL_DATA"
	}
	return addressType
}

func ReverseTxHashString(s string) string {
	var str2 []string
	i := 0
	for i < len(s) {
		a := s[i : i+2]
		str2 = append(str2, a)
		i += 2
	}
	for from, to := 0, len(str2)-1; from < to; from, to = from+1, to-1 {
		str2[from], str2[to] = str2[to], str2[from]
	}
	return strings.Join(str2, "")
}

func GetAddressID(coin, address, addressType string) (addressID string, err error) {
	// if address match bech32, replace the address's hrp
	if match, hrp := IsBech32Adddress(address); match {
		address = hrp + address[len(hrp):]
	}
	var addressHex string
	// todo: taproot激活前，地址：
	// [
	// bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9,
	// bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqsyjer9e,
	// ]
	// 被识别为 WITNESS_V1_TAPROOT, 但是脚本哈希也是 0501 开头，这种情况需要单独区分出来。
	if addressType == consts.BitcoinAddressTypeWitnessUnknown ||
		(addressType == consts.BitcoinAddressTypeWitnessV1Taproot &&
			strings.Contains(strings.Join(TaprootV1AddressBeforeActive, ","), address)) {
		// need support bech32m
		version, _, addrHash, err := bitcoin.Bech32Decode(address)
		if err != nil {
			log.WithFields(log.Fields{"func": "GetAddressID", "coin": coin, "address": address}).
				Warn("Bech32Decode convert Address To Hex failed! error: ", err.Error())
			return "", err
		}
		combined := make([]byte, len(addrHash)+2)
		// consts.BitcoinAddressTypeWitnessUnknown
		combined[0] = 5
		combined[1] = version
		copy(combined[2:], addrHash)
		addressHex = hex.EncodeToString(combined)
	} else {
		tools := bitcoin.NewTools(coin)
		hash, err := tools.AddressToHex(address)
		if err != nil {
			log.WithFields(log.Fields{"func": "GetAddressID", "coin": coin, "address": address}).
				Warn("convert Address To Hex failed! error: ", err.Error())
			return "", err
		}

		var prefix string
		switch addressType {
		case consts.BitcoinAddressTypeP2PKH:
			prefix = PublicKeyHashAddressPrefix
		case consts.BitcoinAddressTypeP2SH:
			prefix = ScriptHashAddressPrefix
		case consts.BitcoinAddressTypeP2WPKHV0:
			prefix = WitnessV0KeyHashAddressPrefix
		case consts.BitcoinAddressTypeP2WSHV0:
			prefix = WitnessV0ScriptHashAddressPrefix
		case consts.BitcoinAddressTypeWitnessUnknown:
			prefix = WitnessUnknownAddressPrefix
		case consts.BitcoinAddressTypeWitnessV1Taproot:
			prefix = WitnessV1TaprootAddressPrefix
		}

		addressHex = prefix + hash
	}

	addressID = strings.ToLower(addressHex)

	return addressID, nil
}

func GuessType(address string) string {
	match1, _ := regexp.MatchString("^[1-9A-Za-z]{26,35}$", address)
	if match1 {
		if address[0:1] == "1" || address[0:1] == "m" || address[0:1] == "n" || address[0:1] == "L" || address[0:1] == "Y" {
			return consts.BitcoinAddressTypeP2PKH
		}
		if address[0:1] == "3" || address[0:1] == "2" || address[0:1] == "M" || address[0:1] == "R" {
			return consts.BitcoinAddressTypeP2SH
		}
	}
	cashAddr, _ := regexp.MatchString("^bitcoincash:[0-9a-zA-Z]{42}$", address)
	if cashAddr {
		if address[12:13] == "q" {
			return consts.BitcoinAddressTypeP2PKH
		} else if address[12:13] == "p" {
			return consts.BitcoinAddressTypeP2SH
		} else {
			return ""
		}
	}

	cashAddrMatch1, _ := regexp.MatchString("^q[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	if cashAddrMatch1 {
		return consts.BitcoinAddressTypeP2PKH
	}
	cashAddrMatch2, _ := regexp.MatchString("^p[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	if cashAddrMatch2 {
		return consts.BitcoinAddressTypeP2SH
	}

	if len(address) == 40 {
		return consts.BitcoinAddressTypeP2WPKHV0
	}

	if len(address) == 64 {
		return consts.BitcoinAddressTypeP2WSHV0
	}

	if match2, hrp := IsBech32Adddress(address); match2 {
		address = hrp + address[len(hrp):]
		ver, res, _ := bech32m.SegwitAddrDecode(hrp, address)
		if ver >= byte(1) && len(res) == 32 {
			return consts.BitcoinAddressTypeWitnessV1Taproot
		}
		_, addrType, _, err := bitcoin.Bech32Decode(address)
		if err != nil {
			log.WithFields(log.Fields{"func": "GuessType", "address": address}).
				Warn("Bech32Decode Address To Hex failed! error: ", err.Error())
			return ""
		}
		return addrType
	}

	return ""
}

func IsBech32Adddress(address string) (match bool, hrp string) {
	match1, _ := regexp.MatchString("^bc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match2, _ := regexp.MatchString("^tb1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match3, _ := regexp.MatchString("^ltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match4, _ := regexp.MatchString("^tltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match5, _ := regexp.MatchString("^royale1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))

	match = match1 || match2 || match3 || match4 || match5

	switch {
	case match1:
		hrp = "bc"
	case match2:
		hrp = "tb"
	case match3:
		hrp = "ltc"
	case match4:
		hrp = "tltc"
	case match5:
		hrp = "royale"
	}

	return
}

func GuessAddressCoinType(address string) string {
	match1, _ := regexp.MatchString("^[1-9A-Za-z]{26,35}$", address)
	if match1 {
		if address[0:1] == "1" || address[0:1] == "m" || address[0:1] == "n" || address[0:1] == "3" || address[0:1] == "2" {
			return consts.CoinTypeBitcoin
		}
		if address[0:1] == "L" || address[0:1] == "M" {
			return consts.CoinTypeLitecoin
		}
	}
	cashAddr, _ := regexp.MatchString("^bitcoincash:[0-9a-zA-Z]{42}$", address)
	if cashAddr {
		return consts.CoinTypeBitcoinCash
	}

	cashAddrMatch1, _ := regexp.MatchString("^q[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	if cashAddrMatch1 {
		return consts.CoinTypeBitcoinCash
	}
	cashAddrMatch2, _ := regexp.MatchString("^p[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	if cashAddrMatch2 {
		return consts.CoinTypeBitcoinCash
	}

	match2, _ := regexp.MatchString("^bc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match3, _ := regexp.MatchString("^tb1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	if match2 || match3 {
		return consts.CoinTypeBitcoin
	}

	match4, _ := regexp.MatchString("^ltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	match5, _ := regexp.MatchString("^tltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	if match4 || match5 {
		return consts.CoinTypeLitecoin
	}

	return ""
}

func HexToAddress(coin, addrHash string) string {
	tools := bitcoin.NewTools(coin)
	return tools.HexToAddress(addrHash)
}

func V3HexToAddress(coin, addrHash string) string {
	prefix := addrHash[0:2]
	hash := addrHash[2:]
	var addrLen int
	switch prefix {
	case PublicKeyHashAddressPrefix:
		addrLen = 40
	case ScriptHashAddressPrefix:
		addrLen = 40
	case WitnessV0KeyHashAddressPrefix:
		addrLen = 40
	case WitnessV0ScriptHashAddressPrefix:
		addrLen = 64
	default:
		addrLen = 40
	}
	addr := prefix + hash[0:addrLen]
	tools := bitcoin.NewTools(coin)
	return tools.HexToAddress(addr)
}

func PKScript2Addr(coin, pkScript string) (encodedAddrs []string, addressType string) {
	tools := bitcoin.NewTools(coin)
	return tools.PKScript2Addr(pkScript)
}

func DecodeTx(coin, rawTx string) (result *btcjson.TxRawDecodeResult, witnessHash string, isSwTx bool, err error) {
	tools := bitcoin.NewTools(coin)
	return tools.DecodeTx(rawTx)
}

func GetAddressType(coin, addr string) (addrType string, err error) {
	tools := bitcoin.NewTools(coin)
	return tools.GetAddressType(addr)
}

// CutCashAddrPrefix 去掉 CashAddr Prefix
func CutCashAddrPrefix(coin, addr string) string {
	var result string
	if coin == consts.CoinTypeBitcoinCash {
		result = strings.TrimLeft(addr, consts.CashAddrPrefix)
	} else {
		result = addr
	}

	return result
}

// IsValidEthereumAddress 判断ETH地址格式
func IsValidEthereumAddress(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}
