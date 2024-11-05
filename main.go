package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jiro4989/subcal/ip"
)

type (
	ErrorCode int
)

const (
	doc = `subcal is a command to calculate subnet mask.

Usage:
	saubcal [options] <ip>...
	saubcal -h | --help
	saubcal -v | --version

Options:
	-h --help                     Print this help.
	-v --version                  Print version.
	-d --delimiter=<DELIMITER>    Set field delimiter. [default:  ]
	-C --color                    Colorize IP address bin.
	-i --ipv4                     Print IPv4 address.
	-c --cidr                     Print CIDR.
	-b --bin                      Print IP address bin.
	-m --mask                     Print Subnet mask.
	-n --no-header                Hide header.`
)

const (
	NoError ErrorCode = iota
	DocoptError
	ParseCIDRError
)

func main() {
	args := ParseArgs()
	os.Exit(int(Main(args)))
}

func Main(args *CmdArgs) ErrorCode {
	if !args.IPv4 && !args.CIDR && !args.Bin && !args.Mask {
		args.IPv4 = true
		args.CIDR = true
		args.Bin = true
		args.Mask = true
	}

	// ヘッダの出力
	if !args.NoHeader {
		fmt.Println(header(args.Delimiter, args.IPv4, args.CIDR, args.Bin, args.Mask))
	}

	// ボディの出力
	for _, ipcidr := range args.Args {
		ipaddr, err := ip.ParseCIDR(ipcidr)
		if err != nil {
			return ParseCIDRError
		}
		fmt.Println(ipaddr.Format(args.Delimiter, args.Color, args.IPv4, args.CIDR, args.Bin, args.Mask))
	}

	return NoError
}

func header(sep string, ipv4Flag, cidrFlag, binFlag, maskFlag bool) string {
	var arr []string
	if ipv4Flag {
		arr = append(arr, "IPv4")
	}
	if cidrFlag {
		arr = append(arr, "CIDR")
	}
	if binFlag {
		arr = append(arr, "Bin")
	}
	if maskFlag {
		arr = append(arr, "Mask")
	}
	return strings.Join(arr, sep)
}
