package command

import (
	"github.com/spf13/cobra"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "show",
	Long:  "show",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCommand.AddCommand(showCommand)
}
