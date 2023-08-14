package blockchain

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Shuixingchen/go-dapp/common/tools/currency"
)

// IsBlockHashOrTransactionHash 判断是否是区块 hash 或者是交易 hash
func IsBlockHashOrTransactionHash(keyword string) bool {
	// 先判断是不是 BTC,BCH,LTC 区块哈希或交易哈希
	re := regexp.MustCompile("^[0-9a-fA-F]{64}$")
	if re.MatchString(keyword) {
		return true
	}
	// 再判断是不是 ETH 区块哈希或交易哈希
	return IsETHBlockHashOrTransactionHash(keyword)
}

// IsETHBlockHashOrTransactionHash 判断是否是 ETH 区块 hash 或者是交易 hash
func IsETHBlockHashOrTransactionHash(keyword string) bool {
	re := regexp.MustCompile("^0x([A-Fa-f0-9]{64})$")
	if re.MatchString(keyword) {
		return true
	} else if len(keyword) == 48 && strings.HasPrefix(keyword, "GENESIS") {
		return true
	}

	return false
}

// IsBTCBlockHash 判断是否是 BTC 区块 hash
func IsBTCBlockHash(keyword string) bool {
	re := regexp.MustCompile("^0{8}[0-9a-fA-F]{56}$")
	return re.MatchString(keyword)
}

// IsBTCTransactionHash 判断是否是 BTC 交易 hash
func IsBTCTransactionHash(keyword string) bool {
	re := regexp.MustCompile("^[0-9a-fA-F]{64}$")
	return re.MatchString(keyword)
}

// IsAddress 判断是否是地址
func IsAddress(address string) bool {
	// 先判断是不是 BTC,BCH,LTC 地址
	match, _ := regexp.MatchString("^[1-9A-Za-z]{26,35}$", address)
	if match {
		return true
	}

	btcSegwitAddressCheck, _ := regexp.MatchString("^bc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	ltcSegwitAddressCheck, _ := regexp.MatchString("^ltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	c1, _ := regexp.MatchString("^bitcoincash:[0-9a-zA-Z]{42}$", strings.ToLower(address))
	c2, _ := regexp.MatchString("^q[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	c3, _ := regexp.MatchString("^p[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	cashAddressCheck := c1 || c2 || c3

	return btcSegwitAddressCheck || ltcSegwitAddressCheck || cashAddressCheck || IsETHAddress(address)
}

// IsETHAddress 判断是否是 ETH 地址
func IsETHAddress(address string) bool {
	matched1, _ := regexp.MatchString("^0x([A-Fa-f0-9]{40})$", strings.ToLower(address))
	matched2, _ := regexp.MatchString("^0xGENESIS([A-Fa-f0-9]{33})$", address)
	return matched1 || matched2
}

// IsLegalAddress 判断是否为合法地址
func IsLegalAddress(coin, address string) (match bool, addressID string) {
	addressType := currency.GuessType(address)
	if addressType == "" {
		match = false
		return
	}
	addressID, err := currency.GetAddressID(coin, address, addressType)
	if err != nil || addressID == "" {
		match = false
		return
	}
	match = true
	return match, addressID
}

// IsLegalBlockInput 检查block height/hash格式
func IsLegalBlockInput(keyWord string) (match bool, format string) {
	if _, err := strconv.ParseInt(keyWord, 10, 64); err == nil {
		match = true
		format = "height"
		return
	}

	if IsBlockHashOrTransactionHash(keyWord) {
		match = true
		format = "hash"
		return
	}

	if keyWord == "latest" {
		match = true
		format = keyWord
		return
	}

	return match, format
}

// InputCheck 检查输入
func InputCheck(keyWord string) bool {
	height, err := strconv.ParseInt(keyWord, 10, 64)
	if err == nil && height >= 0 {
		return true
	}

	if IsBlockHashOrTransactionHash(keyWord) || IsAddress(keyWord) || IsETHTokenSymbolName(keyWord) {
		return true
	}

	return false
}

func IsCashAddr(address string) bool {
	cashAddr, _ := regexp.MatchString("^bitcoincash:[0-9a-zA-Z]{42}$", address)
	if cashAddr && (address[12:13] == "q" || address[12:13] == "p") {
		return true
	}

	cashAddrMatch1, _ := regexp.MatchString("^q[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	if cashAddrMatch1 {
		return true
	}
	cashAddrMatch2, _ := regexp.MatchString("^p[0-9a-zA-Z]{30,50}$", strings.ToLower(address))
	return cashAddrMatch2
}

func IsLTCAddress(address string) bool {
	match1, _ := regexp.MatchString("^[1-9A-Za-z]{26,35}$", address)
	if match1 && (address[0:1] == "L" || address[0:1] == "M") {
		return true
	}

	match2, _ := regexp.MatchString("^ltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	if match2 {
		return true
	}

	match3, _ := regexp.MatchString("^tltc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	return match3
}

func IsBC1Address(address string) bool {
	matched, _ := regexp.MatchString("^bc1[0-9a-zA-Z]{11,71}$", strings.ToLower(address))
	return matched
}

// IsETHTokenSymbolName 暂时过滤掉一些特殊字符
func IsETHTokenSymbolName(name string) bool {
	matched, _ := regexp.MatchString("^[\u4E00-\u9FA5A-Za-z0-9. +_-]+$", name)
	return matched
}
