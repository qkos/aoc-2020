package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day22Cmd represents the day22 command
var day22Cmd = &cobra.Command{
	Use: "day22",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day22 called")
	},
}

func init() {
	rootCmd.AddCommand(day22Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day22Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day22Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
