package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use: "day14",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day14 called")
	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)
}
