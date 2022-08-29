package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Shuixingchen/go-dapp/utils"
	log "github.com/sirupsen/logrus"
)

// HexUint64 is an alias for uint64
type HexUint64 uint64

// UnmarshalJSON implements the json unmarshal behavior for HexUint64.
func (i *HexUint64) UnmarshalJSON(data []byte) error {
	src := string(bytes.Trim(data, `"`))
	// ignore empty case
	if src == "" || strings.ToLower(src) == "null" {
		return nil
	}

	u, err := utils.ParseUint(src)
	*i = HexUint64(u)
	return err
}

// MarshalJSON implements the json marshal behavior for HexUint64.
func (i *HexUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.Uint64ToHex(uint64(*i)))
}

func (i *HexUint64) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

// ToHex converts HexUint64 to hex string.
func (i *HexUint64) ToHex() string {
	return utils.Uint64ToHex(uint64(*i))
}

// Scan implements sql scan interface.
func (i *HexUint64) Scan(src interface{}) error {

	switch src := src.(type) {
	case time.Time:
		*i = HexUint64(src.Unix())
	case int64:
		*i = HexUint64(uint64(src))
	case string:
		u, err := StringToHexUint64(src)
		if err != nil {
			return err
		}
		*i = u
	case []uint8:
		bs := src
		ba := make([]byte, 0, len(bs))
		for _, b := range bs {
			ba = append(ba, byte(b))
		}
		u, err := StringToHexUint64(string(ba))
		if err != nil {
			return err
		}
		*i = u
	default:
		return fmt.Errorf("Incompatible type (%v) for HexUint64", reflect.TypeOf(src))
	}
	return nil
}

// StringToHexUint64 converts a string base 10 to HexUint64 format
func StringToHexUint64(src string) (HexUint64, error) {
	var res HexUint64
	u, err := strconv.ParseUint(src, 10, 64)
	if err != nil {
		return res, err
	}
	res = HexUint64(u)
	return res, nil
}

// HexStringToHexUint64 converts a base64 string to hexUint64 format
func HexStringToHexUint64(src string) (HexUint64, error) {
	val, err := utils.ParseUint(src)
	if err == nil {
		return HexUint64(val), nil
	}
	log.WithFields(log.Fields{"method:": "HexStringToHexUint64", "param:": src}).Error(err)
	return HexUint64(0), err
}

// HexBig is an alias for big.Int.
type HexBig struct {
	*big.Int
}

// NewHexBig returns a new pointer to a HexBig instance.
func NewHexBig() *HexBig {
	return &HexBig{new(big.Int)}
}

// UnmarshalJSON implements the unmarshal json function
func (i *HexBig) UnmarshalJSON(data []byte) error {
	u, err := utils.ParseBigInt(string(bytes.Trim(data, `"`)))
	i.Int = u
	return err
}

// MarshalJSON implements the marshal json function
func (i *HexBig) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.BigToHex(i.Int))
}

// Scan implements sql scan interface.
func (i *HexBig) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		u, err := StringToHexBig(string(src))
		if err != nil {
			return err
		}
		*i = *u
	case int:
		i.Int = big.NewInt(int64(src))
	case int64:
		i.Int = big.NewInt(src)
	case string:
		u, err := StringToHexBig(src)
		if err != nil {
			return err
		}
		*i = *u
	default:
		fmt.Println("Got some unrecognized error here")
		return fmt.Errorf("Incompatible type (%v) for HexBig", reflect.TypeOf(src))
	}
	return nil
}

// StringToHexBig converts string to hexbig type
func StringToHexBig(src string) (*HexBig, error) {
	res := &HexBig{new(big.Int)}
	i, ok := res.SetString(src, 10)
	if !ok {
		return nil, fmt.Errorf("Invalid string value to big.Int: %q", src)
	}
	res.Int = i
	return res, nil
}

func HexStringToHexBig(src string) (*HexBig, error) {
	res := &HexBig{new(big.Int)}
	i, ok := res.SetString(strings.TrimPrefix(src, "0x"), 16)
	if !ok {
		log.WithFields(log.Fields{"method:": "HexStringToHexBig", "param:": src}).Error(ok)
		return nil, fmt.Errorf("Invalid string value to big.Int: %q", src)
	}
	res.Int = i
	return res, nil
}
