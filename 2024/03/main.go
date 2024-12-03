package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// read the file 'input'
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			result += a * b
		}
	}
	fmt.Println(result)

}
