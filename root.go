package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "sop",
	Short: "sop is safety operation",
	Long:  "sop is safety operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root")
	},
}

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().BoolP("color", "c", false, "color")
}
