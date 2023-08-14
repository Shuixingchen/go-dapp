package bch

import (
	"encoding/hex"

	"github.com/bcext/cashutil"
	"github.com/bcext/gcash/chaincfg"
)

func HexToAddress(addr string) string {
	params := chaincfg.MainNetParams
	netID := addr[0:2]
	addressHash, _ := hex.DecodeString(addr[2:])
	switch netID {
	case "01":
		addr, _ := cashutil.NewAddressPubKeyHash(addressHash, &params)
		return addr.EncodeAddress(true)
	case "02":
		addr, _ := cashutil.NewAddressScriptHashFromHash(addressHash, &params)
		return addr.EncodeAddress(true)
	default:
	}

	return ""
}

func AddressToHex(addr string) (string, error) {
	params := chaincfg.MainNetParams
	addressHash, err := cashutil.DecodeAddress(addr, &params)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(addressHash.ScriptAddress()), nil
}
