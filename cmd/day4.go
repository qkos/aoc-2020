package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use: "day4",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		valid := map[string]string{
			"byr": "", //(Birth Year)
			"iyr": "", //(Issue Year)
			"eyr": "", //(Expiration Year)
			"hgt": "", //(Height)
			"hcl": "", //(Hair Color)
			"ecl": "", //(Eye Color)
			"pid": "", //(Passport ID)
			"cid": "", //(Country ID)
		}
		validcount := 0
		for _, line := range lines {
			//fmt.Printf("Line %d -- %s\n", i, line)

			if line == "" {
				if isValid(valid) {
					validcount++
				}
			} else {
				// parse and add to map
				tok := regexp.MustCompile(`[:\- ]`)
				parts := tok.Split(line, -1)
				for i := 0; i < len(parts); i += 2 {
					valid[parts[i]] = parts[i+1]
				}
			}
		}
		if isValid(valid) {
			validcount++
		}
		fmt.Printf("Valid: %d\n", validcount)
	},
}

func isValid(valid map[string]string) bool {
	validp := true
	for k, v := range valid {
		if v == "" && k != "cid" {
			validp = false
		}

		if k == "byr" {
			iv, err := strconv.Atoi(v)
			if err != nil || iv < 1920 || iv > 2002 {
				validp = false
			}
		}

		if k == "iyr" {
			iv, err := strconv.Atoi(v)
			if err != nil || iv < 2010 || iv > 2020 {
				validp = false
			}
		}

		if k == "eyr" {
			iv, err := strconv.Atoi(v)
			if err != nil || iv < 2020 || iv > 2030 {
				validp = false
			}
		}

		if k == "hgt" {
			matches := regexp.MustCompile("(?P<num>[0-9]+)(?P<unit>cm|in)").FindStringSubmatch(v)
			if len(matches) != 3 {
				validp = false
			} else {
				unit := matches[2]
				n, err := strconv.Atoi(matches[1])
				if err != nil {
					validp = false
				} else if unit == "cm" && (n < 150 || n > 193) {
					validp = false
				} else if unit == "in" && (n < 59 || n > 76) {
					validp = false
				}
			}

		}

		if k == "hcl" && !regexp.MustCompile("^#([0-9a-f]{6})$").MatchString(v) {
			validp = false
		}

		if k == "ecl" && !regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$").MatchString(v) {
			validp = false
		}

		if k == "pid" && !regexp.MustCompile("^[0-9]{9}$").MatchString(v) {
			validp = false
		}

		// reset that map
		valid[k] = ""
	}

	return validp
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
