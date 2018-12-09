package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

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
			var arr []string
			ipv4Addr, ipv4Net, err := net.ParseCIDR(ipcidr)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, ipv4Addr.String())

			ones, _ := ipv4Net.Mask.Size()
			arr = append(arr, fmt.Sprintf("%d", ones))

			l := len(ipv4Addr)
			ipbin := fmt.Sprintf("%08b%08b%08b%08b", ipv4Addr[l-4], ipv4Addr[l-3], ipv4Addr[l-2], ipv4Addr[l-1])
			arr = append(arr, ipbin)

			mask := fmt.Sprintf("%08b%08b%08b%08b", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])
			arr = append(arr, mask)

			joined := strings.Join(arr, "\t")
			fmt.Println(joined)
		}
	},
}

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().BoolP("color", "c", false, "color")
}
