package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use: "day16",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day16 called")
	},
}

func init() {
	rootCmd.AddCommand(day16Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day16Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day16Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
