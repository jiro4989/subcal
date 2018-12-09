package main

import (
	"errors"
	"fmt"
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
	var cidraddr CIDRAddress
	ipcidr := strings.Split(s, "/")
	if len(ipcidr) < 2 {
		return cidraddr, errors.New("IP/CIDRが不正 ip/cidr=" + s)
	}

	ip, c := ipcidr[0], ipcidr[1]
	cidr, err := strconv.Atoi(c)
	if err != nil {
		msg := fmt.Sprintf("CIDRが不正 cidr=%s, err=%v", c, err)
		return cidraddr, errors.New(msg)
	}
	cidraddr.IPDec = ip
	cidraddr.CIDR = cidr

	joinedIPDec, err := IPDecToBinString(ip)
	if err != nil {
		return cidraddr, errors.New("IPが不正 ip=" + ip)
	}
	ipBin, err := strconv.ParseUint(joinedIPDec, 2, 64)
	if err != nil {
		return cidraddr, errors.New("IPが不正 joinedIPDec=" + joinedIPDec)
	}
	cidraddr.IPBin = ipBin

	net := joinedIPDec[:cidr]
	host := joinedIPDec[cidr:]
	netu, err := strconv.ParseUint(net, 2, 64)
	if err != nil {
		return cidraddr, err
	}
	hostu, err := strconv.ParseUint(host, 2, 64)
	if err != nil {
		return cidraddr, err
	}
	cidraddr.Network = netu
	cidraddr.Host = hostu

	return cidraddr, nil
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
