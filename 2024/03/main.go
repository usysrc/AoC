package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func openFile() *os.File {
	// read the file 'input'
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	partA()
	partB()
}

func partB() {
	file := openFile()
	defer file.Close()

	scannerB := bufio.NewScanner(file)
	resultB := 0
	do := true
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for scannerB.Scan() {
		line := scannerB.Text()
		matches := mulRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else {
				if do {
					resultB += atoi(match[1]) * atoi(match[2])
				}
			}
		}
	}
	fmt.Println(resultB)
}

func partA() {
	file := openFile()
	defer file.Close()

	scannerA := bufio.NewScanner(file)

	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	resultA := 0
	for scannerA.Scan() {
		line := scannerA.Text()
		matches := mulRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			resultA += atoi(match[1]) * atoi(match[2])
		}
	}
	fmt.Println(resultA)
}
