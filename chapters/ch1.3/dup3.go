package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	files := os.Args[1:]
	count := make(map[string]int)

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stdout, "dup3: %v \n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}

	for line, count := range count {
		if count > 1 {

			fmt.Printf("%d:\t %s \n", count, line)
		}
	}

}
