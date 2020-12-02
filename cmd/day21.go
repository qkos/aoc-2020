package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day21Cmd represents the day21 command
var day21Cmd = &cobra.Command{
	Use: "day21",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day21 called")
	},
}

func init() {
	rootCmd.AddCommand(day21Cmd)
}
