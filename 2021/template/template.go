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
	sum := 0
	for _, v := range data {
		sum += v
	}
	fmt.Println(sum)

}
