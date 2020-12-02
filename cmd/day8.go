package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use: "day8",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day8 called")
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
