package p

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long:  "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pagefriends-api v0.0")
	},
}
