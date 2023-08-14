package bitcoin

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/Shuixingchen/go-dapp/common/tools/currency/bitcoin/bch"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/mempool"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// witnessToHex formats the passed witness stack as a slice of hex-encoded
// strings to be used in a JSON response.
func witnessToHex(witness wire.TxWitness) []string {
	// Ensure nil is returned when there are no entries versus an empty
	// slice so it can properly be omitted as necessary.
	if len(witness) == 0 {
		return nil
	}

	result := make([]string, 0, len(witness))
	for _, wit := range witness {
		result = append(result, hex.EncodeToString(wit))
	}

	return result
}

// createVoutList returns a slice of JSON objects for the outputs of the passed
// transaction.
func createVoutList(mtx *wire.MsgTx, chainParams *chaincfg.Params, filterAddrMap map[string]struct{}) []btcjson.Vout {
	voutList := make([]btcjson.Vout, 0, len(mtx.TxOut))
	for i, v := range mtx.TxOut {
		// The disassembled string will contain [error] inline if the
		// script doesn't fully parse, so ignore the error here.
		disbuf, _ := txscript.DisasmString(v.PkScript)

		// Ignore the error here since an error means the script
		// couldn't parse and there is no additional information about
		// it anyways.
		scriptClass, addrs, reqSigs, _ := txscript.ExtractPkScriptAddrs(
			v.PkScript, chainParams)

		// Encode the addresses while checking if the address passes the
		// filter when needed.
		passesFilter := len(filterAddrMap) == 0
		encodedAddrs := make([]string, len(addrs))
		for j, addr := range addrs {
			encodedAddr := addr.EncodeAddress()
			encodedAddrs[j] = encodedAddr

			// No need to check the map again if the filter already
			// passes.
			if passesFilter {
				continue
			}
			if _, exists := filterAddrMap[encodedAddr]; exists {
				passesFilter = true
			}
		}

		if !passesFilter {
			continue
		}

		var vout btcjson.Vout
		vout.N = uint32(i)
		vout.Value = btcutil.Amount(v.Value).ToBTC()
		vout.ScriptPubKey.Addresses = encodedAddrs
		vout.ScriptPubKey.Asm = disbuf
		vout.ScriptPubKey.Hex = hex.EncodeToString(v.PkScript)
		vout.ScriptPubKey.Type = scriptClass.String()
		vout.ScriptPubKey.ReqSigs = int32(reqSigs)

		voutList = append(voutList, vout)
	}

	return voutList
}

func createVinList(mtx *wire.MsgTx) []btcjson.Vin {
	// Coinbase transactions only have a single txin by definition.
	vinList := make([]btcjson.Vin, len(mtx.TxIn))
	if blockchain.IsCoinBaseTx(mtx) {
		txIn := mtx.TxIn[0]
		vinList[0].Coinbase = hex.EncodeToString(txIn.SignatureScript)
		vinList[0].Sequence = txIn.Sequence
		vinList[0].Witness = witnessToHex(txIn.Witness)
		return vinList
	}

	for i, txIn := range mtx.TxIn {
		// The disassembled string will contain [error] inline
		// if the script doesn't fully parse, so ignore the
		// error here.
		disbuf, _ := txscript.DisasmString(txIn.SignatureScript)

		vinEntry := &vinList[i]
		vinEntry.Txid = txIn.PreviousOutPoint.Hash.String()
		vinEntry.Vout = txIn.PreviousOutPoint.Index
		vinEntry.Sequence = txIn.Sequence
		vinEntry.ScriptSig = &btcjson.ScriptSig{
			Asm: disbuf,
			Hex: hex.EncodeToString(txIn.SignatureScript),
		}

		if mtx.HasWitness() {
			vinEntry.Witness = witnessToHex(txIn.Witness)
		}
	}

	return vinList
}

// DecodeTx decode tx
func (a *Tools) DecodeTx(rawTx string) (result *btcjson.TxRawDecodeResult, witnessHash string, isSwTx bool, err error) {
	if a.Coin == consts.CoinTypeBitcoinCash {
		return bch.DecodeTx(rawTx)
	}
	var serializedTx []byte
	serializedTx, err = hex.DecodeString(rawTx)
	if err != nil {
		return
	}
	var mtx wire.MsgTx
	err = mtx.Deserialize(bytes.NewReader(serializedTx))
	if err != nil {
		return
	}

	var params chaincfg.Params
	if _, ok := MainNetParams[a.Coin]; ok {
		params = MainNetParams[a.Coin]
	} else {
		params = MainNetParams[consts.CoinTypeBitcoin]
	}

	// Create and return the result.
	txReply := btcjson.TxRawDecodeResult{
		Txid:     mtx.TxHash().String(),
		Version:  mtx.Version,
		Locktime: mtx.LockTime,
		Vin:      createVinList(&mtx),
		Vout:     createVoutList(&mtx, &params, nil),
	}
	witnessHash = mtx.WitnessHash().String()
	isSwTx = mtx.HasWitness()
	result = &txReply
	return
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

// GetSizeVsize 获取原始交易的体积和 Vsize
func GetSizeVsize(rawTx string) (size, vsize int32, err error) {
	var serializedTx []byte
	serializedTx, err = hex.DecodeString(rawTx)
	if err != nil {
		return
	}
	var mtx wire.MsgTx
	err = mtx.Deserialize(bytes.NewReader(serializedTx))
	if err != nil {
		return
	}
	vsize = int32(mempool.GetTxVirtualSize(btcutil.NewTx(&mtx)))
	return int32(mtx.SerializeSize()), vsize, nil
}
