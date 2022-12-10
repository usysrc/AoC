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
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			fmt.Println(cycle, x, cycle*x)
			sum += cycle * x
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		instruction := parts[0]

		if instruction == "noop" {
			cycle++
			checkCycle()
		}
		if instruction == "addx" {
			operand, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			cycle++
			checkCycle()
			cycle++
			checkCycle()
			x += operand
		}
	}

	return sum
}

func main() {
	result := readFile()
	fmt.Println("Result is:", result)
}
