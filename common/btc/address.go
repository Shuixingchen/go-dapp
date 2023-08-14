package btc

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/golang/glog"
)

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

// 通过锁定脚本解锁地址信息
func ParsePkScript(pkScript []byte) ScriptPubKey {
	var s ScriptPubKey
	scriptClass, addrs, _, err := txscript.ExtractPkScriptAddrs(pkScript, &chaincfg.MainNetParams)
	if err != nil {
		glog.Infof("ExtractPkScriptAddrs faild err = %v", err)
	} else {
		if len(addrs) == 1 { // 多签地址不处理
			s.Address = addrs[0].EncodeAddress()
		}
	}
	s.Type = scriptClass.String()
	s.Asm, _ = txscript.DisasmString(pkScript[:]) // 翻编译脚本，得到字节码
	s.Hex = hex.EncodeToString(pkScript)          // 脚本转为hex
	return s
}
