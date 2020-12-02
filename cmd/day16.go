package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use: "day16",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day16 called")
	},
}

func init() {
	rootCmd.AddCommand(day16Cmd)
}
