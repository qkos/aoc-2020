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
}
