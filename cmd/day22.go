package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day22Cmd represents the day22 command
var day22Cmd = &cobra.Command{
	Use: "day22",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day22 called")
	},
}

func init() {
	rootCmd.AddCommand(day22Cmd)
}
