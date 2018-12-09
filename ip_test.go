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
		TD{args: "0.0.0.0/24"},
		TD{args: "0.0.0.1/24"},
		TD{args: "0.0.1.0/24"},
		TD{args: "0.1.0.0/24"},
		TD{args: "1.0.0.0/24"},
		TD{args: "172.19.20.11/24"},
	}
	for _, v := range tds {
		got, err := ToCIDRAddress(v.args)
		assert.Equal(t, v.expect, got, v.desc)
		assert.IsType(t, v.err, err, v.desc)
	}
}
