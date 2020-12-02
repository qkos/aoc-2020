package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use: "day13",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day13 called")
	},
}

func init() {
	rootCmd.AddCommand(day13Cmd)
}
