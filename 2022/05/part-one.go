package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseCrates(line string) (crate string) {
	i := 0
	crateLine := ""
	for true {
		if line[i] == '[' {
			i++
			crateLine += string(line[i])
		}
		if i+3 < len(line) && line[i:i+3] == "   " {
			i += 3
			crateLine += " "
		}
		if i+1 >= len(line) {
			break
		}
		i++
	}
	return crateLine
}

func parseMove(line string) (quantity, from, to int) {
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &quantity, &from, &to)
	if err != nil {
		panic(err)
	}
	return quantity, from, to
}

func readFile() string {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	header := true
	scanner := bufio.NewScanner(inputFile)
	var stacks [9][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			header = false
			// invert stacks
			for i, stack := range stacks {
				var newStack []string
				for i := len(stack) - 1; i >= 0; i-- {
					newStack = append(newStack, stack[i])
				}
				stacks[i] = newStack
			}
		}
		if header {
			crateLine := parseCrates(line)
			// add the crates to the stacks
			for i, v := range crateLine {
				if v != ' ' {
					stacks[i] = append(stacks[i], string(v))
				}
			}

		} else if !header && line != "" {
			quantity, from, to := parseMove(line)
			fmt.Println(quantity, from, to)
			for i := 0; i < quantity; i++ {
				fmt.Println(stacks[from-1][len(stacks[from-1])-1])
				stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
				stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
			}
			fmt.Println(stacks)
		}
	}
	str := ""
	for _, v := range stacks {
		str += v[len(v)-1]
	}
	return str
}

func main() {
	topOfStacks := readFile()
	fmt.Println(topOfStacks)
}
