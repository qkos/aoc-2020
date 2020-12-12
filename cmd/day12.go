package cmd

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use: "day12",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("part 1 : %d\n", d12part1(lines))
		fmt.Printf("part 2 : %d\n", d12part2(lines))

	},
}

func d12part2(lines []string) int {
	sx, sy := 0, 0
	wx, wy := 10, 1

	re := regexp.MustCompile("(?P<l>[A-Z]{1})(?P<n>[0-9]+)")
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		a := parts[1]
		v, _ := strconv.Atoi(parts[2])
		//fmt.Printf("cmd: %s %d\n", a, v)
		switch a {
		case "N", "S", "E", "W":
			wx, wy = move(wx, wy, v, a)
		case "L":
			wx, wy = rot(float64(wx), float64(wy), float64(v))
		case "R":
			wx, wy = rot(float64(wx), float64(wy), -float64(v))
		case "F": // forward
			sx, sy = sx+v*wx, sy+v*wy
		}
		//fmt.Printf("S[%d, %d] W[%d, %d]\n\n", sx, sy, wx, wy)
	}

	return int(math.Abs(float64(sx)) + math.Abs(float64(sy)))
}

func rot(x, y, deg float64) (int, int) {
	r := deg * (math.Pi / 180)
	nx := math.Round(x*math.Cos(r) - y*math.Sin(r))
	ny := math.Round(x*math.Sin(r) + y*math.Cos(r))
	return int(nx), int(ny)
}

func d12part1(lines []string) int {
	dir := []string{"E", "S", "W", "N"}
	x, y := 0, 0
	d := 0

	re := regexp.MustCompile("(?P<l>[A-Z]{1})(?P<n>[0-9]+)")
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		a := parts[1]
		v, _ := strconv.Atoi(parts[2])
		//fmt.Printf("cmd: %s %d\n", a, v)
		switch a {
		case "N", "S", "E", "W":
			x, y = move(x, y, v, a)
		case "L":
			i := v / 90
			d = (4 + d - i) % 4
		case "R":
			i := v / 90
			d = (4 + d + i) % 4
		case "F": // forward
			x, y = move(x, y, v, dir[d])
		}
		//fmt.Printf("[%d, %d, %s, %d]\n\n", x, y, a, d)
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func move(x, y, v int, d string) (int, int) {
	//fmt.Printf("moving from [%d, %d, %s, %d]\n", x, y, d, v)
	switch d {
	case "N":
		y += v
	case "S":
		y -= v
	case "E":
		x += v
	case "W":
		x -= v
	}
	//fmt.Printf("arrive at [%d, %d]\n", x, y)
	return x, y
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
