package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkLine(line string, rules map[string][]string) bool {
	lineOk := true

	// iterate over comma separated lines
	split := strings.Split(line, ",")
	for i, s := range split {
		// make sure that there is a rule for this number and the next
		k := i + 1
		if k < len(split) {
			found := false
			if targets, ok := rules[s]; ok {
				for _, target := range targets {
					if target == split[k] {
						found = true
					}
				}
				if !found {
					lineOk = false
				}
			}

			// line is invalid if there is a rule that matches other way around
			if _, ok := rules[split[k]]; ok {
				if ktargets, ok := rules[split[k]]; ok {
					for _, ktarget := range ktargets {
						if ktarget == s {
							lineOk = false
						}
					}
				}
			}
		}
	}
	return lineOk
}

func getMiddleNumber(line string) int {
	split := strings.Split(line, ",")
	// find middle element
	middle := len(split) / 2
	numStr := split[middle]
	if numStr == "" {
		panic("empty number")
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}

func a() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)

	// map a string to a slice of string
	rules := make(map[string][]string)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// if line contains a "|", split it
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			// rules[split[0]] = split[1]
			rules[split[0]] = append(rules[split[0]], split[1])
		} else if line != "" {
			lineOk := checkLine(line, rules)
			if lineOk {
				num := getMiddleNumber(line)
				sum += num
			}
		}

	}
	fmt.Println("Part A:", sum)
}

// order the line by the rules
func orderLine(line string, rules map[string][]string) []string {
	// find the order of the numbers according to the rules
	split := strings.Split(line, ",")
	ordered := make([]string, len(split))
	ordered[0] = split[0]
	for i, s := range split {
		k := i + 1
		if k < len(split) {
			if targets, ok := rules[s]; ok {
				for _, target := range targets {
					if target == split[k] {
						ordered[k] = split[k]
					}
				}
			}
		}
	}
	return ordered
}

func b() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)

	// map a string to a slice of string
	rules := make(map[string][]string)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// if line contains a "|", split it
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			// rules[split[0]] = split[1]
			rules[split[0]] = append(rules[split[0]], split[1])
		} else if line != "" {
			lineOk := checkLine(line, rules)
			if !lineOk {
				split := orderLine(line, rules)
				num := getMiddleNumber(strings.Join(split, ","))
				sum += num
			}
		}
	}
	fmt.Println("Part B:", sum)
}

func main() {
	a()
	// b is broken
	// b()
}
