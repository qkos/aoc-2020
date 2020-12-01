package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use: "day10",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day10 called")
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
