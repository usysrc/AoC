package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	DEFINITIONS = iota
	ORDERS
)

func PartA() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	state := DEFINITIONS
	definitions := map[string]bool{}
	possible := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if line == "" {
			state = ORDERS
			continue
		}
		if state == DEFINITIONS {
			// Parse the definitions
			def := strings.Split(line, ", ")
			for _, d := range def {
				definitions[d] = true
				fmt.Println("Definition:", d)
			}
		} else {
			// order example: brwrr
			// definitions: b, br, w, wr
			// Parse the orders, orders are just strings without any seperators
			order := line
			// Add cache map before check function
			cache := make(map[string]bool)

			var check func(int, string) bool
			check = func(k int, s string) bool {
				// Check cache first
				key := fmt.Sprintf("%d:%s", k, s)
				if result, exists := cache[key]; exists {
					return result
				}

				// Base case: if we've gone through all characters
				if k >= len(order) {
					result := definitions[s]
					cache[key] = result
					return result
				}

				// Check current string
				if _, ok := definitions[s]; ok {
					result := check(k+1, string(order[k])) || check(k+1, s+string(order[k]))
					cache[key] = result
					return result
				}

				// Try appending next character to current string
				result := check(k+1, s+string(order[k]))
				cache[key] = result
				return result
			}
			if check(0, string(order[0])) {
				fmt.Println("[ ] Possible for ", order)
				possible++
			} else {
				fmt.Println("[!] Impossible for ", order)
			}

		}
	}
	fmt.Println("A) Orders possible:", possible)
}

func PartB() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	state := DEFINITIONS
	definitions := map[string]bool{}
	count := uint64(0)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if line == "" {
			state = ORDERS
			continue
		}
		if state == DEFINITIONS {
			// Parse the definitions
			def := strings.Split(line, ", ")
			for _, d := range def {
				definitions[d] = true
				fmt.Println("Definition:", d)
			}
		} else {
			order := line
			cache := map[string]int{}

			var check func(int, string) int
			check = func(k int, s string) int {
				// Check cache first
				key := fmt.Sprintf("%d:%s", k, s)
				if result, exists := cache[key]; exists {
					return result
				}

				// Base case: if we've gone through all characters
				if k >= len(order) {
					if definitions[s] {
						cache[key]++
						return 1
					}
					cache[key] = 0
					return 0
				}

				// Check current string
				total := 0
				if _, ok := definitions[s]; ok {
					total += check(k+1, string(order[k]))
					total += check(k+1, s+string(order[k]))
				} else {
					total += check(k+1, s+string(order[k]))
				}

				cache[key] = total
				return total
			}
			res := check(0, "")
			fmt.Println("Result:", res, "for", order)
			count += uint64(res)
		}
	}
	fmt.Println("B) Orders possible:", count)
}

func main() {
	PartA()
	PartB()
}
