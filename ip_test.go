package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIPBin(t *testing.T) {
	type TD struct {
		args   string
		expect CIDRAddress
		err    error
		desc   string
	}
	tds := []TD{
		TD{
			args: "0.0.0.0/24",
			expect: CIDRAddress{
				IPDec:   "0.0.0.0",
				IPBin:   00000000000000000000000000000000,
				CIDR:    24,
				Network: 000000000000000000000000,
				Host:    00000000,
			},
			err:  nil,
			desc: "0アドレス/24",
		},
		TD{
			args: "0.0.0.1/24",
			expect: CIDRAddress{
				IPDec:   "0.0.0.1",
				IPBin:   00000000000000000000000000000001,
				CIDR:    24,
				Network: 000000000000000000000000,
				Host:    00000001,
			},
			err:  nil,
			desc: "0.0.0.1/24",
		},
		TD{
			args: "0.0.1.0/24",
			expect: CIDRAddress{
				IPDec:   "0.0.1.0",
				IPBin:   00000000000000000000000100000000,
				CIDR:    24,
				Network: 000000000000000000000001,
				Host:    00000000,
			},
			err:  nil,
			desc: "0.0.1.0/24",
		},
	}
	for _, v := range tds {
		got, err := ToCIDRAddress(v.args)
		assert.Equal(t, v.expect, got, v.desc)
		assert.IsType(t, v.err, err, v.desc)
	}
}
