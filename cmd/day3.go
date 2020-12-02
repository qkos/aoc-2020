package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use: "day3",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day3 called")
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
