package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() int {
	w, h := 99, 99
	var height [99][99]int
	var visible [99][99]bool
	inputFile, err := os.Open("./input")

	// w, h := 5, 5
	// var height [5][5]int
	// var visible [5][5]bool
	// inputFile, err := os.Open("./test")
	if err != nil {
		panic(err)
	}

	j := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(j, " ")
		for k, v := range line {
			height[j][k], err = strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			fmt.Print(height[j][k])
			visible[j][k] = false
		}
		fmt.Println()
		j++
	}

	// test if tree is higher than the max before
	testTree := func(max, i, j int) int {
		if height[j][i] > max {
			max = height[j][i]
			visible[j][i] = true
		}
		return max
	}

	// scan horizontal
	for j := 0; j < h; j++ {
		max := -1
		for i := 0; i < w; i++ {
			max = testTree(max, i, j)
		}
		max = -1
		for i := w - 1; i >= 0; i-- {
			max = testTree(max, i, j)
		}
	}

	// scan vertical
	for i := 0; i < w; i++ {
		max := -1
		for j := 0; j < h; j++ {
			max = testTree(max, i, j)
		}
		max = -1
		for j := h - 1; j >= 0; j-- {
			max = testTree(max, i, j)
		}
	}

	sum := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if visible[j][i] {
				sum++
				// fmt.Print("x")
			} else {
				// fmt.Print("_")
			}
		}
		// fmt.Println()
	}

	return sum
}

func main() {
	count := readFile()
	fmt.Println("result is: ", count)
}
