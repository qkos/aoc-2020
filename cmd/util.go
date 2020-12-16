package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileToLines(filename string) (out []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return
}

func ToInts(s, sep string) (out []int) {
	for _, part := range strings.Split(s, sep) {
		i, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("%s is not int", part))
		}
		out = append(out, i)
	}
	return
}
