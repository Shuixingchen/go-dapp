package currency

import (
	"testing"
)

func TestAddressTypeConvert(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: get P2SH", args: args{a: "pubkeyhash"}, want: "P2PKH"},
		{name: "normal: get NULL_DATA, if not exists", args: args{a: "abc"}, want: "NULL_DATA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddressTypeConvert(tt.args.a); got != tt.want {
				t.Errorf("AddressTypeConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseTxHashString(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: reverse string", args: args{a: "123456789101"}, want: "019178563412"},
		{name: "normal: reverse string 2", args: args{a: "abcefg123321abcd"}, want: "cdab213312fgceab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseTxHashString(tt.args.a); got != tt.want {
				t.Errorf("ReverseTxHashString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuessType(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: guess type", args: args{a: "3FZsNnE2PJfhaAeRRtsNijm9WpCv4xvkkz"}, want: "P2SH"},
		{name: "normal: guess type 2", args: args{a: "bc1qaxsxnxdp2s3v8pumkk66zsc0hpfndzx8ygspp0"}, want: "P2WPKH_V0"},
		{name: "normal: guess type 3", args: args{a: "qpk4hk3wuxe2uqtqc97n8atzrrr6r5mleczf9sur4h"}, want: "P2PKH"},
		{name: "normal: guess type 4", args: args{a: "MLA73pJU2XDN87pbTssAfnH2BKuuyC6VNz"}, want: "P2SH"},
		{name: "normal: guess type 5", args: args{a: "bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9"}, want: "WITNESS_V1_TAPROOT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GuessType(tt.args.a); got != tt.want {
				t.Errorf("GuessType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToAddress(t *testing.T) {
	type args struct {
		a    string
		coin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: convert addr hex to address", args: args{a: "01E6176B3A18725E2404A6AB748ABD11700B0D2141",
			coin: "btc"}, want: "1MycUHutP9zPJsCqhVAWDmj4rbhKvtYE11"},
		{name: "normal: convert addr hex to address 2", args: args{a: "010748ACD4B809941AC529ED42379D290094C06EC1",
			coin: "btc"}, want: "1fWnu8Lc5TrtFSUpRPqixxFYq614rK3MW"},
		{name: "normal: convert addr hex to address 2", args: args{a: "038933816f3317abf0b6cc4d131694391858acbb7e",
			coin: "btc"}, want: "bc1q3yeczmenz74lpdkvf5f3d9perpv2ewm7kx0fhd"},
		// 0501da4710964f7852695de2da025290e24af6d8c281de5a0b902b7135fd9fd74d21
		{name: "normal: convert addr hex to address 3", args: args{a: "0501da4710964f7852695de2da025290e24af6d8c281de5a0b902b7135fd9fd74d21",
			coin: "btc"}, want: "bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToAddress(tt.args.coin, tt.args.a); got != tt.want {
				t.Errorf("GuessType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuessAddressCoinType(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "btc", args: args{address: "16k7a5x1G7t2BSMgTAnRjprajNdG68FRg5"}, want: "btc"},
		{name: "btc test2", args: args{address: "bc1qhlzgr84xghaqkad7l0mvs65agv5lj2nam8khll"}, want: "btc"},
		{name: "bch test", args: args{address: "qpk4hk3wuxe2uqtqc97n8atzrrr6r5mleczf9sur4h"}, want: "bch"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GuessAddressCoinType(tt.args.address); got != tt.want {
				t.Errorf("GuessAddressCoinType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetAddressID, https://github.com/btccom/bitcoin-explorer/blob/v4-v0.21.1/src/btcexplorer/Misc.h#L85
func TestGetAddressID(t *testing.T) {
	type args struct {
		coin        string
		address     string
		addressType string
	}
	tests := []struct {
		name          string
		args          args
		wantAddressID string
		wantErr       bool
	}{
		// 23b73c5c07a8b591e4460fa7b22bf3d1707e7e2f2254c315ee8aeecf46d81fd3
		{name: "normal witness unknown", args: args{coin: "BTC", address: "bc1pqyp2cps80mt9m9fllrhqr64u97jh4vkpnes6p72dt3uv4l497m2xx9g5nejjx",
			addressType: "WITNESS_UNKNOWN"},
			wantAddressID: "05010102ac06077ed65d953ff8ee01eabc2fa57ab2c19e61a0f94d5c78cafea5f6d46315", wantErr: false},
		// https://github.com/btcsuite/btcd/btcutil/pull/202, bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh
		{name: "normal witness v1 taproot", args: args{coin: "BTC", address: "bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh",
			addressType: "WITNESS_V1_TAPROOT"},
			wantAddressID: "0501e1ac23b0242bdf8f0aa0ac5eafdb119486a8555be912010903a100db5368df63", wantErr: false},
		// b53e3bc5edbb41b34a963ecf67eb045266cf841cab73a780940ce6845377f141
		{name: "normal", args: args{coin: "BTC", address: "bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqs3wf0qm",
			addressType: "WITNESS_V1_TAPROOT"},
			wantAddressID: "05010101010101010101010101010101010101010101010101010101010101010101", wantErr: false},
		// hash,block_height,block_idx
		// 7641c08f4bd299abfef26dcc6b477938f4a6c2eed2f224d1f5c1c86b4e09739d,654930,380
		{name: "normal taproot", args: args{coin: "BTC", address: "bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5ss52r5n8",
			addressType: "WITNESS_V1_TAPROOT"},
			wantAddressID: "0501da4710964f7852695de2da025290e24af6d8c281de5a0b902b7135fd9fd74d21",
			wantErr:       false},
		// normal p2wpkh, p2sh, p2pkh
		{name: "normal p2pkh", args: args{coin: "BTC", address: "1KFHE7w8BhaENAswwryaoccDb6qcT6DbYY", addressType: "P2PKH"},
			wantAddressID: "01c825a1ecf2a6830c4401620c3a16f1995057c2ab", wantErr: false},
		{name: "normal p2pkh test2, poolin", args: args{coin: "BTC", address: "1E6vTBe9KLCh5ZqEkLoV2Fh5syXVHxHkna", addressType: "P2PKH"},
			wantAddressID: "018fb85581d373ae12b6b865f719241a902bed13ec", wantErr: false},
		{name: "normal p2wpkh", args: args{coin: "BTC", address: "bc1q3yeczmenz74lpdkvf5f3d9perpv2ewm7kx0fhd",
			addressType: "P2WPKH_V0"},
			wantAddressID: "038933816f3317abf0b6cc4d131694391858acbb7e", wantErr: false},
		{name: "normal p2sh", args: args{coin: "BTC", address: "34fVZzBdM4szF9LACjzNGp5fx178robujs", addressType: "P2SH"},
			wantAddressID: "02209e9f6eebe5723fbfce9a95a12a3c33e3f220e3", wantErr: false},
		// bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9
		{name: "normal taproot", args: args{coin: "BTC", address: "bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9",
			addressType: "WITNESS_V1_TAPROOT"},
			wantAddressID: "0501da4710964f7852695de2da025290e24af6d8c281de5a0b902b7135fd9fd74d21",
			wantErr:       false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddressID, err := GetAddressID(tt.args.coin, tt.args.address, tt.args.addressType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddressID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddressID != tt.wantAddressID {
				t.Errorf("GetAddressID() = %v, want %v", gotAddressID, tt.wantAddressID)
			}
		})
	}
}

// TestV3HexToAddress Deprecated
func TestV3HexToAddress(t *testing.T) {
	type args struct {
		a    string
		coin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: convert addr hex to address", args: args{a: "03bfc4819ea645fa0b75befbf6c86a9d4329f92a7d",
			coin: "btc"}, want: "bc1qhlzgr84xghaqkad7l0mvs65agv5lj2nam8khll"},
		{name: "normal: convert addr hex to address 2", args: args{a: "013f00000000000000010000000005f5e100000000000000000000000000000000",
			coin: "btc"}, want: "16k7a5x1G7t2BSMgTAnRjprajNdG68FRg5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := V3HexToAddress(tt.args.coin, tt.args.a); got != tt.want {
				t.Errorf("GuessType() = %v, want %v", got, tt.want)
			}
		})
	}
}
