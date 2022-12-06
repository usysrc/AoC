package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func parseRange(s string) (a, b int) {
	_, err := fmt.Sscanf(s, "%d-%d", &a, &b)
	if err != nil {
		panic(err)
	}
	return a, b
}

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(inputFile)
	reader.Comma = ','
	csvLines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	count := 0
	for _, line := range csvLines {
		a, b := parseRange(line[0])
		c, d := parseRange(line[1])
		if (a <= c && b >= d) || (c <= a && d >= b) {
			fmt.Println(a, b, c, d)
			count++
		}
	}
	return count
}

func main() {
	contained := readFile()
	fmt.Println(contained)
}
