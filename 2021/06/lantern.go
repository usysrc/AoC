package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const FILE = "input.txt"
const cycles = 80

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

func main() {
	data := readFile()
	for i := 0; i < cycles; i++ {
		add := 0
		for j, _ := range data {
			data[j] = data[j] - 1
			if data[j] < 0 {
				data[j] = 6
				add++
			}
		}
		if add > 0 {
			for j := 0; j < add; j++ {
				data = append(data, 8)
			}
		}
	}
	fmt.Println(len(data))
}
