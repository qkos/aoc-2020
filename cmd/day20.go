package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use: "day20",

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
	rootCmd.AddCommand(day20Cmd)
}
