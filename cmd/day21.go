package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day21Cmd represents the day21 command
var day21Cmd = &cobra.Command{
	Use: "day21",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day21 called")
	},
}

func init() {
	rootCmd.AddCommand(day21Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day21Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day21Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
