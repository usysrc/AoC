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
	group := []string{}
	for scanner.Scan() {
		m := make(map[rune]int)
		line := scanner.Text()
		group = append(group, line)

		if len(group) == 3 {
			for _, line := range group {
				// to make sure that every letter is only counted once per line we need a set
				set := make(map[rune]bool)

				for _, letter := range line {

					// check if in set and skip loop if it is
					_, isInSet := set[letter]
					if isInSet {
						continue
					}
					set[letter] = true

					// check if the letter occurs at least three times
					_, ok := m[letter]
					if ok && m[letter] == 2 {
						fmt.Println(string(letter))
						sum += priority(letter)
						delete(m, letter)
					} else if ok {
						m[letter]++
					} else if !ok {
						m[letter] = 1
					}

				}
			}
			group = []string{}
		}

	}
	return sum
}

func main() {
	sum := readFile()
	fmt.Println(sum)
}
