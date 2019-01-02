package command

import (
	"fmt"

	"github.com/jiro4989/subcal/ip"
	"github.com/spf13/cobra"
)

func init() {
	showCommand.Flags().StringP("delimiter", "d", "\t", "color")
	showCommand.Flags().BoolP("color", "c", false, "color")
	showCommand.Flags().BoolP("ipv4", "i", false, "color")
	showCommand.Flags().BoolP("cidr", "r", false, "color")
	showCommand.Flags().BoolP("bin", "b", false, "color")
	showCommand.Flags().BoolP("mask", "m", false, "color")
	showCommand.Flags().BoolP("header", "H", false, "color")
	RootCommand.AddCommand(showCommand)
}

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "show",
	Long:  "show",
	Run: func(cmd *cobra.Command, args []string) {
		checkErr := func(err error) {
			if err != nil {
				panic(err)
			}
		}

		f := cmd.Flags()

		sep, err := f.GetString("delimiter")
		checkErr(err)

		colorFlag, err := f.GetBool("color")
		checkErr(err)

		ipv4Flag, err := f.GetBool("ipv4")
		checkErr(err)

		cidrFlag, err := f.GetBool("cidr")
		checkErr(err)

		binFlag, err := f.GetBool("bin")
		checkErr(err)

		maskFlag, err := f.GetBool("mask")
		checkErr(err)

		if !ipv4Flag && !cidrFlag && !binFlag && !maskFlag {
			ipv4Flag = true
			cidrFlag = true
			binFlag = true
			maskFlag = true
		}

		headerFlag, err := f.GetBool("header")
		checkErr(err)
		if headerFlag {
			fmt.Println(Header(sep, ipv4Flag, cidrFlag, binFlag, maskFlag))
		}

		const (
			maxBin  = 256
			maxCidr = 32
		)
		for i := 0; i < maxBin; i++ {
			for j := 0; j < maxBin; j++ {
				for k := 0; k < maxBin; k++ {
					for l := 0; l < maxBin; l++ {
						for cidr := 1; cidr <= maxCidr; cidr++ {
							s := fmt.Sprintf("%d.%d.%d.%d/%d", i, j, k, l, cidr)
							ipaddr, err := ip.ParseCIDR(s)
							if err != nil {
								panic(err)
							}
							fmt.Println(ipaddr.Format(sep, colorFlag, ipv4Flag, cidrFlag, binFlag, maskFlag))
						}
					}
				}
			}
		}
	},
}
