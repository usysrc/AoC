package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 4; i < len(line); i++ {
			str := line[i-4 : i]
			m := make(map[rune]bool)
			found := true
			for _, c := range str {
				if _, ok := m[c]; ok {
					found = false
					break
				}
				m[c] = true
			}
			if found {
				return i
			}
		}
	}
	return -1
}

func main() {
	position := readFile()
	fmt.Println(position)
}
