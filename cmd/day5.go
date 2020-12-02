package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use: "day5",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day5 called")
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
