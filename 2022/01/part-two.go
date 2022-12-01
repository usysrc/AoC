package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(inputFile)

	max := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > 0 {
				max = append(max, sum)
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
	if sum > 0 {
		max = append(max, sum)
	}

	sort.Slice(max, func(i, j int) bool {
		return max[i] < max[j]
	})

	return max[len(max)-1] + max[len(max)-2] + max[len(max)-3]
}

func main() {
	max := readFile()
	fmt.Println(max)
}
