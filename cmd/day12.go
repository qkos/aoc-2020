package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use: "day12",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day12 called")
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
