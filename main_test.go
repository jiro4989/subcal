package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	type TestData struct {
		desc string
		in   *CmdArgs
		want ErrorCode
	}
	tests := []TestData{
		{
			desc: "正常系: 正常なIP",
			in: &CmdArgs{
				Args: []string{"192.168.255.0/24"},
			},
			want: NoError,
		},
		{
			desc: "正常系: 全てのオプション",
			in: &CmdArgs{
				Delimiter: ",",
				Color:     true,
				IPv4:      true,
				CIDR:      true,
				Bin:       true,
				Mask:      true,
				NoHeader:  true,
				Args:      []string{"192.168.255.0/24"},
			},
			want: NoError,
		},
		{
			desc: "異常系: 不正なIP",
			in: &CmdArgs{
				Args: []string{"192a.168.255.0/24"},
			},
			want: ParseCIDRError,
		},
		// {desc: "引数不足", in: []string{"./bin/subcal"}, want: DocoptError},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := Main(tt.in)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}

func TestHeader(t *testing.T) {
	type TestData struct {
		desc   string
		inSep  string
		inIPv4 bool
		inCIDR bool
		inBin  bool
		inMask bool
		want   string
	}
	tests := []TestData{
		{
			desc:   "すべてのオプションを指定",
			inSep:  ",",
			inIPv4: true,
			inCIDR: true,
			inBin:  true,
			inMask: true,
			want:   "IPv4,CIDR,Bin,Mask",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := header(
				tt.inSep,
				tt.inIPv4,
				tt.inCIDR,
				tt.inBin,
				tt.inMask,
			)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
