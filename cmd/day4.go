package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use: "day4",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day4 called")
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
