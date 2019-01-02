// Package ip provides calc IP CIDR
package ip

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
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
		if colorFlag {
			b1 := ip.Bin[:ip.CIDR]
			b2 := ip.Bin[ip.CIDR:]
			s := fmt.Sprintf("%s%s", color.RedString(b1), color.GreenString(b2))
			arr = append(arr, s)
		} else {
			arr = append(arr, ip.Bin)
		}
	}
	if maskFlag {
		arr = append(arr, ip.Mask)
	}
	return strings.Join(arr, sep)
}
