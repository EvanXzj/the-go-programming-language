// Dup2 prints the count and text of lines that appear more than once
// the the input. It read from stdin or form a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)

				continue
			}

			countLines(f, counts)

			defer f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// map 是指针传递还是传递的副本， 会修改原counts里面的值吗？
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	if input.Scan() {
		counts[input.Text()]++
	}
}
