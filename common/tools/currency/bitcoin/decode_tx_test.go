package bitcoin

import "testing"

func TestGetSizeVsize(t *testing.T) {
	type args struct {
		rawTx string
	}
	tests := []struct {
		name      string
		args      args
		wantSize  int32
		wantVsize int32
		wantErr   bool
	}{
		{name: "normal", args: args{rawTx: "0200000001eb7c67d578e6e38b1856f37173877cae3b3a" +
			"1088f1e39784bd98e8a405dd8309000000006946304302202783f18b24aa1044aaefb07" +
			"4787b72033d8a55331e2f686df00e49bb35e2508d021f404e2940c3d3b54004f14b2105e7" +
			"acc98d67471fa679b688c3882a763cf2b7012102d073a99b86ee34196bac612d3d7e2b813" +
			"35a04c8fd7fd23877ba7a2c2b2423bcffffffff01e5e00400000000001976a9149a2b9ccf" +
			"ccd6067ad235d7201a93d52dc0150c5888ac00000000"}, wantSize: 190, wantVsize: 190, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, gotVsize, err := GetSizeVsize(tt.args.rawTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSizeVsize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSize != tt.wantSize {
				t.Errorf("GetSizeVsize() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
			if gotVsize != tt.wantVsize {
				t.Errorf("GetSizeVsize() gotVsize = %v, want %v", gotVsize, tt.wantVsize)
			}
		})
	}
}
