package cmd

import (
	"bufio"
	"os"
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
