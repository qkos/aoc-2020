package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use: "day9",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		max := 25 // max
		var nums []int
		var invalid int
		for l, line := range lines {

			num, err := strconv.Atoi(line)
			if err != nil {
				panic("not a number")
			}

			// compute
			if l >= max {
				can := false
				last := nums[len(nums)-max:]
				valid := toMap(last)
				for _, v := range last {
					diff := num - v
					if _, ok := valid[diff]; ok {
						can = true
					}
				}
				if !can {
					invalid = num
				}
			}

			nums = append(nums, num)

		}

		minmax := 0
		for i := 0; i < len(nums)-1; i++ {
			sum := nums[i]
			min := sum
			max := sum
			// find a contiguous sum
			for j := i + 1; j < len(nums); j++ {
				sum += nums[j]
				if nums[j] < min {
					min = nums[j]
				}
				if nums[j] > max {
					max = nums[j]
				}
				if sum == invalid {
					minmax = min + max
					break
				}
			}
		}

		fmt.Printf("Part 1: %d\n", invalid)
		fmt.Printf("Part 2: %d\n", minmax)

	},
}

func toMap(list []int) map[int]int {
	mp := map[int]int{}
	for _, v := range list {
		mp[v] = 1
	}
	return mp
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
