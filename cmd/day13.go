package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use: "day13",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day13 called")
	},
}

func init() {
	rootCmd.AddCommand(day13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
