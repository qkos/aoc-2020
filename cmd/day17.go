package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use: "day17",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}
		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Part 1 : %d\n", d17part1(lines))
		fmt.Printf("Part 2 : %d\n", d17part2(lines))
	},
}

func d17part2(lines []string) int {

	cycle := 6
	start := len(lines)
	max := len(lines) + 2*6
	c := max * 2 // enough buffer
	dim := make([][][][]int, c)

	// initialize everything
	for i := 0; i < c; i++ {
		dim[i] = make([][][]int, c)
		for j := 0; j < c; j++ {
			dim[i][j] = make([][]int, c)
			for k := 0; k < c; k++ {
				dim[i][j][k] = make([]int, c)
			}
		}
	}

	z, w := 0, 0
	for x, line := range lines {
		for y := 0; y < len(line); y++ {
			var a int
			if line[y] == '#' {
				a = 1
			}
			dim[x+max][y+max][z+max][w+max] = a
		}
	}

	for i := 1; i <= cycle; i++ {

		next := map[string]int{}
		for x := max - i; x < (max + start + i); x++ {
			for y := max - i; y < (max + start + i); y++ {
				for z := max - i; z < (max + start + i); z++ {
					for w := max - i; w < (max + start + i); w++ {

						var ac int

						for x1 := x - 1; x1 < x+2; x1++ {
							for y1 := y - 1; y1 < y+2; y1++ {
								for z1 := z - 1; z1 < z+2; z1++ {
									for w1 := w - 1; w1 < w+2; w1++ {

										if x == x1 && y == y1 && z == z1 && w == w1 {
											continue
										}
										if dim[x1][y1][z1][w1] == 1 {
											ac++
										}
									}
								}
							}
						}

						//fmt.Printf("[%d - %d, %d, %d] -- ac: %d\n", i, x, y, z, ac)

						if dim[x][y][z][w] == 1 && (ac == 2 || ac == 3) {
							next[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)] = 1
						} else if dim[x][y][z][w] == 0 && ac == 3 {
							next[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)] = 1
						} else if dim[x][y][z][w] == 1 {
							next[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)] = 0
						}
					}
				}
			}
		}

		// make modification
		for x := max - i; x < (max + start + i); x++ {
			for y := max - i; y < (max + start + i); y++ {
				for z := max - i; z < (max + start + i); z++ {
					for w := max - i; w < (max + start + i); w++ {
						dim[x][y][z][w] = next[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)]
					}
				}
			}
		}
	}

	var ttl int
	for x := 0; x < c; x++ {
		for y := 0; y < c; y++ {
			for z := 0; z < c; z++ {
				for w := 0; w < c; w++ {
					ttl += dim[x][y][z][w]
				}
			}
		}
	}

	return ttl
}

func d17part1(lines []string) int {

	cycle := 6
	start := len(lines)
	max := len(lines) + 2*6
	c := max * 2 // enough buffer
	dim := make([][][]int, c)

	// initialize everything
	for i := 0; i < c; i++ {
		dim[i] = make([][]int, c)
		for j := 0; j < c; j++ {
			dim[i][j] = make([]int, c)
		}
	}

	z := 0
	for x, line := range lines {
		for y := 0; y < len(line); y++ {
			var a int
			if line[y] == '#' {
				a = 1
			}
			dim[x+max][y+max][z+max] = a
		}
	}

	for i := 1; i <= cycle; i++ {

		next := map[string]int{}
		for x := max - i; x < (max + start + i); x++ {
			for y := max - i; y < (max + start + i); y++ {
				for z := max - i; z < (max + start + i); z++ {

					var ac int

					for x1 := x - 1; x1 < x+2; x1++ {
						for y1 := y - 1; y1 < y+2; y1++ {
							for z1 := z - 1; z1 < z+2; z1++ {
								if x == x1 && y == y1 && z == z1 {
									continue
								}
								if dim[x1][y1][z1] == 1 {
									ac++
								}
							}
						}
					}

					//fmt.Printf("[%d - %d, %d, %d] -- ac: %d\n", i, x, y, z, ac)

					if dim[x][y][z] == 1 && (ac == 2 || ac == 3) {
						next[fmt.Sprintf("%d,%d,%d", x, y, z)] = 1
					} else if dim[x][y][z] == 0 && ac == 3 {
						next[fmt.Sprintf("%d,%d,%d", x, y, z)] = 1
					} else if dim[x][y][z] == 1 {
						next[fmt.Sprintf("%d,%d,%d", x, y, z)] = 0
					}
				}
			}
		}

		// make modification
		for x := max - i; x < (max + start + i); x++ {
			for y := max - i; y < (max + start + i); y++ {
				for z := max - i; z < (max + start + i); z++ {
					dim[x][y][z] = next[fmt.Sprintf("%d,%d,%d", x, y, z)]
				}
			}
		}
	}

	var ttl int
	for x := 0; x < c; x++ {
		for y := 0; y < c; y++ {
			for z := 0; z < c; z++ {
				ttl += dim[x][y][z]
			}
		}
	}

	return ttl
}

func init() {
	rootCmd.AddCommand(day17Cmd)
}
