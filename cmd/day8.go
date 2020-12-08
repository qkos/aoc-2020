package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use: "day8",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		// part 1
		acc, lst := run(lines)
		fmt.Printf("acc: %d - %d\n", acc, lst)

		// part 2
		re := strings.NewReplacer("nop", "jmp", "jmp", "nop")
		for l := 0; l < len(lines); l++ {

			line := lines[l]
			ins := string(line[0:3])

			if ins != "nop" && ins != "jmp" {
				continue
			}

			// replace the instruction
			nls := make([]string, len(lines))
			copy(nls, lines)
			nls[l] = re.Replace(line)
			sec, last := run(nls)

			//fmt.Printf("running %d -- [%s] = [%d, %d]\n", l, strings.Join(nls, ", "), sec, last)
			if last == len(lines)-1 {
				fmt.Printf("sec: %d - %d\n", sec, last)
			}
		}
	},
}

func run(lines []string) (acc, last int) {
	l := 0
	visited := map[int]bool{}
	for {

		if _, ok := visited[l]; ok || l >= len(lines) {
			// second run
			break
		}
		visited[l] = true
		last = l
		line := lines[l]
		ins := string(line[0:3])
		sig := string(line[4])
		val, _ := strconv.Atoi(string(line[5:]))
		//fmt.Printf("%s %s %d\n", ins, sig, val)
		switch ins {
		case "nop":
			l++
		case "acc":
			if sig == "+" {
				acc += val
			} else {
				acc -= val
			}
			l++
		case "jmp":
			if sig == "+" {
				l += val
			} else {
				l -= val
			}
		}
		//fmt.Printf("> %s | %d, %d\n", line, acc, l)
	}
	return
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
