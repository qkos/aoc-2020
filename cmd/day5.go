package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use: "day5",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		max := 0
		seats := [128][8]int{}
		for _, line := range lines {
			lastx := ""
			lasty := ""
			ymin := 0
			ymax := 128
			xmin := 0
			xmax := 8
			for _, c := range line {
				if string(c) == "F" {
					ymax -= (ymax - ymin) / 2
					lasty = string(c)
				} else if string(c) == "B" {
					ymin += (ymax - ymin) / 2
					lasty = string(c)
				} else if string(c) == "L" {
					xmax -= (xmax - xmin) / 2
					lastx = string(c)
				} else if string(c) == "R" {
					xmin += (xmax - xmin) / 2
					lastx = string(c)
				}
			}

			row := ymin
			col := xmin
			if lastx == "R" {
				col = xmax - 1
			} else if lasty == "B" {
				row = ymax - 1
			}

			id := row*8 + col
			if id > max {
				max = id
			}
			seats[row][col] = 1
			// debug
			// fmt.Printf("%d, %d, %d, %d = [%d, %d, %d]\n", ymin, ymax, xmin, xmax, row, col, id)
		}
		fmt.Printf("Max: %d\n", max)

		// find missing seat
		for i, rows := range seats {
			sum := 0
			missingCol := 0
			for j, col := range rows {
				if col == 0 {
					missingCol = j
				} else {
					sum += 1
				}
			}
			if sum == 7 {
				fmt.Printf("Missing [%d, %d] = id: %d\n", i, missingCol, i*8+missingCol)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
