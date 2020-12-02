package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use: "day9",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day9 called")
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
