package command

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(showCommand)
}

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "show",
	Long:  "show",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
