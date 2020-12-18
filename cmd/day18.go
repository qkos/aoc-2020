package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use: "day18",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Part 1: %d\n", d18part1(lines))
		fmt.Printf("Part 2: %d\n", d18part2(lines))

	},
}

func d18part1(lines []string) (ttl int) {
	for _, line := range lines {
		ttl += d18cs(line)
	}
	return
}

func d18part2(lines []string) (ttl int) {
	for _, line := range lines {
		ttl += d18compute(d18normalise(strings.Split(strings.NewReplacer("(", "( ", ")", " )").Replace(line), " ")))
	}
	return
}

func d18normalise(xx []string) []string {

	sig := "+"
	var pos int
	for pos < len(xx) {

		// find the next +
		for ; pos < len(xx); pos++ {
			if xx[pos] == sig {
				break
			}
		}

		// no more sig
		if pos >= len(xx) {
			break
		}

		var nx []string
		left, right := d18wrap(d18cpr(xx[0:pos])), d18wrap(d18cp(xx[pos+1:]))

		nx = append(nx, xx[0:pos-len(left)]...)
		nx = append(nx, "(")
		nx = append(nx, left...)
		nx = append(nx, sig)
		nx = append(nx, right...)
		nx = append(nx, ")")
		nx = append(nx, xx[pos+len(right)+1:]...)

		pos += 2
		if len(left) > 1 {
			pos++
		}
		xx = nx

	}

	return xx
}

func d18wrap(xx []string) []string {
	if len(xx) == 1 {
		return xx
	}
	return append([]string{"("}, append(xx, ")")...)
}

func d18cpr(tokens []string) (res []string) {
	//fmt.Printf("cpr(%v)\n", tokens)
	lt := len(tokens)
	if tokens[lt-1] != ")" {
		return tokens[lt-1 : lt]
	}
	// find the next sub expression and remainder
	res = append(res, tokens[lt-1])
	c := 1
	for i := lt - 2; i >= 0 && c != 0; i-- {
		s := tokens[i]
		res = append([]string{tokens[i]}, res...)
		if s == ")" {
			// return
			c++
		} else if s == "(" {
			c--
		}
	}
	return res[1 : len(res)-1]
}

func d18cp(tokens []string) (res []string) {

	if tokens[0] != "(" {
		return tokens[0:1]
	}

	// find the next sub expression and remainder
	res = tokens[0:1]
	c := 1
	for i := 1; i < len(tokens) && c != 0; i++ {
		s := tokens[i]
		res = append(res, s)
		if s == "(" {
			// return
			c++
		} else if s == ")" {
			c--
		}
	}
	return res[1 : len(res)-1]
}

func d18cs(s string) int {
	return d18compute(strings.Split(strings.NewReplacer("(", "( ", ")", " )").Replace(s), " "))
}

func d18compute(tokens []string) int {
	var left int
	for i := 0; i < len(tokens); i++ {

		s := tokens[i]

		switch s {
		case "*", "+":
			// peek forward
			if tokens[i+1] == "(" {
				// we need to evaluate the rest
				right := d18cp(tokens[i+1:])
				left = d18calc(left, right, s)
				i += len(right) + 1 // fast forward
			} else {
				left = d18calc(left, []string{tokens[i+1]}, s)
				i++
			}
		case "(":
			rest := d18cp(tokens[i:])
			left = d18compute(rest)
			i += len(rest) // fast forward
		case ")":
			continue
		default:
			left, _ = strconv.Atoi(s)
		}
	}
	return left
}

func d18calc(l int, right []string, sig string) int {
	var r int
	if len(right) == 1 {
		r, _ = strconv.Atoi(right[0])
	} else {
		r = d18compute(right)
	}

	if sig == "*" {
		return l * r
	} else {
		return l + r
	}

}

func init() {
	rootCmd.AddCommand(day18Cmd)
}
