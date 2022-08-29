package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// ParseUint parse hex string value to uint64.
func ParseUint(value string) (uint64, error) {
	if value == "" || !strings.HasPrefix(value, "0x") {
		return 0, nil
	}
	i, err := strconv.ParseUint(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Uint64ToHex convets uint64 into hexadecimal representation.
func Uint64ToHex(i uint64) string {
	return fmt.Sprintf("0x%x", i)
}

// ParseBigInt parse hex string value to big.Int.
func ParseBigInt(value string) (*big.Int, error) {
	if value == "0x" {
		value = "0x0"
	}
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)
	return &i, err
}

// BigToHex covert big.Int to hexadecimal representation.
func BigToHex(bigInt *big.Int) string {
	res := fmt.Sprintf("%x", bigInt.Bytes())
	if res == "" {
		return "0x0"
	}
	return "0x" + res
}
