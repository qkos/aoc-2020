package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use: "day20",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		p1, p2 := d20(lines)
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	},
}

const tz = 9
const mz = 12
const mz10 = mz * 10
const mz8 = mz * 8

type TilesMap map[int]Tile

func (tm TilesMap) Dump() string {
	x, _ := json.MarshalIndent(tm, "", "  ")
	return string(x)
}

func (tm TilesMap) IDs() (ts []int) {
	for t := range tm {
		ts = append(ts, t)
	}
	return
}

func (tm TilesMap) Tiles() (ts Tiles) {
	for _, t := range tm {
		ts = append(ts, t)
	}
	return
}

type Tiles []Tile

type TileImage [10][10]rune

func (ti TileImage) Rotate(k int) TileImage {
	if k == 0 {
		return ti
	}
	nti := ti.Rotate90() // rotate once
	for i := 2; i < k; i++ {
		nti = nti.Rotate90()
	}
	return nti
}

func (ti TileImage) Rotate90() (nti TileImage) {
	for i := 0; i < tz+1; i++ {
		for j := 0; j < tz+1; j++ {
			nti[i][j] = ti[tz-j][i]
		}
	}
	return
}

func (ti TileImage) FlipUp() (nti TileImage) {
	for i := 0; i < (tz+1)/2; i++ {
		nti[i] = ti[tz-i]
		nti[tz-i] = ti[i]
	}
	return
}

func (ti TileImage) FlipSides() (nti TileImage) {
	for i := 0; i < tz+1; i++ {
		for j := 0; j < (tz+1)/2; j++ {
			nti[i][j] = ti[i][tz-j]
			nti[i][tz-j] = ti[i][j]
		}
	}
	return
}

type FullImage [mz8][mz8]rune

func (ti FullImage) Rotate(k int) FullImage {
	if k == 0 {
		return ti
	}
	nti := ti.Rotate90() // rotate once
	for i := 2; i < k; i++ {
		nti = nti.Rotate90()
	}
	return nti
}

func (ti FullImage) Rotate90() (nti FullImage) {
	for i := 0; i < mz8; i++ {
		for j := 0; j < mz8; j++ {
			nti[i][j] = ti[mz8-j-1][i]
		}
	}
	return
}

func (ti FullImage) FlipUp() (nti FullImage) {
	for i := 0; i < (mz8)/2; i++ {
		nti[i] = ti[mz8-i-1]
		nti[mz8-i-1] = ti[i]
	}
	return
}

func (ti FullImage) FlipSides() (nti FullImage) {
	for i := 0; i < mz8; i++ {
		for j := 0; j < mz8/2; j++ {
			nti[i][j] = ti[i][mz8-j-1]
			nti[i][mz8-j-1] = ti[i][j]
		}
	}
	return
}

func (ti FullImage) All() []FullImage {
	f := []FullImage{ti, ti.FlipSides(), ti.FlipUp()}
	for i := 0; i < 3; i++ {
		tr := ti.Rotate(i + 1)
		f = append(f, tr, tr.FlipUp(), tr.FlipSides())
	}
	return f
}

type Tile struct {
	Id    int
	Sides []int
	Image TileImage
}

func (t Tile) String() string {
	return fmt.Sprintf("%d: %v", t.Id, t.Sides)
}

func (t Tile) Rotate(k int) Tile {
	if k == 0 {
		return t
	}
	// initial state
	next := append([]int{}, t.Sides...)
	for i := 1; i < k; i++ {
		next = []int{d20rev(next[3]), next[0], d20rev(next[1]), next[2]}
	}
	return Tile{t.Id, next, t.Image.Rotate(k)}
}

func (t Tile) FlipUp() Tile {
	ts := t.Sides
	return Tile{t.Id, []int{ts[2], d20rev(ts[1]), ts[0], d20rev(ts[3])}, t.Image.FlipUp()}
}

func (t Tile) FlipSides() Tile {
	ts := t.Sides
	return Tile{t.Id, []int{d20rev(ts[0]), ts[3], d20rev(ts[2]), ts[1]}, t.Image.FlipSides()}
}

