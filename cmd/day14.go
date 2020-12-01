package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use: "day14",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day14 called")
	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day14Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day14Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
