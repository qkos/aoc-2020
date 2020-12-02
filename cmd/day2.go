package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use: "day2",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		validCount := 0  // count for part 1
		validCount2 := 0 // count for part 2
		tok := regexp.MustCompile(`[:\- ]`)
		for _, line := range lines {
			parts := tok.Split(line, -1)
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			valid := parts[2]
			password := parts[4]
			cn := 0
			for _, c := range password {
				if string(c) == valid {
					cn++
				}
			}
			if cn >= min && cn <= max {
				validCount++
			}
			// check for part 2
			first := string(password[min-1]) == valid
			second := string(password[max-1]) == valid
			if (first || second) && !(first && second) {
				validCount2++
			}
		}

		fmt.Printf("Valid (part 1): %d\n", validCount)
		fmt.Printf("Valid (part 2): %d\n", validCount2)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
