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

	onediff := 0
	threediff := 0

	n := 0
	for {
		if n == max {
			threediff++
			break
		}
		if m[n+1] {
			onediff++
			n = n+1
		} else if m[n+2] {
			n = n+2
		} else if m[n+3] {
			threediff++
			n = n+3
		}
	}

	fmt.Println(onediff*threediff)

}