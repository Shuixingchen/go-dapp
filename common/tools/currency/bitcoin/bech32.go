package bitcoin

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin/bech32m"

	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/btcsuite/btcd/btcutil/bech32"
)

// Bech32Decode due to btcsuite/utils pkg can't decode witness_unknown address
func Bech32Decode(address string) (version byte, addrType string, addrHash []byte, err error) {
	_, data, spec, err := bech32m.Decode(address)
	if err != nil {
		return byte(0), "", nil, err
	}
	if len(data) < 1 {
		return byte(0), "", nil, fmt.Errorf("empty data section")
	}
	version = data[0]
	if version >= byte(1) && spec == bech32m.Bech32m {
		addrType = consts.BitcoinAddressTypeWitnessV1Taproot
		addrHash, err = bech32m.Convertbits(data[1:], 5, 8, false)
		if len(addrHash) != 32 {
			addrType = consts.BitcoinAddressTypeWitnessUnknown
		}
		return data[0], addrType, addrHash, err
	}

	_, decoded, err := bech32.Decode(address)
	if err != nil {
		return version, "", addrHash, err
	}

	addrHash, err = bech32.ConvertBits(decoded[1:], 5, 8, false)
	if err != nil {
		return version, "", addrHash, err
	}

	addrType = consts.BitcoinAddressTypeWitnessUnknown
	// https://bitcoincore.org/en/segwit_wallet_dev/
	if version == byte(0) {
		switch len(addrHash) {
		case 20:
			addrType = consts.BitcoinAddressTypeP2WPKHV0
		case 32:
			addrType = consts.BitcoinAddressTypeP2WSHV0
		default:
			addrType = consts.BitcoinAddressTypeWitnessUnknown
		}
	}

	return version, addrType, addrHash, nil
}

// Bech32Encode due to btcsuite/utils pkg can't encode witness_unknown address
func Bech32Encode(addrHash, hrp string, version int64) (address string, err error) {
	decoded, err := hex.DecodeString(strings.ToLower(addrHash))
	if err != nil {
		return "", err
	}
	if version == 1 {
		address, err = bech32m.SegwitAddrEncode(hrp, byte(version), decoded)
		if err != nil {
			return "", err
		}
		return address, nil
	}
	conv, err := bech32.ConvertBits(decoded, 8, 5, true)
	if err != nil {
		return "", err
	}
	// Concatenate the witness version and program, and encode the resulting
	// bytes using bech32 encoding.
	combined := make([]byte, len(conv)+1)
	combined[0] = byte(version)
	copy(combined[1:], conv)
	encoded, err := bech32.Encode(hrp, combined)
	if err != nil {
		return "", err
	}

	return encoded, nil
}
