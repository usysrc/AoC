package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func main() {
	numbers := make([]int, 0)
	for _, line := range readFile() {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	m := make(map[int]bool)
	max := 0
	for _, num := range numbers {
		m[num] = true
		if num > max {
			max = num
		}
	}

	solutionsFound := 0

	cache := make(map[int]int)

	var find func(i int)
	find = func(i int){
		if i >= max + 3 {
			solutionsFound++
		}
		if m[i] {
			for k:=1; k<4; k++ {
				if val, ok := cache[i+k]; ok {
					solutionsFound += val
				} else {
					oldFounds := solutionsFound
					find(i+k)
					cache[i+k] = solutionsFound - oldFounds
				}
			}

		} else {
			return
		}
	}

	for k:=0; k<4; k++ {
		find(k)
	}

	fmt.Println(solutionsFound)
}