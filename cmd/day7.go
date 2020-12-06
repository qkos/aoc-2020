package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use: "day7",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		for l, line := range lines {
			fmt.Printf("Input line %d -- %s\n", l, line)
		}
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
