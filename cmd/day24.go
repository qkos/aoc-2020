package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day24Cmd represents the day24 command
var day24Cmd = &cobra.Command{
	Use: "day24",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day24 called")
	},
}

func init() {
	rootCmd.AddCommand(day24Cmd)
}
