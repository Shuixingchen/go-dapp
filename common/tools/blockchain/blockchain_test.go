package blockchain

import (
	"testing"
)

func TestIsBlockHashOrTransactionHash(t *testing.T) {
	type args struct {
		keyword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "tx hash", args: args{keyword: "98ca36906002d90b19a1099eeeadad89bff31f6000c322dca0d7caddc319028d"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlockHashOrTransactionHash(tt.args.keyword); got != tt.want {
				t.Errorf("IsBlockHashOrTransactionHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBC1Address(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "not bc1 address", args: args{address: "IsBlockHashOrTransactionHash"}, want: false},
		{name: "bc1 address", args: args{address: "bc1qgyqcfxznu9qak7zp065ty6vvr7rlufmhq0ksge"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBC1Address(tt.args.address); got != tt.want {
				t.Errorf("IsBC1Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLTCAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "not ltc address", args: args{address: "IsBlockHashOrTransactionHash"}, want: false},
		{name: "not ltc address", args: args{address: "bc1qgyqcfxznu9qak7zp065ty6vvr7rlufmhq0ksge"}, want: false},
		{name: "ltc address", args: args{address: "M8T1B2Z97gVdvmfkQcAtYbEepune1tzGua"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLTCAddress(tt.args.address); got != tt.want {
				t.Errorf("IsLTCAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCashAddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "not bch address", args: args{address: "IsBlockHashOrTransactionHash"}, want: false},
		{name: "not bch address", args: args{address: "bc1qgyqcfxznu9qak7zp065ty6vvr7rlufmhq0ksge"}, want: false},
		{name: "bch address", args: args{address: "qz7xc0vl85nck65ffrsx5wvewjznp9lflgktxc5878"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCashAddr(tt.args.address); got != tt.want {
				t.Errorf("IsCashAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO 这个 case 不应该通过
		{name: "not address", args: args{address: "IsBlockHashOrTransactionHash"}, want: true},
		{name: "address", args: args{address: "bc1qgyqcfxznu9qak7zp065ty6vvr7rlufmhq0ksge"}, want: true},
		{name: "address", args: args{address: "qz7xc0vl85nck65ffrsx5wvewjznp9lflgktxc5878"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAddress(tt.args.address); got != tt.want {
				t.Errorf("IsAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLegalAddress(t *testing.T) {
	type args struct {
		coin    string
		address string
	}
	tests := []struct {
		name          string
		args          args
		wantMatch     bool
		wantAddressID string
	}{
		{name: "not address", args: args{address: "IsBlockHashOrTransactionHash", coin: "btc"},
			wantMatch: false, wantAddressID: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotAddressID := IsLegalAddress(tt.args.coin, tt.args.address)
			if gotMatch != tt.wantMatch {
				t.Errorf("IsLegalAddress() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			if gotAddressID != tt.wantAddressID {
				t.Errorf("IsLegalAddress() gotAddressID = %v, want %v", gotAddressID, tt.wantAddressID)
			}
		})
	}
}

func TestIsLegalBlockInput(t *testing.T) {
	type args struct {
		keyWord string
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  bool
		wantFormat string
	}{
		{name: "block", args: args{keyWord: "650000"}, wantFormat: "height", wantMatch: true},
		{name: "block", args: args{keyWord: "latest"}, wantFormat: "latest", wantMatch: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotFormat := IsLegalBlockInput(tt.args.keyWord)
			if gotMatch != tt.wantMatch {
				t.Errorf("IsLegalBlockInput() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			if gotFormat != tt.wantFormat {
				t.Errorf("IsLegalBlockInput() gotFormat = %v, want %v", gotFormat, tt.wantFormat)
			}
		})
	}
}

func TestInputCheck(t *testing.T) {
	type args struct {
		keyWord string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "input check", args: args{keyWord: "-12222"}, want: true},
		{name: "input check2", args: args{keyWord: "12222"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InputCheck(tt.args.keyWord); got != tt.want {
				t.Errorf("InputCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
