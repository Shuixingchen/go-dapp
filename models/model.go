package models

import "math/big"

type Token struct {
	Addr            string   `json:"tokenAddress"`
	Name            string   `json:"name"`
	Symbol          string   `json:"sym"`
	Decimals        int64    `json:"decimals"`
	InitTotalSupply *big.Int `json:"initTotalSupply"` // the maximum supply of a ERC20 token can be 2^256 - 1
	Type            string   `json:"tokenType"`
}
