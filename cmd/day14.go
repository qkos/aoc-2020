package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use: "day14",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic("missing input file")
		}

		lines, err := FileToLines(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Part 1: %d\n", day14part1(lines))
		fmt.Printf("Part 2: %d\n", day14part2(lines))

	},
}

func day14part2(lines []string) uint64 {
	mem := map[uint64]uint64{}
	var mask string
	var set, masks []uint64
	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			mask = strings.Split(line, " = ")[1]
			set, masks = GetMask2(mask)
		} else if strings.HasPrefix(line, "mem[") {
			value, _ := strconv.Atoi(strings.Split(line, " = ")[1])
			addr, _ := strconv.Atoi(strings.Split(strings.NewReplacer("[", " ", "]", " ").Replace(line), " ")[1])
			// fmt.Printf("doing mem[%d] = %d + %v + %v\n", addr, value, set, masks)
			applyMasks(mem, applySetMask(uint64(addr), set, masks), uint64(value), masks)
		}
	}

	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func applySetMask(num uint64, set, mask []uint64) uint64 {
	for _, pos := range set {
		num |= 1 << pos
	}
	for _, pos := range mask {
		num &= ^(1 << pos)
	}
	return num
}

func applyMasks(mem map[uint64]uint64, num, val uint64, masks []uint64) uint64 {

	// set the bit before
	mem[num] = val

	// go through the masks
	for i, pos := range masks {
		// two possibilities [0, 1]
		// case of 1
		applyMasks(mem, applySetMask(num, []uint64{pos}, nil), val, masks[i+1:])
		// case of 0
		applyMasks(mem, num, val, masks[i+1:])
	}
	return num
}

func GetMask2(mask string) (set, masks []uint64) {
	for i, v := range mask {
		pos := uint64(len(mask)-i) - 1
		if v == '1' {
			set = append(set, pos)
		} else if v == 'X' {
			masks = append(masks, pos)
		}
	}
	return
}

func day14part1(lines []string) uint64 {
	mem := map[int]uint64{}
	var mask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			mask = strings.Split(line, " = ")[1]
		} else if strings.HasPrefix(line, "mem[") {
			value, _ := strconv.Atoi(strings.Split(line, " = ")[1])
			addr, _ := strconv.Atoi(strings.Split(strings.NewReplacer("[", " ", "]", " ").Replace(line), " ")[1])
			//fmt.Printf("doing mem[%d] = %d + %s\n", addr, value, mask)
			mem[addr] = Mask(uint64(value), mask)
		}
	}

	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func Mask(num uint64, mask string) uint64 {
	for i, v := range mask {
		pos := uint64(len(mask)-i) - 1
		if v == '1' {
			num |= 1 << pos
		} else if v == '0' {
			num &= ^(1 << pos)
		}
	}
	return num
}

func init() {
	rootCmd.AddCommand(day14Cmd)
}
