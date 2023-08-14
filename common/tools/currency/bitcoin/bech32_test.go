package bitcoin

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

type TestBech32DecodeParam struct {
	Name        string
	Args        string
	Want        string
	WantVersion byte
	WantErr     bool
}

type TestBech32DecodeParams []TestBech32DecodeParam

func TestBech32DecodeByFile(t *testing.T) {
	cases := []struct {
		fixture   string
		returnErr bool
		name      string
	}{
		{
			fixture:   "testdata/invalid.csv",
			returnErr: true,
			name:      "InvalidFile",
		},
		{
			fixture:   "testdata/witness_unknown_addresses.csv",
			returnErr: false,
			name:      "witness unknown addresses",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := NewTestBech32DecodeParams(tc.fixture)
			returnedErr := err != nil
			if returnedErr != tc.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v", tc.returnErr, returnedErr)
			}
			for _, i := range g {
				_, got, addr, err := Bech32Decode(i.Args)
				if (err != nil) != i.WantErr {
					t.Errorf("Tools.Bech32Decode() Name: %v error = %v, wantErr %v", i.Name, err, i.WantErr)
					return
				}
				if err != nil {
					t.Log(err)
				}
				t.Log(len(addr))
				if got != i.Want {
					t.Errorf("Bech32Decode() = %v, name = %v, want %v", got, i.Name, i.Want)
				}
			}
		})
	}
}

func NewTestBech32DecodeParams(filepath string) (TestBech32DecodeParams, error) {
	var params TestBech32DecodeParams
	csvFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return params, err
		}

		if len(line) < 3 {
			return params, fmt.Errorf("Invalid file structure")
		}
		wantErr, _ := strconv.ParseBool(line[2])
		params = append(params, TestBech32DecodeParam{
			Name:    line[0],
			Args:    line[0],
			Want:    line[1],
			WantErr: wantErr,
		})
	}

	return params, nil
}

func TestBech32Decode(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantVersion byte
		wantErr     bool
	}{
		// taproot test case
		{name: "normal: btc bech32 version 1", args: args{a: "bc1pqyp2cps80mt9m9fllrhqr64u97jh4vkpnes6p72dt3uv4l497m2xx9g5nejjx"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 version 1 test1", args: args{a: "bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqsyjer9e"},
			want: "WITNESS_V1_TAPROOT", wantVersion: byte(1), wantErr: false},
		// b7e3c981b06983cdb0082dd8e386ce6715805a75e1b8b3ce2df45c6e48172b6a,bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqwl4cdu
		// https://www.walletexplorer.com/txid/b7e3c981b06983cdb0082dd8e386ce6715805a75e1b8b3ce2df45c6e48172b6a
		// wallet explorer is bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqmr95g7
		{name: "normal: btc bech32 version 1 test2", args: args{a: "bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqwl4cdu"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 version 1 test3", args: args{a: "bc1pw2knldczhudzzydsns4lree0fafdfn4j4nw0e5xx82lhpfvuxmtqmr95g7"},
			want: "WITNESS_V1_TAPROOT", wantVersion: byte(1), wantErr: false},
		// 23b73c5c07a8b591e4460fa7b22bf3d1707e7e2f2254c315ee8aeecf46d81fd3
		// WITNESS_UNKNOWN test case.
		{name: "normal: btc bech32 segwit v1", args: args{a: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k633322m7s8v"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 segwit v1", args: args{a: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k63332l8wuzw"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(1), wantErr: false},
		// https://github.com/btcsuite/btcd/btcutil/pull/202, bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh
		{name: "normal: btc bech32 version 1 test2", args: args{a: "bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh"},
			want: "WITNESS_V1_TAPROOT", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 version 1 test3", args: args{a: "bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqzk5jj0"},
			want: "WITNESS_V1_TAPROOT", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 version 1 test4", args: args{a: "bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7kt5nd6y"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(1), wantErr: false},
		{name: "normal: btc bech32 version 2", args: args{a: "bc1zw508d6qejxtdg4y5r3zarvaryvaxxpcs"},
			want: "WITNESS_UNKNOWN", wantVersion: byte(2), wantErr: false},
		{name: "normal: btc bech32 version 0", args: args{a: "bc1qr5x0qckcjj7nfckmpahl84nl7xdkz7sy6kefw2"},
			want: "P2WPKH_V0", wantVersion: byte(0), wantErr: false},
		{name: "normal: ltc bech32 version 0", args: args{a: "ltc1qfzlzyt267l4elwslxdxg589d5407dnlrjnlvpv"},
			want: "P2WPKH_V0", wantVersion: byte(0), wantErr: false},
		{name: "normal: btcv bech32 version 0", args: args{a: "royale1qzraycj2erawmvr4uddsva4dwurrsctq9kzwz8fv0a3cphzmajm4qedvjey"},
			want: "P2WSH_V0", wantVersion: byte(0), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version, got, _, err := Bech32Decode(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tools.Bech32Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Log(err)
			}
			if got != tt.want {
				t.Errorf("Bech32Decode() = %v, want %v", got, tt.want)
			}
			if version != tt.wantVersion {
				t.Errorf("Bech32Decode() version = %v, want %v", version, tt.wantVersion)
			}
		})
	}
}

func TestBech32Encode(t *testing.T) {
	type args struct {
		a   string
		hrp string
		b   int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {name: "normal: bech32 version 1", args: args{a: "02ac06077ed65d953ff8ee01eabc2fa57ab2c19e61a0f94d5c78cafea5f6d46315",
		//	hrp: "bc", b: 1}, want: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k633322m7s8v"},
		{name: "normal: bech32 version 1", args: args{a: "0101010101010101010101010101010101010101010101010101010101010101",
			hrp: "bc", b: 1}, want: "bc1pqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqsyjer9e"},
		{name: "normal: bech32 version 1", args: args{a: "da4710964f7852695de2da025290e24af6d8c281de5a0b902b7135fd9fd74d21",
			hrp: "bc", b: 1}, want: "bc1pmfr3p9j00pfxjh0zmgp99y8zftmd3s5pmedqhyptwy6lm87hf5sspknck9"},
		{name: "normal: bech32 version 1", args: args{a: "02ac06077ed65d953ff8ee01eabc2fa57ab2c19e61a0f94d5c78cafea5f6d46315",
			hrp: "bc", b: 1}, want: "bc1pq2kqvpm76ewe20lcacq740p054at9sv7vxs0jn2u0r90af0k63332l8wuzw"},
		{name: "normal: bech32 version 0", args: args{a: "bfc4819ea645fa0b75befbf6c86a9d4329f92a7d",
			hrp: "bc", b: 0}, want: "bc1qhlzgr84xghaqkad7l0mvs65agv5lj2nam8khll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Bech32Encode(tt.args.a, tt.args.hrp, tt.args.b); got != tt.want {
				t.Errorf("Bech32Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
