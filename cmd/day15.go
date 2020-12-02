package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use: "day15",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day15 called")
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)
}
