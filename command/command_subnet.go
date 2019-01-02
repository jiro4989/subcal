package command

import (
	"fmt"
	"strings"

	"github.com/jiro4989/subcal/ip"
	"github.com/spf13/cobra"
)

func init() {
	subnetCommand.Flags().StringP("delimiter", "d", "\t", "color")
	subnetCommand.Flags().BoolP("color", "c", false, "color")
	subnetCommand.Flags().BoolP("ipv4", "i", false, "color")
	subnetCommand.Flags().BoolP("cidr", "r", false, "color")
	subnetCommand.Flags().BoolP("bin", "b", false, "color")
	subnetCommand.Flags().BoolP("mask", "m", false, "color")
	subnetCommand.Flags().BoolP("header", "H", false, "color")
	RootCommand.AddCommand(subnetCommand)
}

var subnetCommand = &cobra.Command{
	Use:   "subnet",
	Short: "subnet",
	Long:  "subnet",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("stdin")
			return
		}

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

		for _, ipcidr := range args {
			ipaddr, err := ip.ParseCIDR(ipcidr)
			if err != nil {
				panic(err)
			}

			fmt.Println(ipaddr.Format(sep, colorFlag, ipv4Flag, cidrFlag, binFlag, maskFlag))
		}
	},
}

func Header(sep string, ipv4Flag, cidrFlag, binFlag, maskFlag bool) string {
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
