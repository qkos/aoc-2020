package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// day21Cmd represents the day21 command
var day21Cmd = &cobra.Command{
	Use: "day21",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		p1, p2 := d21(lines)
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	},
}

func d21(lines []string) (cnt int, alg string) {

	var ofoods []map[string]int
	var foods []map[string]int
	var aas []string
	aa := map[string][]int{}
	for l, line := range lines {
		parts := strings.Split(strings.NewReplacer("(", "", ")", "").Replace(line), " contains ")
		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1], ", ")
		mi, mo := map[string]int{}, map[string]int{}
		for _, i := range ingredients {
			mi[i]++
			mo[i]++
		}
		foods = append(foods, mi)
		ofoods = append(ofoods, mo)
		for _, a := range allergens {
			aa[a] = append(aa[a], l)
		}
	}

	// from highest to lowest
	for fi := len(foods); fi > 0; fi-- {
		var ac []string
		for a, fs := range aa {
			if len(fs) == fi {
				ac = append(ac, a)
				continue
			}
		}

		for _, a := range ac {
			fs := aa[a]
			// count all ingredients
			ic := map[string]int{}
			for _, f := range fs {
				for i := range foods[f] {
					ic[i]++
				}
			}
			// remove the item matching ing
			for i, c := range ic {
				if c == len(fs) {
					for _, f := range foods {
						delete(f, i)
					}
				}
			}
		}
	}

	// Begin Part 2
	for _, f := range foods {
		cnt += len(f)
		for i := range f {
			for _, of := range ofoods {
				delete(of, i)
			}
		}
	}

	mci := map[string][]string{}
	// from highest to lowest
	for fi := len(ofoods); fi > 0; fi-- {
		var ac []string
		for a, fs := range aa {
			if len(fs) == fi {
				ac = append(ac, a)
				continue
			}
		}

		for _, a := range ac {
			fs := aa[a]
			// count all ingredients
			ic := map[string]int{}
			for _, f := range fs {
				for i := range ofoods[f] {
					ic[i]++
				}
			}

			// remove the item matching ing
			for i, c := range ic {
				if c == len(fs) {
					mci[a] = append(mci[a], i)
				}
			}
		}
	}

	mcf := map[string]string{}
	for len(mcf) != len(mci) {
		for a, ings := range mci {
			if len(ings) == 1 {
				mcf[a] = mci[a][0]
				// remove this from all list
				for ia, ings := range mci {
					mci[ia] = removeStringItem(ings, mcf[a])
				}

				break
			}
		}
	}

	for a := range aa {
		aas = append(aas, a)
	}
	sort.Strings(aas)
	for i := 0; i < len(aas); i++ {
		aas[i] = mcf[aas[i]]
	}
	alg = strings.Join(aas, ",")

	return
}

func removeStringItem(list []string, item string) []string {
	for i, v := range list {
		if v == item {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func init() {
	rootCmd.AddCommand(day21Cmd)
}
