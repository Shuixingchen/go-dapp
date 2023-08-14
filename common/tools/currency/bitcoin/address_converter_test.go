package bitcoin

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestTools_GetAddressType(t *testing.T) {
	type fields struct {
		Coin string
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantAddrType string
		wantErr      bool
	}{
		{name: "bc1 test1", fields: fields{Coin: "BTC"}, args: args{addr: "bc1qu4yhw0r8zs94zlazkp367vt2ppa6pt00zpmjmln4l9xu8quzgjds84axh0"},
			wantAddrType: "P2WSH_V0", wantErr: false},
		{name: "p2pkh test1", fields: fields{Coin: "BTC"}, args: args{addr: "1GX28yLjVWux7ws4UQ9FB4MnLH4UKTPK2z"},
			wantAddrType: "P2PKH", wantErr: false},
		{name: "p2sh test1", fields: fields{Coin: "BTC"}, args: args{addr: "33aF1P2uA8XuQUHtxtNQq3YVrKU8JNEequ"},
			wantAddrType: "P2SH", wantErr: false},
		{name: "bc1 test2", fields: fields{Coin: "BTC"}, args: args{addr: "bc1qr438mvrsh9lwlymwt7zdug2mcpq3206z3ve4t6"},
			wantAddrType: "P2WPKH_V0", wantErr: false},
		// b7e3c981b06983cdb0082dd8e386ce6715805a75e1b8b3ce2df45c6e48172b6a,bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqwl4cdu
		// https://www.walletexplorer.com/txid/b7e3c981b06983cdb0082dd8e386ce6715805a75e1b8b3ce2df45c6e48172b6a
		// wallet explorer is bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqmr95g7
		{name: "bc1 witness version1", fields: fields{Coin: "BTC"}, args: args{addr: "bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqwl4cdu"},
			wantAddrType: "WITNESS_UNKNOWN", wantErr: false},
		{name: "bc1 witness version1 test2", fields: fields{Coin: "BTC"}, args: args{addr: "bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqmr95g7"},
			wantAddrType: "WITNESS_V1_TAPROOT", wantErr: false},
		{name: "bc1 witness no version test", fields: fields{Coin: "BTC"}, args: args{addr: "bc1pveaamy78cq5hvl74zmfw52fxyjun3lh7lgt44j03ygx02zyk8lesgk06f6"},
			wantAddrType: "WITNESS_V1_TAPROOT", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Tools{
				Coin: tt.fields.Coin,
			}
			gotAddrType, err := a.GetAddressType(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tools.GetAddressType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddrType != tt.wantAddrType {
				t.Errorf("Tools.GetAddressType() = %v, want %v", gotAddrType, tt.wantAddrType)
			}
		})
	}
}

