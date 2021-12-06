package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const FILE = "input.txt"
const cycles = 256

func readFile() []int {
	inputFile, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(inputFile).ReadAll()
	if err != nil {
		panic(err)
	}

	var data []int

	for _, line := range csvLines {
		for _, num := range line {
			number, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			data = append(data, number)
		}

	}
	return data
}

type Value struct {
	num    int
	amount int
}

func main() {
	data := readFile()
	counts := make(map[int]int, 0)

	for i := 0; i < 9; i++ {
		counts[i] = 0
	}

	for _, v := range data {
		counts[v]++
	}
	for i := 0; i < cycles; i++ {
		newcounts := make(map[int]int, 0)
		for j := 0; j < 9; j++ {
			newcounts[j] = 0
		}

		if counts[0] > 0 {
			newcounts[8] = counts[0]
			newcounts[6] = counts[0]
			counts[0] = 0
		}
		for j := 1; j < 9; j++ {
			newcounts[j-1] += counts[j]
		}
		counts = newcounts
	}
	sum := 0
	for j := 0; j < 9; j++ {
		sum += counts[j]
	}
	fmt.Println(sum)
}
