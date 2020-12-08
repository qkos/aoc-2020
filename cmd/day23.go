package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use: "day23",

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
	rootCmd.AddCommand(day23Cmd)
}
