package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use: "day7",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		rules := map[string]map[string]int{}
		for _, line := range lines {
			spec := strings.Split(line, " bags contain ")
			left := spec[0]
			rep := strings.NewReplacer(" bags", "", " bag", "", ".", "")
			combinations := strings.Split(rep.Replace(spec[1]), ", ")

			combs := map[string]int{}
			for _, comb := range combinations {
				if comb == "no other" {
					// no other
					break
				}
				cp := strings.Split(comb, " ")
				mixc, _ := strconv.Atoi(cp[0])
				mix := strings.Join(cp[1:], " ")
				combs[mix] = mixc
			}
			rules[left] = combs

			// fmt.Printf("%s = %#v\n", left, combs)
		}
		ttl := calculateAll(rules)
		ttlSum := sum(map[string]int{}, rules, "shiny gold")
		fmt.Printf("count: %d, sum: %d\n", ttl, ttlSum)
	},
}

func sum(mem map[string]int, rules map[string]map[string]int, start string) (ttl int) {

	if s, ok := mem[start]; ok {
		return s
	}
	mp := rules[start]
	for k, v := range mp {
		if _, ok := mem[k]; !ok {
			mem[k] = sum(mem, rules, k)
		}
		ttl += v * (mem[k] + 1)
	}
	mem[start] = ttl
	return mem[start]
}

func calculateAll(rules map[string]map[string]int) (ttl int) {
	// given rules calculate
	for _, v := range rules {
		sum := calculate(rules, v)
		if sum > 0 {
			ttl++
		}
	}
	return
}

func calculate(rules map[string]map[string]int, combs map[string]int) (ttl int) {

	for k := range combs {
		if k == "shiny gold" {
			ttl += 1
		} else {
			ttl += calculate(rules, rules[k])
		}
	}
	return ttl
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
