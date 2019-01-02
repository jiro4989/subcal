package command

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var RootCommand = &cobra.Command{
	Use:   "subcal",
	Short: "subcal",
	Long:  "subcal",
	Run:   func(cmd *cobra.Command, args []string) {},
}
