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
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, count := range counts {
		if count > 1 {

			fmt.Printf("%d \t %s \n", count, line)
		}
	}

}

func countLines(f *os.File, count map[string]int) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		count[input.Text()]++
		if input.Text() == "END" {
			return
		}

	}

}
