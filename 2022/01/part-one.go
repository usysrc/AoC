package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(inputFile)
	max := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			number, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += number
		}

	}
	return max
}

func main() {
	max := readFile()
	fmt.Println(max)
}
