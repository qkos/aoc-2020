package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use: "day23",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day23 called")
	},
}

func init() {
	rootCmd.AddCommand(day23Cmd)
}
