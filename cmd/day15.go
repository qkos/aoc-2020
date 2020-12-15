package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use: "day15",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{0, 3, 6})) // 436
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{1, 3, 2})) // 1
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{2, 1, 3})) // 10
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{1, 2, 3})) // 27
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{2, 3, 1})) // 78
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{3, 2, 1})) // 438
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{3, 1, 2})) // 1836
		fmt.Printf("Part 1: %d\n", day15part1(2020, []int{9, 12, 1, 4, 17, 0, 18}))
		fmt.Printf("Part 2: %d\n", day15part1(30000000, []int{9, 12, 1, 4, 17, 0, 18}))
	},
}

func day15part1(lim int, in []int) int {

	visited := map[int]int{}
	for i, v := range in {
		visited[v] = i + 1
	}

	for i := len(in) - 1; i < lim; i++ {
		diff, v := 0, in[i]
		if last, ok := visited[v]; ok {
			diff = i + 1 - last
		}
		visited[v] = i + 1
		in = append(in, diff)
	}
	return in[lim-1]
}

func init() {
	rootCmd.AddCommand(day15Cmd)
}
