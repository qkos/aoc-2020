package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use: "day20",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day20 called")
	},
}

func init() {
	rootCmd.AddCommand(day20Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day20Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day20Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
