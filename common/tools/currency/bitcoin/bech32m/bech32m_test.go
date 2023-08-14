// Copyright (c) 2017 Takatoshi Nakagawa
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package bech32m

import (
	"encoding/hex"
	"strings"
	"testing"
)

func segwitScriptpubkey(witver byte, witprog []byte) []byte {
	// Construct a Segwit scriptPubKey for a given witness program.
	if witver != 0 {
		witver += 0x50
	}
	return append(append([]byte{witver}, byte(len(witprog))), witprog...)
}

var validBech32 = []string{
	"A12UEL5L",
	"a12uel5l",
	"an83characterlonghumanreadablepartthatcontainsthenumber1andtheexcludedcharactersbio1tt5tgs",
	"abcdef1qpzry9x8gf2tvdw0s3jn54khce6mua7lmqqqxw",
	"11qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqc8247j",
	"split1checkupstagehandshakeupstreamerranterredcaperred2y9e3w",
	"?1ezyfcl",
}

var validBech32m = []string{
	"A1LQFN3A",
	"a1lqfn3a",
	"an83characterlonghumanreadablepartthatcontainsthetheexcludedcharactersbioandnumber11sg7hg6",
	"abcdef1l7aum6echk45nj3s0wdvt2fg8x9yrzpqzd3ryx",
	"11llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllludsr8",
	"split1checkupstagehandshakeupstreamerranterredcaperredlc445v",
	"?1v759aa",
}

var invalidBech32 = []string{
	" 1nwldj5",         // HRP character out of range
	"\x7F" + "1axkwrx", // HRP character out of range
	"\x80" + "1eym55h", // HRP character out of range
	// overall max length exceeded
	"an84characterslonghumanreadablepartthatcontainsthenumber1andtheexcludedcharactersbio1569pvx",
	"pzry9x0s0muk",      // No separator character
	"1pzry9x0s0muk",     // Empty HRP
	"x1b4n0q5v",         // Invalid data character
	"li1dgmt3",          // Too short checksum
	"de1lg7wt" + "\xFF", // Invalid character in checksum
	"A1G7SGD8",          // checksum calculated with uppercase form of HRP
	"10a06t8",           // empty HRP
	"1qzzfhee",          // empty HRP
}

var invalidBech32m = []string{
	" 1xj0phk",         // HRP character out of range
	"\x7F" + "1g6xzxy", // HRP character out of range
	"\x80" + "1vctc34", // HRP character out of range
	// overall max length exceeded
	"an84characterslonghumanreadablepartthatcontainsthetheexcludedcharactersbioandnumber11d6pts4",
	"qyrz8wqd2c9m",  // No separator character
	"1qyrz8wqd2c9m", // Empty HRP
	"y1b0jsk6g",     // Invalid data character
	"lt1igcx5c0",    // Invalid data character
	"in1muywd",      // Too short checksum
	"mm1crxm3i",     // Invalid character in checksum
	"au1s5cgom",     // Invalid character in checksum
	"M1VUXWEZ",      // Checksum calculated with uppercase form of HRP
	"16plkw9",       // Empty HRP
	"1p2gdwpf",      // Empty HRP
}

var validAddress = [][]string{
	{"BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4", "0014751e76e8199196d454941c45d1b3a323f1433bd6"},
	{"bc1puxkz8vpy900c7z4q4302lkc3jjr2s42mayfqzzgr5yqdk5mgma3s0kntlh", "5120e1ac23b0242bdf8f0aa0ac5eafdb119486a8555be912010903a100db5368df63"},
	{"tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7",
		"00201863143c14c5166804bd19203356da136c985678cd4d27a1b8c6329604903262"},
	{"bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7kt5nd6y",
		"5128751e76e8199196d454941c45d1b3a323f1433bd6751e76e8199196d454941c45d1b3a323f1433bd6"},
	{"BC1SW50QGDZ25J", "6002751e"},
	{"bc1zw508d6qejxtdg4y5r3zarvaryvaxxpcs", "5210751e76e8199196d454941c45d1b3a323"},
	{"tb1qqqqqp399et2xygdj5xreqhjjvcmzhxw4aywxecjdzew6hylgvsesrxh6hy",
		"0020000000c4a5cad46221b2a187905e5266362b99d5e91c6ce24d165dab93e86433"},
	{"tb1pqqqqp399et2xygdj5xreqhjjvcmzhxw4aywxecjdzew6hylgvsesf3hn0c",
		"5120000000c4a5cad46221b2a187905e5266362b99d5e91c6ce24d165dab93e86433"},
	{"bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqzk5jj0",
		"512079be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"},
}

