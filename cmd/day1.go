package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use: "day1",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		original := map[int]int{}
		for i, line := range lines {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
			}

			original[num] = 2020 - num
		}

		for k, diff := range original {
			if _, ok := original[diff]; ok {
				// we found a 2 solution
				fmt.Printf("Found pair - %d x %d = %d\n", k, diff, k*diff)
			} else {
				// find next diff
				for k1 := range original {
					// see if we get 0
					diff2 := diff - k1
					if _, ok := original[diff2]; ok {
						fmt.Printf("Found triple - %d x %d x %d = %d\n", k, k1, diff2, k*k1*diff2)
						break
					}
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)

}
