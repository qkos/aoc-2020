package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day25Cmd represents the day25 command
var day25Cmd = &cobra.Command{
	Use: "day25",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day25 called")
	},
}

func init() {
	rootCmd.AddCommand(day25Cmd)

}
