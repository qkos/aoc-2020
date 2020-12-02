package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use: "day11",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day11 called")
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
