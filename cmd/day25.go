package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day25Cmd represents the day25 command
var day25Cmd = &cobra.Command{
	Use: "day25",

	Run: func(cmd *cobra.Command, args []string) {
		r := bfs(3248366)
		fmt.Printf("Final: %d\n", tf(4738476, r))
	},
}

func tf(sn, l uint64) uint64 {
	v := uint64(1)
	for i := 0; i < int(l); i++ {
		v *= sn
		v = v % uint64(20201227)
	}
	return v
}

func bfs(e uint64) uint64 {
	v := uint64(1)
	for i := uint64(1); ; i++ {
		v *= 7
		v = v % uint64(20201227)
		if v == e {
			return i
		}
	}
}

func init() {
	rootCmd.AddCommand(day25Cmd)
}
