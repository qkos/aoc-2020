package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use: "day18",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day18 called")
	},
}

func init() {
	rootCmd.AddCommand(day18Cmd)
}
