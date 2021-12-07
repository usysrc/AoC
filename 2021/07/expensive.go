package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

const FILE = "input.txt"

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
	min := 100000000

	max := 0
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	for i := 0; i < max; i++ {
		sum := 0
		for _, vv := range data {
			n := int(math.Abs(float64(vv - i)))
			p := 0
			for k := 0; k < n+1; k++ {
				p += k
			}
			sum += p

		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min)
}
