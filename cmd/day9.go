package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use: "day9",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day9 called")
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
