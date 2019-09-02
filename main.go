package main

import (
	"fmt"
	"os"

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
		Header    bool
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
	-H --header       delimiter`
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

	for _, ipcidr := range config.IP {
		ipaddr, err := ip.ParseCIDR(ipcidr)
		if err != nil {
			return 1
		}
		fmt.Println(ipaddr.Format(config.Delimiter, config.Color, config.IPv4, config.CIDR, config.Bin, config.Mask))
	}

	return 0
}
