package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string][]string)
	files := os.Args[1:]

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, filename)
		f.Close()
	}

	for line, count := range counts {
		if count > 1 {

			fmt.Printf("%d \t %s \t %s \n", count, line, removeDuplicateStr(filename[line]))
		}
	}

}

func countLines(f *os.File, count map[string]int, filename map[string][]string) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		count[input.Text()]++
		filename[input.Text()] = append(filename[input.Text()], f.Name())

	}

}

func removeDuplicateStr(strSlice []string) []string {
	allkeys := make(map[string]bool)
	list := []string{}
	for _, fileName := range strSlice {

		{
			if _, present := allkeys[fileName]; !present {
				allkeys[fileName] = true
				list = append(list, fileName)

			}
		}

	}
	return list
}