var invalidAddress = []string{
	// Invalid HRP
	"tc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq5zuyut",
	// Invalid checksum algorithm (bech32 instead of bech32m)
	"bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqh2y7hd",
	// Invalid checksum algorithm (bech32 instead of bech32m)
	"tb1z0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqglt7rf",
	// Invalid checksum algorithm (bech32 instead of bech32m)
	"BC1S0XLXVLHEMJA6C4DQV22UAPCTQUPFHLXM9H8Z3K2E72Q4K9HCZ7VQ54WELL",
	// Invalid checksum algorithm (bech32m instead of bech32)
	"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kemeawh",
	// Invalid checksum algorithm (bech32m instead of bech32)
	"tb1q0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq24jc47",
	// Invalid character in checksum
	"bc1p38j9r5y49hruaue7wxjce0updqjuyyx0kh56v8s25huc6995vvpql3jow4",
	// Invalid witness version
	"BC130XLXVLHEMJA6C4DQV22UAPCTQUPFHLXM9H8Z3K2E72Q4K9HCZ7VQ7ZWS8R",
	// Invalid program length (1 byte)
	"bc1pw5dgrnzv",
	// Invalid program length (41 bytes)
	"bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7v8n0nx0muaewav253zgeav",
	// Invalid program length for witness version 0 (per BIP141)
	"BC1QR508D6QEJXTDG4Y5R3ZARVARYV98GJ9P",
	// Mixed case
	"tb1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq47Zagq",
	// More than 4 padding bits
	"bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7v07qwwzcrf",
	// Non-zero padding in 8-to-5 conversion
	"tb1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vpggkg4j",
	// Empty data section
	"bc1gmk9yu",
}

var invalidAddressEnc = [][]interface{}{
	{"BC", 0, 20},
	{"bc", 0, 21},
	{"bc", 17, 32},
	{"bc", 1, 1},
	{"bc", 16, 41},
}

func TestValidChecksum(t *testing.T) {
	// Test checksum creation and validation.
	specs := []int{Bech32, Bech32m}
	for _, spec := range specs {
		tests := validBech32m
		if spec == Bech32 {
			tests = validBech32
		}
		for _, test := range tests {
			_, _, dspec, err := Decode(test)
			if err != nil {
				t.Errorf("NG : %s / %+v", test, err)
				continue
			}
			if spec != dspec {
				t.Errorf("NG : %s", test)
				continue
			}
			pos := strings.LastIndex(test, "1")
			test2 := test[:pos+1] + string(test[pos+1]^1) + test[pos+2:]
			_, _, dspec, err = Decode(test2)
			if err == nil {
				t.Errorf("NG : %s", test2)
				continue
			}
			if dspec > 0 {
				t.Errorf("NG : %s, spec: %v, dspec: %v", test, spec, dspec)
				continue
			}
			t.Logf("OK : %s", test)
		}
	}
}

func TestInvalidChecksum(t *testing.T) {
	// Test checksum creation and validation.
	specs := []int{Bech32, Bech32m}
	for _, spec := range specs {
		tests := invalidBech32m
		if spec == Bech32 {
			tests = invalidBech32
		}
		for _, test := range tests {
			_, _, dspec, err := Decode(test)
			if err == nil {
				if spec == dspec {
					t.Errorf("NG : %s", test)
					continue
				}
			}
			t.Logf("OK : %s", err)
		}
	}
}

func TestValidAddress(t *testing.T) {
	// Test whether valid addresses decode to the correct output.
	for _, test := range validAddress {
		address := test[0]
		hexscript := test[1]
		hrp := "bc"
		witver, witprog, err := SegwitAddrDecode(hrp, address)
		if err != nil {
			hrp = "tb"
			witver, witprog, err = SegwitAddrDecode(hrp, address)
		}
		if err != nil {
			t.Errorf("NG : %s / %+v", test, err)
			continue
		}
		scriptpubkey := segwitScriptpubkey(witver, witprog)
		if hexscript != hex.EncodeToString(scriptpubkey) {
			t.Errorf("NG : %s", test)
			continue
		}
		addr, err := SegwitAddrEncode(hrp, witver, witprog)
		if err != nil {
			t.Errorf("NG : %s / %+v", test, err)
			continue
		}
		if !strings.EqualFold(strings.ToLower(address), addr) {
			t.Errorf("NG : %s", test)
			continue
		}
		t.Logf("OK : %s", test)
	}
}

func TestInvalidAddress(t *testing.T) {
	// Test whether invalid addresses fail to decode.
	for _, test := range invalidAddress {
		ver, _, err := SegwitAddrDecode("bc", test)
		if err == nil {
			t.Errorf("NG %d : %s", ver, test)
			continue
		}
		t.Logf("OK : %v", err)
		_, _, err = SegwitAddrDecode("tb", test)
		if err == nil {
			t.Errorf("NG : %s", test)
			continue
		}
		t.Logf("OK : %v", err)
	}
}

func TestInvalidAddressEnc(t *testing.T) {
	// Test whether address encoding fails on invalid input.
	for _, test := range invalidAddressEnc {
		hrp := test[0].(string)
		version := test[1].(int)
		length := test[2].(int)
		prog := make([]byte, length)
		_, err := SegwitAddrEncode(hrp, byte(version), prog)
		if err == nil {
			t.Logf("NG : %+v", test)
			t.Errorf("%+v", err)
			continue
		}
		t.Logf("OK : %v", err)
	}
}
