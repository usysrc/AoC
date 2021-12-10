package main

import (
	"bufio"
	"fmt"
	"os"
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
	costsTable := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	costs := 0
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
		if illegal != "" {
			costs += costsTable[illegal]
		}
	}
	fmt.Println(costs)

}
