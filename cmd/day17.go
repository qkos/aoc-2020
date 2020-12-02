package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use: "day17",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day17 called")
	},
}

func init() {
	rootCmd.AddCommand(day17Cmd)
}