// Match number of matches
func (t Tile) PMatches(o Tile) (m int) {

	for _, s := range t.Sides {
		for _, os := range o.Sides {
			if s == os {
				m++
			}
		}
	}
	return m
}

func (t Tile) All() (v Tiles) {
	v = Tiles{t, t.FlipSides(), t.FlipUp()}
	for i := 0; i < 3; i++ {
		tr := t.Rotate(i + 1)
		v = append(v, tr, tr.FlipUp(), tr.FlipSides())
	}
	return v
}

func (t Tile) MatchAll(tiles TilesMap) TileMatch {
	var mc Tiles
	for id, ot := range tiles {
		if t.Id == id {
			continue // skip same tile
		}
		for _, it := range []Tile{ot, ot.FlipSides(), ot.FlipUp()} {
			if t.PMatches(it) > 0 {
				mc = append(mc, it)
			}
		}
	}
	return TileMatch{t, mc}
}

// Exact fit
func (t Tile) Fit(o Tile) (mp MatchPos) {
	return MatchPos{
		t.Sides[0] == o.Sides[2],
		t.Sides[1] == o.Sides[3],
		t.Sides[2] == o.Sides[0],
		t.Sides[3] == o.Sides[1],
	}
}

func (t Tile) FitAll(tiles TilesMap) TileMatch {
	tm := TileMatch{Tile: t}
	for id, ot := range tiles {
		if t.Id == id {
			continue // skip same tile
		}
		for _, tot := range ot.All() {
			if t.Fit(tot).Fit() {
				tm.Matches = append(tm.Matches, tot)
			}
		}
	}
	return tm
}

type TileMatch struct {
	Tile    Tile
	Matches Tiles
}

type MatchPos [4]bool

func (mp MatchPos) Fit() bool {
	return mp[0] || mp[1] || mp[2] || mp[3]
}

// get first unique tiles
func (ts Tiles) Uniq() (tc TilesMap) {
	tc = TilesMap{}
	for _, t := range ts {
		if _, ok := tc[t.Id]; !ok {
			tc[t.Id] = t
		}
	}
	return
}

func d20rev(n int) int {
	var r int
	for i := tz; i >= 0; i-- {
		if n&(1<<(tz-i)) != 0 {
			r |= 1 << i
		}
	}
	return r
}

