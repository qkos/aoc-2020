package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// day19Cmd represents the day19 command
var day19Cmd = &cobra.Command{
	Use: "day19",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		p1, p2 := d19(lines)
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	},
}

func d19(lines []string) (int, int) {
	rules := map[string][][]string{}

	var inputs []string
	for _, line := range lines {

		if line == "" {
			continue
		}

		r := strings.Split(line, ": ")

		if len(r) == 1 {
			inputs = append(inputs, line)
			continue
		}

		seqs := strings.Split(r[1], " | ")
		for _, s := range seqs {
			parts := strings.Split(s, " ")
			var rseq []string
			for _, p := range parts {
				if p == `"a"` || p == `"b"` {
					rseq = append(rseq, strings.ReplaceAll(p, `"`, ""))
				} else {
					rseq = append(rseq, p)
				}
			}
			rules[r[0]] = append(rules[r[0]], rseq)
		}
	}

	mem := map[string][]string{}
	all := map[string]int{}
	rep := map[string]string{}
	var max, cl int
	for _, s := range d19matches(mem, "0", rules) {
		if len(s) > max {
			max = len(s)
		}
		all[s] = 1
	}

	// Handle
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	for _, r := range []string{"42", "31"} {
		for _, s := range mem[r] {
			rep[s] = r
			cl = len(s)
		}
	}

	var ttl, ttl2 int
	for _, line := range inputs {

		if _, ok := all[line]; ok {
			ttl++
			ttl2++
		} else if len(line)%cl == 0 {

			// potentially matching loopy rules
			chunks := len(line) / cl
			if chunks < 4 {
				continue // no hope
			}
			mc := make([]string, chunks)
			for i := 0; i < chunks; i++ {
				if rule, ok := rep[line[i*cl:(i+1)*cl]]; ok {
					mc[i] = rule
				}
			}

			var a, b int

			// count 42 forward
			for _, r := range mc {
				if r != "42" {
					break
				}
				a++
			}

			// count 31 backward
			for i := chunks - 1; i >= 0; i-- {
				if mc[i] != "31" {
					break
				}
				b++
			}

			if a > 2 && b > 0 && a > b && (a+b) == chunks {
				ttl2++
			}

		}
	}

	return ttl, ttl2
}

func d19matches(mem map[string][]string, rule string, rules map[string][][]string) []string {
	if m, ok := mem[rule]; ok {
		return m
	}
	// build all possibilities
	var mx []string
	for _, seq := range rules[rule] {
		var smx []string
		for _, s := range seq {
			// find s
			if _, ok := rules[s]; ok {
				// what to do with next??
				smx = d19cross(smx, d19matches(mem, s, rules))
			} else {
				// this is the value
				smx = d19cross(smx, seq)
			}
		}
		mx = append(mx, smx...)
	}
	mem[rule] = mx
	return mx
}

func d19cross(left, right []string) (x []string) {
	if len(right) == 0 {
		return left
	}
	if len(left) == 0 {
		return right
	}
	for _, l := range left {
		for _, r := range right {
			x = append(x, l+r)
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(day19Cmd)
}
