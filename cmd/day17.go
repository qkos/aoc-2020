package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use: "day17",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day17 called")
	},
}

func init() {
	rootCmd.AddCommand(day17Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day17Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day17Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
