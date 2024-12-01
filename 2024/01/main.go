package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	// read file 'input' and parse it, format is as follows:
	// number three spaces number
	// example: 12345   12345

	// lets read the file
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line
	left := make([]int, 0)
	right := make([]int, 0)
	for {
		var a, b int
		_, err := fmt.Fscanf(file, "%d   %d\n", &a, &b)
		if err != nil {
			break
		}
		fmt.Println(a, b)
		left = append(left, a)
		right = append(right, b)
	}
	// sort with the standard library
	sort.Ints(left)
	sort.Ints(right)

	// iterate over left
	sumA := 0
	for i := 0; i < len(left); i++ {
		// get the abs distance between left and right
		dist := int(math.Abs(float64(left[i] - right[i])))
		sumA += dist
	}
	fmt.Println("Answer A", sumA)

	// count occurences of each number in right
	counts := make(map[int]int)
	for _, v := range right {
		counts[v]++
	}

	sumB := 0
	for i := 0; i < len(left); i++ {
		if counts[left[i]] == 0 {
			continue
		}
		sumB += left[i] * counts[left[i]]
	}
	fmt.Println("Answer B", sumB)
}