func D20PrintRunes10(rr [mz10][mz10]rune) {
	for i, row := range rr {
		if i%10 == 0 {
			fmt.Println()
		}
		for j, r := range row {
			if j%10 == 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func D20PrintRunes8(rr [mz8][mz8]rune) {
	for _, row := range rr {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func D20NewTile(id int, lines []string) Tile {
	var a, b, c, d int
	var sides []int
	var image TileImage
	for i, line := range lines {
		for j, s := range line {
			image[i][j] = s
			// top
			if i == 0 && s == '#' {
				a |= 1 << (tz - j)
			}

			// right
			if j == tz && s == '#' {
				b |= 1 << (tz - i)
			}

			// bottom
			if i == tz && s == '#' {
				c |= 1 << (tz - j)
			}

			// left
			if j == 0 && s == '#' {
				d |= 1 << (tz - i)
			}

		}
	}
	sides = append(sides, a, b, c, d)

	return Tile{id, sides, image}
}

func d20(lines []string) (mul, ttl int) {

	tiles := TilesMap{}

	// tiles always 10 by 10
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		re := regexp.MustCompile("Tile (?P<id>[0-9]+):")
		if strings.HasPrefix(line, "Tile ") {
			id, _ := strconv.Atoi(re.FindStringSubmatch(line)[1])
			tiles[id] = D20NewTile(id, lines[i+1:i+11])
			i += 9
		}
	}

	// find 4 corners
	mul = 1
	var cornerTiles []Tile
	mc := map[int][]int{}
	for id, t := range tiles {
		for _, it := range t.All() {
			tm := it.MatchAll(tiles)
			res := tm.Matches.Uniq().IDs()
			if len(res) == 2 {
				mc[id] = res
				mul *= id
				cornerTiles = append(cornerTiles, tiles[id].All()...)
				//fmt.Printf("Maybe corner: %d\n", id) // DEBUG
				break
			}
		}
	}

	var fmf [mz][mz]Tile
corner:
	for _, fm := range cornerTiles {

		//fmt.Printf("--- Trying corner %d\n", fm.Id) // DEBUG
		var fml [mz][mz]Tile
		fml[0][0] = fm // this is the first tile

		for i := 0; i < mz-1; i++ {
			for j := 0; j < mz-1; j++ {

				t := fml[i][j]
				tfm := t.FitAll(tiles)

				var f1, f2 bool
				for _, tm := range tfm.Matches {
					fit := t.Fit(tm)
					if !f1 && fit[1] && (i == 0 || i > 0 && tm.Fit(fml[i-1][j+1]).Fit()) {
						fml[i][j+1] = tm
						f1 = true
					} else if !f2 && fit[2] && (j == 0 || j > 0 && tm.Fit(fml[i+1][j-1]).Fit()) {
						fml[i+1][j] = tm
						f2 = true
					}
					if f1 && f2 {
						break
					}
				}
				if !f1 || !f2 {
					continue corner // we didn't find any match
				}
			}
		}

		// get the last tile
		lt := fml[mz-2][mz-1]
		for _, tm := range lt.FitAll(tiles).Matches {
			fit := lt.Fit(tm)
			if fit[2] && tm.Fit(fml[11][10]).Fit() {
				fml[mz-1][mz-1] = tm
				break
			}
		}

		if fml[mz-1][mz-1].Id != 0 {
			// we got to the last corner
			fmf = fml
			break
		}

	}

	//for i := 0; i < mz; i++ { // DEBUG
	//	for j := 0; j < mz; j++ { // DEBUG
	//		fmt.Printf("%s | ", fmf[i][j].String()) // DEBUG
	//	} // DEBUG
	//	fmt.Println() // DEBUG
	//} // DEBUG

	var image [mz10][mz10]rune
	var imaget FullImage
	var smh, smh8 int

	for i := 0; i < mz; i++ {
		for j := 0; j < mz; j++ {

			ti := fmf[i][j].Image

			for k := 0; k < tz+1; k++ {
				for l := 0; l < tz+1; l++ {
					// fill the fkn image
					image[i*(tz+1)+k][j*(tz+1)+l] = ti[k][l]
					if k > 0 && k < 9 && l > 0 && l < 9 {
						imaget[i*(tz-1)+k-1][j*(tz-1)+l-1] = ti[k][l]
						if ti[k][l] == '#' {
							smh8++
						}
					}
					if ti[k][l] == '#' {
						smh++
					}
				}
			}
		}
	}

	//D20PrintRunes10(image)                 // DEBUG
	//fmt.Println()                       // DEBUG
	//fmt.Println("After border removed") // DEBUG
	//fmt.Println()                       // DEBUG
	//D20PrintRunes8(imaget)                 // DEBUG
	//fmt.Println()                       // DEBUG

	// now find the sea monster [3][20], 15 #s
	sm := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	// kill me now
	for _, im := range imaget.All() {
		var smc int
		for i := 0; i < mz8-3; i++ {
			for j := 0; j < mz8-20; j++ {
				// find matches
				var smm int
				for k, smr := range sm {
					for l, smc := range smr {
						if smc == '#' && im[i+k][j+l] == smc {
							smm++
						}
					}
				}
				if smm == 15 {
					// made it here
					smc++
				}
			}
		}

		if smc == 0 {
			continue
		}

		ttl = smh8 - (15 * smc)
		//fmt.Printf("Found %d sea monsters\n", smc)          // DEBUG
		//fmt.Printf("Hash in image: %d\n", smh8)             // DEBUG
		//fmt.Printf("%d - (%d * 15) = %d\n", smh8, smc, ttl) // DEBUG
		break
	}

	return
}

func init() {
	rootCmd.AddCommand(day20Cmd)
}
