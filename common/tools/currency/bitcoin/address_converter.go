package bitcoin

import (
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin/bech32m"

	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin/bch"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"

	log "github.com/sirupsen/logrus"
)

const (
	PublicKeyHashAddressPrefix       = "01"
	ScriptHashAddressPrefix          = "02"
	WitnessV0KeyHashAddressPrefix    = "03"
	WitnessV0ScriptHashAddressPrefix = "04"
	WitnessUnknownPrefix             = "05"
	WitnessV1TaprootAddressPrefix    = "06"
)

type Tools struct {
	Coin string
}

func NewTools(coin string) *Tools {
	return &Tools{
		Coin: coin,
	}
}

func (a *Tools) HexToAddress(addr string) string {
	if a.Coin == consts.CoinTypeBitcoinCash {
		return bch.HexToAddress(addr)
	}

	var params chaincfg.Params
	if _, ok := MainNetParams[a.Coin]; ok {
		params = MainNetParams[a.Coin]
	} else {
		params = MainNetParams[consts.CoinTypeBitcoin]
	}

	netID := addr[0:2]
	addressHash, _ := hex.DecodeString(addr[2:])
	switch netID {
	case PublicKeyHashAddressPrefix:
		addr, err := btcutil.NewAddressPubKeyHash(addressHash, &params)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("NewAddressPubKeyHash Address To Hex failed! error: ", err.Error())
			return ""
		}
		return addr.EncodeAddress()
	case ScriptHashAddressPrefix:
		addr, err := btcutil.NewAddressScriptHashFromHash(addressHash, &params)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("NewAddressScriptHashFromHash Address To Hex failed! error: ", err.Error())
			return ""
		}
		return addr.EncodeAddress()
	case WitnessV0KeyHashAddressPrefix:
		addr, err := btcutil.NewAddressWitnessPubKeyHash(addressHash, &params)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("NewAddressWitnessPubKeyHash Address To Hex failed! error: ", err.Error())
			return ""
		}
		return addr.EncodeAddress()
	case WitnessV0ScriptHashAddressPrefix:
		addr, err := btcutil.NewAddressWitnessScriptHash(addressHash, &params)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("NewAddressWitnessScriptHash Address To Hex failed! error: ", err.Error())
			return ""
		}
		return addr.EncodeAddress()
		// case WitnessUnknownPrefix:
	case WitnessUnknownPrefix:
		// bech32 witness_unknown
		version, err := strconv.ParseInt(addr[2:4], 16, 64)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Error(err.Error())
			return ""
		}

		address, err := Bech32Encode(addr[4:], params.Bech32HRPSegwit, version)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("Bech32Encode Address To Hex failed! error: ", err.Error())
			return ""
		}
		return address
	case WitnessV1TaprootAddressPrefix:
		// bech32 witness_v1_taproot
		// 指定 version = 1
		address, err := Bech32Encode(addr[2:], params.Bech32HRPSegwit, 1)
		if err != nil {
			log.WithFields(log.Fields{"func": "HexToAddress", "address": addr}).
				Warn("Bech32Encode Address To Hex failed! error: ", err.Error())
			return ""
		}
		return address
	default:
		// implements
	}
	return ""
}

// PKScript2Addr pkscript to address and type
func (a *Tools) PKScript2Addr(pkScript string) (encodedAddrs []string, addressType string) {
	var params chaincfg.Params
	if _, ok := MainNetParams[a.Coin]; ok {
		params = MainNetParams[a.Coin]
	} else {
		params = MainNetParams[consts.CoinTypeBitcoin]
	}

	scriptClass, addrs, _, _ := txscript.ExtractPkScriptAddrs(
		[]byte(pkScript), &params)
	encodedAddrs = make([]string, len(addrs))
	for j, addr := range addrs {
		encodedAddr := addr.EncodeAddress()
		encodedAddrs[j] = encodedAddr
	}
	addressType = scriptClass.String()
	// NullDataTy, NonStandardTy, segwit v1
	if len(addrs) == 0 {
		witver, err := strconv.Atoi(hex.EncodeToString([]byte(pkScript))[1:2])
		if err != nil {
			return
		}
		a, err := bech32m.SegwitAddrEncode(params.Bech32HRPSegwit, byte(witver), []byte(pkScript)[2:])
		if err != nil {
			return
		}
		// temp: witness_v1_taproot
		addressType = "witness_v1_taproot"
		return []string{a}, addressType
	}
	return
}

