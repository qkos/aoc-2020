package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day19Cmd represents the day19 command
var day19Cmd = &cobra.Command{
	Use: "day19",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day19 called")
	},
}

func init() {
	rootCmd.AddCommand(day19Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day19Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day19Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
