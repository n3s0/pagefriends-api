package p

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "pagefriends-api cli interface",
	Long:  "pagefriends-api cli interface",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}
