package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use: "day20",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day20 called")
	},
}

func init() {
	rootCmd.AddCommand(day20Cmd)
}
