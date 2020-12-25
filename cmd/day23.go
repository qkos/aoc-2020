package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use: "day23",

	Run: func(cmd *cobra.Command, args []string) {

		input := []int{1, 9, 8, 7, 5, 3, 4, 6, 2}
		fmt.Printf("Part 1: %v\n", d23(input, 100).Part1())
		for i := 9; i < 1000000; i++ {
			input = append(input, i+1)
		}
		fmt.Printf("Part 2: %v\n", d23(input, 10000000).Part2())
	},
}

type Node struct {
	Value      int
	Next, Prev *Node
}

type NodeList []*Node

func (nl NodeList) Contains(i int) bool {
	for _, n := range nl {
		if n.Value == i {
			return true
		}
	}
	return false
}

func (n *Node) NextN(s int) (arr NodeList) {
	last := n
	for i := 0; i < s; i++ {
		last = last.Next
		arr = append(arr, last)
	}
	return
}

func (n *Node) Part1() (s string) {
	for _, v := range n.NextN(8) {
		s += fmt.Sprintf("%d", v.Value)
	}
	return
}

func (n *Node) Part2() int {
	return n.Next.Value * n.Next.Next.Value
}

func (n *Node) Connect(o *Node) {
	if n == nil || o == nil {
		return
	}
	if n.Next != nil {
		n.Next.Prev = nil
	}
	if o.Prev != nil {
		o.Prev.Next = nil
	}
	n.Next = o
	o.Prev = n
}

func DecSize(i, size int) (d int) {
	d = i - 1
	if d == 0 {
		d = size
	}
	return
}

func d23(input []int, moves int) *Node {

	size := len(input)
	var start, last *Node
	nodes := map[int]*Node{}
	for _, v := range input {
		nodes[v] = &Node{v, nil, nil}
		if start == nil {
			start = nodes[v]
			last = start
		} else {
			last.Connect(nodes[v])
			last = last.Next
		}
	}
	last.Connect(start) // make it circular

	var p NodeList
	var d int
	c := start
	for i := 0; i < moves; i++ {
		// here we go

		p = c.NextN(3)
		for d = DecSize(c.Value, size); p.Contains(d); d = DecSize(d, size) {
		}

		// find d
		dn := nodes[d]
		// first detach p
		c.Connect(p[2].Next)
		// put p after dn
		p[2].Connect(dn.Next)
		dn.Connect(p[0])
		c = c.Next
	}
	return nodes[1]
}

func init() {
	rootCmd.AddCommand(day23Cmd)
}
