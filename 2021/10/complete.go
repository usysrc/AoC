package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const FILE = "input.txt"

func readFile() []string {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	lines := readFile()
	mirror := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	costsTable := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	costs := make([]int, 0)
	for _, line := range lines {
		stack := make([]string, 0)
		illegal := ""
		for _, r := range line {
			str := string(r)
			n := len(stack) - 1
			if str == "(" {
				stack = append(stack, str)
			}
			if str == ")" {
				if stack[n] != "(" {
					illegal = str
					break
				}
				stack = stack[:n]
			}

			if str == "[" {
				stack = append(stack, str)
			}
			if str == "]" {
				if stack[n] != "[" {
					illegal = str
					break
				}
				stack = stack[:n]
			}

			if str == "{" {
				stack = append(stack, str)
			}
			if str == "}" {
				if stack[n] != "{" {
					illegal = str
					break
				}
				stack = stack[:n]
			}

			if str == "<" {
				stack = append(stack, str)
			}
			if str == ">" {
				if stack[n] != "<" {
					illegal = str
					break
				}
				stack = stack[:n]
			}
		}
		if illegal == "" {
			c := 0

			for len(stack) > 0 {
				n := len(stack) - 1
				c *= 5
				c += costsTable[mirror[stack[n]]]
				stack = stack[:n]
			}
			costs = append(costs, c)

		}
	}
	sort.Ints(costs)
	idx := int(len(costs) / 2)
	fmt.Println(costs[idx])
}
