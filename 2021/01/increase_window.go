package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readFile() []int {
	inputFile, err := os.Open("./input.txt")
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
	if len(data) <= 1 {
		return
	}
	increase := 0
	last := 0
	for i := 0; i < len(data); i++ {
		sum := 0
		for k := 0; k < 3; k++ {
			if i+k < len(data) {
				sum += data[i+k]
			}
		}
		if sum > last && last != 0 {
			increase++
		}
		last = sum
		fmt.Println(sum, increase)
	}
	fmt.Println(increase)
}
