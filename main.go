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
)

const (
	doc = `subcal はサブネットマスクを簡単に参照するためのコマンドです。

Usage:
	saubcal [options] <ip>...
	saubcal -h | --help
	saubcal -v | --version

Options:
	-h --help         このヘルプを出力する。
	-v --version      バージョン情報を出力する。
	-d --delimiter=<DELIMITER>   delimiter [default:  ]
	-C --color        delimiter
	-i --ipv4         delimiter
	-c --cidr         delimiter
	-b --bin          delimiter
	-m --mask         delimiter
	-n --no-header       delimiter`
)

func main() {
	os.Exit(Main(os.Args))
}

func Main(argv []string) int {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
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
			return 1
		}
		fmt.Println(ipaddr.Format(config.Delimiter, config.Color, config.IPv4, config.CIDR, config.Bin, config.Mask))
	}

	return 0
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
