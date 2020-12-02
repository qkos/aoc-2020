package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use: "day10",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day10 called")
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
