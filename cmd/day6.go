package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use: "day6",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		sum := 0
		allAnswered := 0
		gs := 0
		mp := map[string]int{}
		for _, line := range lines {
			if line == "" {
				// reset
				sum += len(mp)
				allAnswered += all(mp, gs)
				mp = map[string]int{}
				gs = 0
				continue
			}

			// add to map
			for _, c := range line {
				cs := string(c)
				if _, ok := mp[cs]; !ok {
					mp[cs] = 1
				} else {
					mp[cs] += 1
				}
			}
			gs++
		}
		sum += len(mp)
		allAnswered += all(mp, gs)

		fmt.Printf("Sum: %d\n", sum)
		fmt.Printf("All: %d\n", allAnswered)
	},
}

func all(mp map[string]int, gs int) (ttl int) {
	for _, v := range mp {
		if v >= gs {
			ttl++
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
