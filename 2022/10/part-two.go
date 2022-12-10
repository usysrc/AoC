package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() int {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	cycle := 0
	x := 1
	sum := 0
	checkCycle := func() {
		if cycle%40 == 0 {
			fmt.Println()
		}
		if x >= (cycle%40)-1 && x <= (cycle%40)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		cycle++
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		instruction := parts[0]

		if instruction == "noop" {
			checkCycle()
		}
		if instruction == "addx" {
			operand, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			checkCycle()
			checkCycle()
			x += operand
		}
	}
	fmt.Println()

	return sum
}

func main() {
	result := readFile()
	fmt.Println("Result is:", result)
}
