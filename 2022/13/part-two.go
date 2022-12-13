package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

	var dividerPackerTwo any
	var dividerPackerSix any
	if err := json.Unmarshal([]byte("[[2]]"), &dividerPackerTwo); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte("[[6]]"), &dividerPackerSix); err != nil {
		panic(err)
	}
	signals := []any{}
	for _, pair := range pairs {
		signals = append(signals, pair.left)
		signals = append(signals, pair.right)
	}
	signals = append(signals, dividerPackerTwo, dividerPackerSix)

	sort.Slice(signals, func(i, j int) bool {
		return comparePair(Pair{left: signals[i], right: signals[j]}) < 0
	})

	decoderKey := 1
	for i, v := range signals {
		str, _ := json.Marshal(v)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	fmt.Println("Decoder Key is: ", decoderKey)
}
