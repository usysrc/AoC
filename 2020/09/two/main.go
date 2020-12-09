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

func calculateMinMax(lines []int, p int, q int, huge int) (int, int){
	min, max := huge, 0
	for i:=p; i<q; i++ {
		if lines[i] < min {
			min = lines[i]
		}
		if lines[i] > max {
			max = lines[i]
		}
	}
	return min, max
}

func linesToIntArray() []int {
	lines := make([]int, 0)

	for _, line := range readFile() {
		val, err := strconv.Atoi(line)
		if err != nil{
			log.Fatal(err)
		}
		lines = append(lines, val)
	}
	return lines
}

func main() {

	lines := linesToIntArray()

	p := 0
	q := 0
	target := 14360655
	value := 0

	for p < len(lines) {
		for ; q < len(lines) && value < target; q++ {
			value += lines[q]
		}
		if value == target {
			min, max := calculateMinMax(lines, p, q, target)
			fmt.Println(min+max)
			return
		}
		value -= lines[p]
		p++
	}
}