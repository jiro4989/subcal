package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

type IP struct {
	IPAddress string
	CIDR      int
	Bin       string
	Mask      string
}

func ParseCIDR(s string) (IP, error) {
	ipv4Addr, ipv4Net, err := net.ParseCIDR(s)
	if err != nil {
		return IP{}, err
	}

	var ip IP
	ip.IPAddress = ipv4Addr.String()

	ones, _ := ipv4Net.Mask.Size()
	ip.CIDR = ones

	l := len(ipv4Addr)
	ip.Bin = fmt.Sprintf("%08b%08b%08b%08b", ipv4Addr[l-4], ipv4Addr[l-3], ipv4Addr[l-2], ipv4Addr[l-1])
	ip.Mask = fmt.Sprintf("%08b%08b%08b%08b", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])

	return ip, nil
}

func (ip IP) Format(sep string, colorFlag, ipv4Flag, cidrFlag, binFlag, maskFlag bool) string {
	var arr []string
	if ipv4Flag {
		arr = append(arr, ip.IPAddress)
	}
	if cidrFlag {
		arr = append(arr, fmt.Sprintf("%d", ip.CIDR))
	}
	if binFlag {
		arr = append(arr, ip.Bin)
	}
	if maskFlag {
		arr = append(arr, ip.Mask)
	}
	return strings.Join(arr, sep)
}

var RootCommand = &cobra.Command{
	Use:   "subcal",
	Short: "subcal",
	Long:  "subcal",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("stdin")
			return
		}
		for _, ipcidr := range args {
			ip, err := ParseCIDR(ipcidr)
			if err != nil {
				panic(err)
			}
			fmt.Println(ip)
		}
	},
}

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().StringP("sep", "s", "\t", "color")
	RootCommand.Flags().BoolP("color", "c", false, "color")
	RootCommand.Flags().BoolP("ipv4", "i", false, "color")
	RootCommand.Flags().BoolP("cidr", "d", false, "color")
	RootCommand.Flags().BoolP("bin", "b", false, "color")
	RootCommand.Flags().BoolP("mask", "m", false, "color")
	RootCommand.Flags().BoolP("noheader", "H", false, "color")
}
