package cryptos

import (
	"crypto/sha256"
	"hash"
	"reflect"
	"testing"
)

func TestHMACsha256(t *testing.T) {
	type args struct {
		message string
		secret  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// https://www.devglan.com/online-tools/hmac-sha256-online base64
		{name: "normal: get HMACsha256", args: args{message: "abcde", secret: "123456"},
			want: "maUBb9typ59e5IfvIlNroej3qlAXh1jOO85OL1CAewk="},
		{name: "error: get HMACsha256 empty secret", args: args{message: "abcde", secret: ""},
			want: "tX4dts+xEIds9cjafzlDqwTBb6q9HsyViAWvCs0jGrc="},
		{name: "error: get HMACsha256 empty message", args: args{message: "", secret: "12345"},
			want: "1w2IzZrfn5KEcsyV9YsUFamF3wP04jhBME6h0dsFQz4="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HMACsha256(tt.args.message, tt.args.secret); got != tt.want {
				t.Errorf("HMACsha256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: get MD5", args: args{text: "abcde"}, want: "ab56b4d92b40713acc5af89985d4b786"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.text); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: get sha256", args: args{text: "abcde"},
			want: "36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.args.text); got != tt.want {
				t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringHasher(t *testing.T) {
	type args struct {
		algorithm hash.Hash
		text      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal: get sha256 stringHasher", args: args{algorithm: sha256.New(), text: "abcde"},
			want: "36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringHasher(tt.args.algorithm, tt.args.text); got != tt.want {
				t.Errorf("stringHasher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleHashH(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want [32]byte
	}{
		{name: "normal: get DoubleHashH", args: args{b: []byte("1234")},
			want: [32]byte{207, 195, 43, 97, 219, 11, 205, 215, 28, 186, 114, 11, 101, 249, 251, 110, 107, 116, 176,
				4, 76, 45, 31, 94, 122, 106, 31, 144, 73, 161, 207, 155}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleHashH(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleHashH() = %v, want %v", got, tt.want)
			}
		})
	}
}
