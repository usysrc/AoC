package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(r rune) int {
	charValue := int(r)
	if charValue >= 96 {
		return charValue - 96
	}
	return 26 + charValue - 64
}

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	sum := 0
	for scanner.Scan() {
		m := make(map[rune]int)
		line := scanner.Text()
		for _, letter := range line[0 : len(line)/2] {
			m[letter] = 1
		}
		for _, letter := range line[len(line)/2:] {
			if _, ok := m[letter]; ok {
				sum += priority(letter)
				delete(m, letter)
			}
		}
	}
	return sum
}

func main() {
	sum := readFile()
	fmt.Println(sum)
}
