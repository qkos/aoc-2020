package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use: "day13",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		var et int
		var bs []uint64
		var bsi []int
		for l, line := range lines {
			if l == 0 {
				// first line
				et, _ = strconv.Atoi(line)
			} else {
				for _, b := range strings.Split(line, ",") {
					if b != "x" {
						bi, _ := strconv.Atoi(b)
						bs = append(bs, uint64(bi))
						bsi = append(bsi, bi)
					} else {
						bs = append(bs, 0)
						bsi = append(bsi, 0)
					}
				}
			}
		}
		fmt.Printf("Part 1: %d\n", day13part1(et, bsi))

		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{7, 13, 0, 0, 59, 0, 31, 19}))
		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{17, 0, 13, 19}))
		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{67, 7, 59, 61}))
		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{67, 0, 7, 59, 61}))
		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{67, 7, 0, 59, 61}))
		//fmt.Printf("Part 2 test: %d\n", day13part2([]uint64{1789, 37, 47, 1889}))
		fmt.Printf("Part 2: %d\n", day13part2(bs))

	},
}

type Bus struct {
	No    uint64
	Index uint64
}

func day13part2(buses []uint64) uint64 {
	t, w := uint64(0), uint64(1)
	for i, b := range buses {
		if b != 0 {
			for i64 := uint64(i); (t+i64)%b != 0; t += w {
			}
			w *= b
		}
	}
	return t
}

func day13part1(time int, buses []int) int {
	sort.Ints(buses)
	minb, minv := 0, 0
	for _, b := range buses {
		if b == 0 {
			continue
		}
		x := time / b
		next := b * (x + 1)
		if minb == 0 || next < minv {
			minb, minv = b, next
		}
	}
	return minb * (minv - time)
}

func init() {
	rootCmd.AddCommand(day13Cmd)
}
