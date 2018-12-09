package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// CIDRAddress は 0.0.0.0/24 みたいなIP/CIDR文字列
type CIDRAddress struct {
	IPDec   string // 10進数IP
	IPBin   uint64 // 2進数IP
	CIDR    int    // /24 みたいなネットワーク部とホスト部の境界値
	Network uint64 // ネットワーク部
	Host    uint64 // ホスト部
}

// ToIPBin はIP/CIDRを2進数に変換する
func ToCIDRAddress(s string) (CIDRAddress, error) {
	ipv4Addr, ipv4Net, err := net.ParseCIDR(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	l := len(ipv4Addr)
	fmt.Printf("ipv4addr: %08b%08b%08b%08b\n", ipv4Addr[l-4], ipv4Addr[l-3], ipv4Addr[l-2], ipv4Addr[l-1])
	fmt.Printf("IP:       %08b%08b%08b%08b\n", ipv4Net.IP[0], ipv4Net.IP[1], ipv4Net.IP[2], ipv4Net.IP[3])
	fmt.Printf("Mask:     %08b%08b%08b%08b\n", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])

	return CIDRAddress{IPDec: "xxx"}, nil
}

func IPDecToBinString(ip string) (string, error) {
	ipDecs := strings.Split(ip, ".")
	if len(ipDecs) < 4 {
		return "", errors.New("IPが不正 ip=" + ip)
	}
	var ipBins []uint64
	for _, v := range ipDecs {
		bin, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("IPが不正 ip=%s, bin=%b, err=%v", ip, bin, err)
			return "", errors.New(msg)
		}
		ipBins = append(ipBins, bin)
	}

	joinedIPDec := fmt.Sprintf("%08b%08b%08b%08b", ipBins[0], ipBins[1], ipBins[2], ipBins[3])

	return joinedIPDec, nil
}
