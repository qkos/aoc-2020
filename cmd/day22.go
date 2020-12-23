package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day22Cmd represents the day22 command
var day22Cmd = &cobra.Command{
	Use: "day22",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Part 1: %d\n", d22(lines, false))
		fmt.Printf("Part 2: %d\n", d22(lines, true))
	},
}

func d22(lines []string, recurse bool) (score int) {

	var decks [2][]int
	var di int
	for _, line := range lines {
		if line == "" {
			di++
			continue
		} else if strings.HasPrefix(line, "Player") {
			continue
		}

		// do the thing
		c, _ := strconv.Atoi(line)
		decks[di] = append(decks[di], c)
	}

	// start game
	winner, decks := d22game2(decks, 1, recurse)
	for i, v := range decks[winner] {
		score += (len(decks[winner]) - i) * v
	}
	return
}

func d22game2(decks [2][]int, game int, recurse bool) (int, [2][]int) {

	// start game
	round := 1
	p1mem, p2mem := map[string]bool{}, map[string]bool{}
	var winner int
	for len(decks[0]) != 0 && len(decks[1]) != 0 {
		// do the thing
		p1h, p2h := fmt.Sprintf("%v", decks[0]), fmt.Sprintf("%v", decks[1])

		if _, ok := p1mem[p1h]; ok {
			return 0, decks
		}

		if _, ok := p2mem[p2h]; ok {
			return 1, decks
		}

		p1mem[p1h] = true
		p1mem[p2h] = true

		p1, p2 := decks[0][0], decks[1][0]

		decks[0] = decks[0][1:]
		decks[1] = decks[1][1:]

		// check if we have to play a sub game
		if recurse && len(decks[0]) >= p1 && len(decks[1]) >= p2 {
			// we have to play a sub game and determine winner
			winner, _ = d22game2([2][]int{
				append([]int(nil), decks[0][0:p1]...),
				append([]int(nil), decks[1][0:p2]...),
			}, game+1, recurse)
		} else if p1 > p2 {
			winner = 0
		} else {
			winner = 1
		}

		if winner == 0 {
			decks[0] = append(decks[0], p1, p2)
		} else {
			decks[1] = append(decks[1], p2, p1)
		}
		round++
	}

	return winner, decks
}

func init() {
	rootCmd.AddCommand(day22Cmd)
}
