package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use: "day12",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day12 called")
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day12Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day12Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
