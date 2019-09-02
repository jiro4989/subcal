package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/subcal/ip"
)

type (
	Config struct {
		Delimiter string `docopt:"-d,--delimiter"`
		Color     bool
		IPv4      bool `docopt:"--ipv4"`
		CIDR      bool `docopt:"--cidr"`
		Bin       bool
		Mask      bool
		NoHeader  bool
		IP        []string `docopt:"<ip>"`
	}
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
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return DocoptError
	}

	if !config.IPv4 && !config.CIDR && !config.Bin && !config.Mask {
		config.IPv4 = true
		config.CIDR = true
		config.Bin = true
		config.Mask = true
	}

	// ヘッダの出力
	if !config.NoHeader {
		fmt.Println(header(config.Delimiter, config.IPv4, config.CIDR, config.Bin, config.Mask))
	}

	// ボディの出力
	for _, ipcidr := range config.IP {
		ipaddr, err := ip.ParseCIDR(ipcidr)
		if err != nil {
			return ParseCIDRError
		}
		fmt.Println(ipaddr.Format(config.Delimiter, config.Color, config.IPv4, config.CIDR, config.Bin, config.Mask))
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
