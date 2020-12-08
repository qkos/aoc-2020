package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use: "day14",

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
	rootCmd.AddCommand(day14Cmd)
}
