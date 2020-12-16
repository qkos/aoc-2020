package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Range [2]int

func (r Range) IsValid(v int) bool {
	return v >= r[0] && v <= r[1]
}

type Rule struct {
	Name   string
	Ranges []Range
}

func (r Rule) IsValid(v int) (valid bool) {
	for _, r := range r.Ranges {
		valid = valid || r.IsValid(v)
	}
	return
}

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use: "day16",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		var mtx []int
		var ntx [][]int
		var rules []Rule
		for l := 0; l < len(lines); l++ {
			if lines[l] == "" {
				continue
			} else if lines[l] == "your ticket:" {
				l++
				// parse
				mtx = ToInts(lines[l], ",")
			} else if lines[l] == "nearby tickets:" {
				for l = l + 1; l < len(lines); l++ {
					ntx = append(ntx, ToInts(lines[l], ","))
				}
			} else {
				parts := strings.Split(lines[l], ": ")
				rule := Rule{Name: parts[0]}
				for _, pair := range strings.Split(parts[1], " or ") {
					rng := strings.Split(pair, "-")
					left, _ := strconv.Atoi(rng[0])
					right, _ := strconv.Atoi(rng[1])
					rule.Ranges = append(rule.Ranges, Range{left, right})
				}
				rules = append(rules, rule)
			}
		}

		fmt.Printf("Part 1: %d\n", d16p1(rules, mtx, ntx))
		fmt.Printf("Part 2: %d\n", d16p2(rules, mtx, ntx))
	},
}

func d16p2(rules []Rule, mtx []int, ntx [][]int) (m int) {
	var fields []map[int][]string
	for _, tix := range ntx {
		flds := map[int][]string{}
		for i, v := range tix {
			valid := false
			for _, rule := range rules {
				rv := rule.IsValid(v)
				valid = valid || rv
				if rv {
					flds[i] = append(flds[i], rule.Name)
				}
			}
		}
		// correct
		if len(flds) == len(mtx) {
			// maybe correct
			fields = append(fields, flds)
		}
	}

	// time to get crazy
	var rf [][]string
	for i := 0; i < len(mtx); i++ {
		fc := map[string]int{}
		for _, mf := range fields {
			// count it
			for _, mff := range mf[i] {
				fc[mff]++
			}
		}
		var ifds []string
		for k, v := range fc {
			if v == len(fields) {
				ifds = append(ifds, k)
			}
		}
		rf = append(rf, ifds)
	}

	m = 1
	rff := map[string]int{}
	for i := 0; i < len(mtx); i++ {
		for k, v := range rf {
			if len(v) == i+1 {
				// we have a potential
				for _, f := range v {
					if _, ok := rff[f]; !ok {
						rff[f] = k
						if strings.HasPrefix(f, "departure ") {
							m *= mtx[k]
						}
					}
				}
			}
		}
	}

	return
}

func d16p1(rules []Rule, mtx []int, ntx [][]int) (e int) {
	for _, tix := range ntx {
		for _, v := range tix {
			valid := false
			for _, rule := range rules {
				valid = valid || rule.IsValid(v)
			}
			if !valid {
				e += v
			}
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(day16Cmd)
}
