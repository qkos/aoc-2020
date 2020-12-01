package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use: "day15",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day15 called")
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day15Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day15Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
