package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day24Cmd represents the day24 command
var day24Cmd = &cobra.Command{
	Use: "day24",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day24 called")
	},
}

func init() {
	rootCmd.AddCommand(day24Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day24Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day24Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
