package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use: "day3",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		// part 2 slopes
		slopes := [][]int{
			[]int{1, 1},
			[]int{3, 1},
			[]int{5, 1},
			[]int{7, 1},
			[]int{1, 2},
		}

		var res []int
		max := 1

		for _, slope := range slopes {
			xslope := slope[0]
			yslope := slope[1]
			xpos := xslope
			treecount := 0
			for i, line := range lines {
				if i < yslope {
					continue
				}
				if i%yslope != 0 {
					continue
				}
				ll := len(line)
				// Implement this
				if string(line[xpos%ll]) == "#" {
					// free
					treecount++
				}
				xpos += xslope
			}
			res = append(res, treecount)
			max = max * treecount
		}

		fmt.Printf("Trees (part 1): %d\n", res[1])
		fmt.Printf("Trees: %d, Mul: %d\n", res, max)
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
