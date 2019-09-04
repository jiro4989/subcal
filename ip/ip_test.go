package ip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCIDR(t *testing.T) {
	type TestData struct {
		desc    string
		in      string
		want    IP
		wantErr bool
	}
	tests := []TestData{
		{
			desc: "通常のIPアドレス",
			in:   "254.1.255.0/24",
			want: IP{
				IPAddress: "254.1.255.0",
				CIDR:      24,
				Bin:       "11111110000000011111111100000000",
				Mask:      "11111111111111111111111100000000",
			},
			wantErr: false,
		},
		{
			desc:    "不正なIPアドレス",
			in:      "a254.1.255.0/24",
			want:    IP{},
			wantErr: true,
		},
		{
			desc:    "不正なIPアドレス (数値の不正)",
			in:      "254..255.0/24",
			want:    IP{},
			wantErr: true,
		},
		{
			desc:    "不正なIPアドレス (CIDRが空)",
			in:      "192.168.255.0/",
			want:    IP{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ParseCIDR(tt.in)
			if tt.wantErr {
				assert.NotNil(t, err, tt.desc)
			} else {
				assert.Equal(t, tt.want, got, tt.desc)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	type TestData struct {
		desc    string
		inIP    IP
		inSep   string
		inColor bool
		inIPv4  bool
		inCIDR  bool
		inBin   bool
		inMask  bool
		want    string
	}
	tests := []TestData{
		{
			desc: "通常のIPアドレス",
			inIP: IP{
				IPAddress: "254.1.255.0",
				CIDR:      24,
				Bin:       "11111110000000011111111100000000",
				Mask:      "11111111111111111111111100000000",
			},
			inSep:   " ",
			inColor: true,
			inIPv4:  true,
			inCIDR:  true,
			inBin:   true,
			inMask:  true,
			want:    "254.1.255.0 24 11111110000000011111111100000000 11111111111111111111111100000000",
		},
		{
			desc: "通常のIPアドレス",
			inIP: IP{
				IPAddress: "254.1.255.0",
				CIDR:      24,
				Bin:       "11111110000000011111111100000000",
				Mask:      "11111111111111111111111100000000",
			},
			inSep:   ",",
			inColor: false,
			inIPv4:  true,
			inCIDR:  true,
			inBin:   true,
			inMask:  true,
			want:    "254.1.255.0,24,11111110000000011111111100000000,11111111111111111111111100000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.inIP.Format(
				tt.inSep,
				tt.inColor,
				tt.inIPv4,
				tt.inCIDR,
				tt.inBin,
				tt.inMask,
			)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
