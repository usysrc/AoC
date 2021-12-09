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

var sums map[int]int

// gauss with memoization
func getSum(n int) int {
	if _, ok := sums[n]; !ok {
		sums[n] = n * (n + 1) / 2
	}
	return sums[n]
}

func main() {
	sums = make(map[int]int, 0)
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
			p := getSum(n)
			sum += p
			if sum >= min {
				continue
			}
		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min)
}
