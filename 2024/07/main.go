package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Concatenate a and b using arithmetic operations
func concatenate(a, b int64) int64 {
	temp := b

	// Handle the case when b is 0
	if temp == 0 {
		a *= 10
	} else {
		for temp > 0 {
			a *= 10
			temp /= 10
		}
	}
	return a + b
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	sumA := 0
	var sumB int64
	sumB = 0

	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, ":")
		if len(str) == 2 {
			left := strings.TrimSpace(str[0])
			right := strings.TrimSpace(str[1])
			var leftInt int
			fmt.Sscanf(left, "%d", &leftInt)
			rightSlice := strings.Split(right, " ")
			// convert rightSlice to []int
			rightIntSlice := make([]int, len(rightSlice))
			for i, v := range rightSlice {
				fmt.Sscanf(v, "%d", &rightIntSlice[i])
			}

			// Part A
			var calcA func(int, int) bool
			calcA = func(i int, current int) bool {
				if i >= len(rightIntSlice) {
					return current == leftInt
				}

				r := rightIntSlice[i]

				if calcA(i+1, r+current) {
					return true
				}

				if calcA(i+1, r*current) {
					return true
				}

				return false
			}
			res := calcA(0, 0)
			if res {
				sumA += leftInt
			}

			// Part B
			var calcB func(int64, int64) bool
			calcB = func(i int64, current int64) bool {
				if current > int64(leftInt) {
					return false
				}
				if i >= int64(len(rightIntSlice)) {
					return current == int64(leftInt)
				}

				r := int64(rightIntSlice[i])

				if calcB(i+1, r+current) {
					return true
				}

				if calcB(i+1, r*current) {
					return true
				}

				if calcB(i+1, concatenate(current, r)) {
					return true
				}

				return false
			}
			res = calcB(0, 0)
			if res {
				sumB += int64(leftInt)
			}

		}
	}
	fmt.Println("Part A:", sumA)
	fmt.Println("Part B:", sumB)
}
