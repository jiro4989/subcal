package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	Version   bool
	Delimiter string
	Color     bool
	IPv4      bool
	CIDR      bool
	Bin       bool
	Mask      bool
	NoHeader  bool
	Args      []string
}

const (
	appName = "subcal"

	helpMsgHelp      = "print help"
	helpMsgVersion   = "print version"
	helpMsgDelimiter = "set field delimiter"
	helpMsgColor     = "colorize ip address bin"
	helpMsgIPv4      = "print ipv4 address"
	helpMsgCIDR      = "print cidr"
	helpMsgBin       = "print ipv4 address bin"
	helpMsgMask      = "print subnet mask"
	helpMsgNoHeader  = "hide header"
)

func ParseArgs() *CmdArgs {
	args := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.BoolVar(&args.Version, "v", false, helpMsgVersion)
	flag.StringVar(&args.Delimiter, "d", " ", helpMsgDelimiter)
	flag.BoolVar(&args.Color, "c", false, helpMsgColor)
	flag.BoolVar(&args.NoHeader, "n", false, helpMsgNoHeader)
	flag.BoolVar(&args.IPv4, "ipv4", false, helpMsgIPv4)
	flag.BoolVar(&args.CIDR, "cidr", false, helpMsgCIDR)
	flag.BoolVar(&args.Bin, "bin", false, helpMsgBin)
	flag.BoolVar(&args.Mask, "mask", false, helpMsgMask)
	flag.Parse()
	args.Args = flag.Args()

	return &args
}

func flagHelpMessage() {
	out := os.Stderr
	fmt.Fprintln(out, fmt.Sprintf("%s is a command to calculate subnet mask", appName))
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Usage:")
	fmt.Fprintln(out, fmt.Sprintf("  %s [OPTIONS] <ip>...", appName))
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Examples:")
	fmt.Fprintln(out, fmt.Sprintf("  %s 10.0.1.0/24", appName))
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Options:")

	flag.PrintDefaults()
}
