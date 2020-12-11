package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use: "day11",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		var mp [][]string
		for _, line := range lines {
			//fmt.Printf("Input line %d -- %s\n", l, line)
			var ln []string
			for _, c := range line {
				cs := string(c)
				ln = append(ln, cs)
			}
			mp = append(mp, ln)
		}

		fmt.Printf("part1: %d\n", Adjacent(mp))
		fmt.Printf("part2: %d\n", FirstUp(mp))
	},
}

// PrintR debug yo
func PrintR(left [][]string) {
	for _, rows := range left {
		for _, v := range rows {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func FirstUp(prev [][]string) int {
	it := 0
	rows := len(prev)
	cols := len(prev[0])
	max := rows
	if cols > max {
		max = cols
	}
	for {

		//PrintR(prev)
		sit := it%2 == 0
		var next [][]string

		for i, row := range prev {
			var mpr []string
			for j, v := range row {
				if sit && v == "L" && count(max, i, j, prev) == 0 {
					mpr = append(mpr, "#")
				} else if !sit && v == "#" && count(max, i, j, prev) >= 5 {
					mpr = append(mpr, "L")
				} else {
					mpr = append(mpr, v)
				}
			}
			next = append(next, mpr)
		}
		it++
		//PrintR(next)
		//fmt.Println("----------------------------------")
		if eq(prev, next) {
			// done!
			return countOccupied(next)
		}
		prev = next

	}
}

func Adjacent(prev [][]string) int {
	it := 0
	for {

		//PrintR(prev)
		sit := it%2 == 0
		var next [][]string
		for i, row := range prev {
			var mpr []string
			for j, v := range row {
				if sit && v == "L" && count(1, i, j, prev) == 0 {
					mpr = append(mpr, "#")
				} else if !sit && v == "#" && count(1, i, j, prev) >= 4 {
					mpr = append(mpr, "L")
				} else {
					mpr = append(mpr, v)
				}
			}
			next = append(next, mpr)
		}
		it++
		//PrintR(next)
		//fmt.Println("----------------------------------")
		if eq(prev, next) {
			// done!
			return countOccupied(next)
		}
		prev = next

	}
}

func countOccupied(left [][]string) (total int) {
	for _, rows := range left {
		for _, v := range rows {
			if v == "#" {
				total++
			}
		}
	}
	return
}

func eq(left, right [][]string) bool {
	for i, rows := range left {
		for j, v := range rows {
			if v != right[i][j] {
				return false
			}
		}
	}
	return true
}

func count(max, i, j int, mp [][]string) (cnt int) {
	// check ha
	rows := len(mp)
	cols := len(mp[0])
	v := map[int]string{}
	for it := 1; it <= max && len(v) != 8; it++ {
		if _, ok := v[0]; !ok && i-it >= 0 && j-it >= 0 && mp[i-it][j-it] != "." {
			v[0] = mp[i-it][j-it]
		}
		if _, ok := v[1]; !ok && i-it >= 0 && mp[i-it][j] != "." {
			v[1] = mp[i-it][j]
		}
		if _, ok := v[2]; !ok && i-it >= 0 && j < cols-it && mp[i-it][j+it] != "." {
			v[2] = mp[i-it][j+it]
		}
		if _, ok := v[3]; !ok && j-it >= 0 && mp[i][j-it] != "." {
			v[3] = mp[i][j-it]
		}
		if _, ok := v[4]; !ok && j < cols-it && mp[i][j+it] != "." {
			v[4] = mp[i][j+it]
		}
		if _, ok := v[5]; !ok && i < rows-it && j-it >= 0 && mp[i+it][j-it] != "." {
			v[5] = mp[i+it][j-it]
		}
		if _, ok := v[6]; !ok && i < rows-it && mp[i+it][j] != "." {
			v[6] = mp[i+it][j]
		}
		if _, ok := v[7]; !ok && i < rows-it && j < cols-it && mp[i+it][j+it] != "." {
			v[7] = mp[i+it][j+it]
		}
	}

	for _, val := range v {
		if val == "#" {
			cnt++
		}
	}
	//fmt.Printf("[%d, %d, %d, %d] -- %#v\n", max, i, j, cnt, v)

	return
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