func TestTools_AddressToHex(t *testing.T) {
	type fields struct {
		Coin string
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "bc1 p2wpkh V0", fields: fields{Coin: "BTC"}, args: args{addr: "bc1qr438mvrsh9lwlymwt7zdug2mcpq3206z3ve4t6"},
			want: "1d627db070b97eef936e5f84de215bc041153f42", wantErr: false},
		// 	{"BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4", "0014751e76e8199196d454941c45d1b3a323f1433bd6"},
		{name: "bech32 test3", fields: fields{Coin: "BTC"}, args: args{addr: "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4"},
			want: "751e76e8199196d454941c45d1b3a323f1433bd6", wantErr: false},
		{name: "empty test case", fields: fields{Coin: ""}, args: args{addr: ""}, want: "", wantErr: false},
		{name: "p2sh", fields: fields{Coin: "BTC"}, args: args{addr: "34Z3Vraosmdr6G9PETCB7hM1fvGHLx2JJH"},
			want: "1f665d257910a34b2622b74d4d4ca9ece48e9f85", wantErr: false},
		{name: "p2pkh 压缩格式", fields: fields{Coin: "BTC"}, args: args{addr: "162zktHBAcHWJrhSmB4NKA2r9GtabZQs3y"},
			want: "373941f767cc77a0f438070377805f3c6a1697e7", wantErr: false},
		{name: "p2sh 未压缩格式，结果也是未压缩格式", fields: fields{Coin: "BTC"}, args: args{addr: "33D9Fd5Rq2NvgnLU4yqVoVfiHa34eT4kTu"},
			want: "10aab8f824c7561a2c8d24b64dabf964c0b3722a", wantErr: false},
		// witness unknown, 23b73c5c07a8b591e4460fa7b22bf3d1707e7e2f2254c315ee8aeecf46d81fd3
		{name: "bech32 segwit unknown", fields: fields{Coin: "BTC"}, args: args{addr: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k633322m7s8v"},
			want: "", wantErr: true},
		// https://github.com/btcsuite/btcd/btcutil/pull/202, normal bech32m addr
		{name: "bech32m test2", fields: fields{Coin: "BTC"}, args: args{addr: "bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh"},
			want: "e1ac23b0242bdf8f0aa0ac5eafdb119486a8555be912010903a100db5368df63", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Tools{
				Coin: tt.fields.Coin,
			}
			got, err := a.AddressToHex(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tools.AddressToHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Tools.AddressToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_PKScript2Addr(t *testing.T) {
	type fields struct {
		Coin string
	}
	type args struct {
		pkScript string
	}
	script1, _ := hex.DecodeString("76a914128004ff2fcaf13b2b91eb654b1dc2b674f7ec6188ac")
	script2, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	script3, _ := hex.DecodeString("5210751e76e8199196d454941c45d1b3a323")
	script4, _ := hex.DecodeString("512079be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantEncodedAddrs []string
		wantAddressType  string
	}{
		{name: "", fields: fields{Coin: "BTC"}, args: args{pkScript: string(script1)},
			wantEncodedAddrs: []string{"12gpXQVcCL2qhTNQgyLVdCFG2Qs2px98nV"}, wantAddressType: "pubkeyhash"},
		// https://github.com/bitcoin/bips/blob/master/bip-0350.mediawiki#test-vectors-for-v0-v16-native-segregated-witness-addresses
		{name: "bech32m test1", fields: fields{Coin: "BTC"}, args: args{pkScript: string(script2)},
			wantEncodedAddrs: []string{"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4"}, wantAddressType: "witness_v0_keyhash"},
		{name: "bech32m test2", fields: fields{Coin: "BTC"}, args: args{pkScript: string(script3)},
			wantEncodedAddrs: []string{"bc1zw508d6qejxtdg4y5r3zarvaryvaxxpcs"}, wantAddressType: "witness_v1_taproot"},
		{name: "bech32m test2", fields: fields{Coin: "BTC"}, args: args{pkScript: string(script4)},
			wantEncodedAddrs: []string{"bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqzk5jj0"}, wantAddressType: "witness_v1_taproot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Tools{
				Coin: tt.fields.Coin,
			}
			gotEncodedAddrs, gotAddressType := a.PKScript2Addr(tt.args.pkScript)
			if !reflect.DeepEqual(gotEncodedAddrs, tt.wantEncodedAddrs) {
				t.Errorf("Tools.PKScript2Addr() gotEncodedAddrs = %v, want %v", gotEncodedAddrs, tt.wantEncodedAddrs)
			}
			if gotAddressType != tt.wantAddressType {
				t.Errorf("Tools.PKScript2Addr() gotAddressType = %v, want %v", gotAddressType, tt.wantAddressType)
			}
		})
	}
}

func TestTools_HexToAddress(t *testing.T) {
	type fields struct {
		Coin string
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// {name: "test1", fields: fields{Coin: "BTC"}, args: args{addr: "0000"}, want: "bc1q9zpgru"},
		// {name: "bech32m test1", fields: fields{Coin: "BTC"}, args: args{addr: "050102ac06077ed65d953ff8ee01eabc2fa57ab2c19e61a0f94d5c78cafea5f6d46315"},
		//	want: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k633322m7s8v"},
		// 23b73c5c07a8b591e4460fa7b22bf3d1707e7e2f2254c315ee8aeecf46d81fd3
		{name: "bech32m test1", fields: fields{Coin: "BTC"}, args: args{addr: "050102ac06077ed65d953ff8ee01eabc2fa57ab2c19e61a0f94d5c78cafea5f6d46315"},
			want: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k63332l8wuzw"},
		// bc1pveaamy78cq5hvl74zmfw52fxyjun3lh7lgt44j03ygx02zyk8lesgk06f6
		{name: "bech32m test2", fields: fields{Coin: "BTC"}, args: args{addr: "06667bdd93c7c029767fd516d2ea292624b938fefefa175ac9f1220cf508963ff3"},
			want: "bc1pveaamy78cq5hvl74zmfw52fxyjun3lh7lgt44j03ygx02zyk8lesgk06f6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Tools{
				Coin: tt.fields.Coin,
			}
			if got := a.HexToAddress(tt.args.addr); got != tt.want {
				t.Errorf("Tools.HexToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
