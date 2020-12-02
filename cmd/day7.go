package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use: "day7",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day7 called")
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
