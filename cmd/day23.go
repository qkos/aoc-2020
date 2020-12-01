package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use: "day23",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day23 called")
	},
}

func init() {
	rootCmd.AddCommand(day23Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day23Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day23Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
