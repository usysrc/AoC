package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Pair struct {
	left  any
	right any
}

// read the file and generate pairs
func readFile() []Pair {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}
	pairs := []Pair{}
	pair := Pair{}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			pairs = append(pairs, pair)
			pair = Pair{}
		} else {
			var term any
			err := json.Unmarshal([]byte(line), &term)
			if err != nil {
				panic(err)
			}
			if pair.left == nil {
				pair.left = term
			} else {
				pair.right = term
			}
		}
	}
	return pairs
}

func comparePair(pair Pair) int {
	// both are numbers
	a, okA := pair.left.(float64)
	b, okB := pair.right.(float64)
	if okA && okB {
		return int(a) - int(b)
	}

	var leftList []any
	var rightList []any

	switch pair.left.(type) {
	case []any, []float64:
		leftList = pair.left.([]any)
	case float64:
		leftList = []any{pair.left}
	}

	switch pair.right.(type) {
	case []any, []float64:
		rightList = pair.right.([]any)
	case float64:
		rightList = []any{pair.right}
	}

	for i := range leftList {
		if len(rightList) <= i {
			return 1
		}
		if val := comparePair(Pair{left: leftList[i], right: rightList[i]}); val != 0 {
			return val
		}
	}
	if len(rightList) == len(leftList) {
		return 0
	}
	return -1
}

// sum the indexes of pairs in right order
func comparePairs(pairs []Pair) int {
	sum := 0
	for i, pair := range pairs {
		if comparePair(pair) <= 0 {
			// fmt.Println(pair, i+1)
			sum += i + 1
		}
	}
	return sum
}

func main() {
	pairs := readFile()
	sum := comparePairs(pairs)
	fmt.Println("result is:", sum)
}
