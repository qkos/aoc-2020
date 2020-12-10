package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use: "day10",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		var jolts []int
		joltsMap := map[int]int{}
		for _, line := range lines {
			jolt, _ := strconv.Atoi(line)
			jolts = append(jolts, jolt)
			joltsMap[jolt] = 0
		}

		sort.Ints(jolts)
		mj := jolts[len(jolts)-1] // max
		jolts = append(jolts, mj+3)

		diff := map[int]int{1: 0, 2: 0, 3: 0}
		start := 0
		for _, j := range jolts {
			if j < start {
				continue
			}
			for k := range diff {
				if start+k == j {
					diff[k]++
					//fmt.Printf("%d choose %d - diff %d\n", start, j, k)
					start += k
				}
			}
		}

		cnt := comb(map[int]int{}, 0, 0, jolts)

		//fmt.Printf("Max: %d\n", mj)
		//fmt.Printf("Diffs: %#v\n", diff)
		fmt.Printf("part1: %d\n", diff[3]*diff[1])
		fmt.Printf("part2: %d\n", cnt)
	},
}

func comb(mem map[int]int, start, i int, ls []int) (count int) {
	if i == len(ls)-1 {
		return 1 // last one
	}
	for ; i < len(ls); i++ {
		v := ls[i]
		diff := v - start
		if diff > 3 {
			break
		}
		if diff > 0 && diff < 4 {
			if c, ok := mem[i]; ok {
				count += c
			} else {
				mem[i] = comb(mem, v, i, ls)
				count += mem[i]

			}
		}
	}
	return count
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
