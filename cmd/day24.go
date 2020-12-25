package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day24Cmd represents the day24 command
var day24Cmd = &cobra.Command{
	Use: "day24",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		p1, p2 := d24(lines)
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	},
}

type Hex [2]int

func (h Hex) Move(d string) Hex {
	switch d {
	case "e": // right
		return Hex{h[0] + 1, h[1]}
	case "w": // left
		return Hex{h[0] - 1, h[1]}
	case "se":
		return Hex{h[0], h[1] + 1}
	case "sw":
		return Hex{h[0] - 1, h[1] + 1}
	case "ne":
		return Hex{h[0] + 1, h[1] - 1}
	case "nw":
		return Hex{h[0], h[1] - 1}
	}
	return h
}

func (h Hex) Adjacents() (a []Hex) {
	for _, d := range []string{"e", "w", "se", "sw", "ne", "nw"} {
		a = append(a, h.Move(d))
	}
	return
}

func (h Hex) Path(path string) Hex {
	nh := h
	for i := 0; i < len(path); i++ {
		if path[i] == 'e' || path[i] == 'w' {
			nh = nh.Move(path[i : i+1])
		} else {
			nh = nh.Move(path[i : i+2])
			i++
		}
	}
	return nh
}

type MapTile map[Hex]string

func (mt MapTile) Flip(h Hex) {
	if v, ok := mt[h]; ok && v == "b" {
		mt[h] = "w"
	} else {
		mt[h] = "b"
	}
}

func (mt MapTile) Count(w string) (c int) {
	for _, v := range mt {
		if v == w {
			c++
		}
	}
	return
}

func (mt MapTile) Val(h Hex) string {
	if v, ok := mt[h]; ok && v == "b" {
		return "b"
	}
	return "w"
}

func d24(lines []string) (p1, p2 int) {

	start := Hex{}
	tile := MapTile{}

	// part 1
	for _, line := range lines {
		nt := start.Path(line)
		tile.Flip(nt)
	}
	p1 = tile.Count("b")

	// part 2
	for i := 0; i < 100; i++ {

		cf := MapTile{}
		mwf := MapTile{}
		// determine black tiles to flip
		for h, v := range tile {
			if v == "b" {
				var cnt int

				for _, nh := range h.Adjacents() {
					if tile.Val(nh) == "b" {
						cnt++
					} else {
						// we have a white tile, check maybe
						mwf[nh] = "w"
					}
				}
				if cnt == 0 || cnt > 2 {
					cf[h] = v
				}
			}
		}
		for w, wv := range mwf {
			var wcnt int
			for _, nh := range w.Adjacents() {
				if tile.Val(nh) == "b" {
					wcnt++
				}
			}
			if wcnt == 2 {
				cf[w] = wv
			}
		}
		for h := range cf {
			tile.Flip(h)
		}
	}

	p2 = tile.Count("b")

	return
}

func init() {
	rootCmd.AddCommand(day24Cmd)
}
