package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line and parse the integers
	lines := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		strValues := strings.Fields(lineStr)
		intValues := make([]int, len(strValues))
		for i, str := range strValues {
			intValues[i], err = strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
		}
		lines = append(lines, intValues)
	}
	sumA := 0
	for _, line := range lines {
		if isLevelSafeA(line) {
			sumA++
		}
	}
	fmt.Println("Sum of A: ", sumA)

	sumB := 0
	for _, line := range lines {
		if isLevelSafeB(line) {
			sumB++
		}
	}
	fmt.Println("Sum of B: ", sumB)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func isLevelSafeA(intValues []int) bool {
	if len(intValues) <= 2 {
		return false
	}
	safe := true
	ascending := intValues[0] < intValues[len(intValues)-1]
	descending := intValues[0] > intValues[len(intValues)-1]
	if (!ascending && !descending) || (intValues[0] == intValues[1]) {
		safe = false
		return false
	}
	last := intValues[0]
	for _, v := range intValues[1:] {
		if v == last {
			safe = false
			break
		}
		if ascending && (v < last || v-last > 3) {
			safe = false
			break
		}
		if descending && (v > last || last-v > 3) {
			safe = false
			break
		}
		last = v
	}
	return safe
}

// if intValues is unsafe try to remove one value and check if it is safe
func isLevelSafeB(intValues []int) bool {
	if isLevelSafeA(intValues) {
		return true
	}
	// fmt.Println("unsafe ", intValues)
	for i := 0; i < len(intValues); i++ {
		// create a new slice without the value at index i
		values := make([]int, 0)
		for j := 0; j < len(intValues); j++ {
			if j != i {
				values = append(values, intValues[j])
			}
		}

		// fmt.Println("testing ", values)
		if isLevelSafeA(values) {
			// fmt.Println("safe! ", values)
			return true
		}
	}
	return false
}