// AddressToHex address to 地址公钥 hash160，压缩格式转的也是压缩格式
func (a *Tools) AddressToHex(addr string) (string, error) {
	if addr == "" {
		return "", nil
	}
	if a.Coin == consts.CoinTypeBitcoinCash {
		return bch.AddressToHex(addr)
	}
	params := MainNetParams[consts.CoinTypeBitcoin]
	if _, ok := MainNetParams[a.Coin]; ok {
		params = MainNetParams[a.Coin]
	}
	// bech32m address
	ver, data, err := bech32m.SegwitAddrDecode(params.Bech32HRPSegwit, addr)
	if ver >= byte(1) && len(data) == 32 && err == nil {
		return hex.EncodeToString(data), nil
	}
	// btcutil DecodeAddress only support version 0 for witness program
	addressHash, err := btcutil.DecodeAddress(addr, &params)
	if err != nil {
		log.WithFields(log.Fields{"func": "AddressToHex", "address": addr}).
			Warn("DecodeAddress Address To Hex failed! error: ", err.Error())
		return "", err
	}
	return hex.EncodeToString(addressHash.ScriptAddress()), nil
}

func (a *Tools) GetAddressType(addr string) (addrType string, err error) {
	params := MainNetParams[consts.CoinTypeBitcoin]
	if _, ok := MainNetParams[a.Coin]; ok {
		params = MainNetParams[a.Coin]
	}
	// check if addr is bech32 format
	match1, _ := regexp.MatchString("^bc1[0-9a-zA-Z]{11,71}$", strings.ToLower(addr))
	match2, _ := regexp.MatchString("^tb1[0-9a-zA-Z]{11,71}$", strings.ToLower(addr))
	match3, _ := regexp.MatchString("^ltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(addr))
	match4, _ := regexp.MatchString("^tltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(addr))
	match5, _ := regexp.MatchString("^royale1[0-9a-zA-Z]{11,71}$", strings.ToLower(addr))
	if match2 || match3 || match4 || match5 || match1 {
		// support bech32m address
		ver, res, _ := bech32m.SegwitAddrDecode(params.Bech32HRPSegwit, addr)
		if ver >= byte(1) && len(res) == 32 {
			return consts.BitcoinAddressTypeWitnessV1Taproot, nil
		}
		_, addrType, _, err = Bech32Decode(addr)
		if err != nil {
			log.WithFields(log.Fields{"func": "GetAddressType", "address": addr}).
				Warn("Bech32Decode Address To Hex failed! error: ", err.Error())
			return "", err
		}
		return addrType, nil
	}

	decoded, err := btcutil.DecodeAddress(addr, &params)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAddressType", "address": addr}).
			Warn("DecodeAddress Address To Hex failed! error: ", err.Error())
		return "", err
	}
	switch decoded.(type) {
	case *btcutil.AddressPubKeyHash:
		addrType = consts.BitcoinAddressTypeP2PKH
	case *btcutil.AddressScriptHash:
		addrType = consts.BitcoinAddressTypeP2SH
	case *btcutil.AddressPubKey:
		addrType = consts.BitcoinAddressTypePubkey
	case *btcutil.AddressWitnessPubKeyHash:
		addrType = consts.BitcoinAddressTypeP2WPKHV0
	case *btcutil.AddressWitnessScriptHash:
		addrType = consts.BitcoinAddressTypeP2WSHV0
	}

	return addrType, nil
}
