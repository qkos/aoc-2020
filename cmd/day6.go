package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use: "day6",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day6 called")
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
