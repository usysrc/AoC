package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id         int
	items      []int
	throwsTo   []Monkey
	operation  func(old int) int
	test       func(item int) bool
	throwTrue  int
	throwFalse int
	inspects   int
}

func parseThrow(line string) int {
	// get the right side starting from the colon
	split := strings.Split(line, ":")
	if len(split) < 2 {
		panic("split length too short")
	}

	throw := split[1]
	var monkeyID int
	_, err := fmt.Sscanf(throw, " throw to monkey %d", &monkeyID)
	if err != nil {
		fmt.Println(throw, line)
		panic(err)
	}

	return monkeyID
}

// returns a test func from a string in the form of "  Test: divisible by 19"
func parseTest(line string) func(item int) bool {
	// get the right side starting from the colon
	split := strings.Split(line, ":")
	if len(split) < 2 {
		panic("split length too short")
	}

	// parse out the divisor
	test := split[1]
	var divisor int
	_, err := fmt.Sscanf(test, " divisible by %d", &divisor)
	if err != nil {
		fmt.Println(test, line)
		panic(err)
	}

	// return the test func
	return func(item int) bool {
		return item%divisor == 0
	}
}

// returns a func for an operation in the form of "  Operation: new = old * 19"
func parseOperation(line string) func(old int) int {
	// get the right side starting from the colon
	split := strings.Split(line, ":")
	if len(split) < 2 {
		panic("split length too short")
	}

	// parse the operation
	operation := split[1]
	var operator string
	var operandString string
	_, err := fmt.Sscanf(operation, " new = old %s %s", &operator, &operandString)
	if err != nil {
		fmt.Println(operation, line)
		panic(err)
	}

	// generate functions for with two variables
	if operandString == "old" {
		if operator == "*" {
			return func(old int) int {
				return old * old
			}
		} else if operator == "+" {
			return func(old int) int {
				return old + old
			}
		}
	}

	// generate functions with one variable and an operand
	operand, err := strconv.Atoi(operandString)
	if err != nil {
		panic(err)
	}
	if operator == "*" {
		return func(old int) int {
			return old * operand
		}
	} else if operator == "+" {
		return func(old int) int {
			return old + operand
		}
	} else {
		panic("operation parsing error, operator not found")
	}
}

// return the items from a line in the form of "  Starting items: 79, 98"
func parseItems(line string) []int {
	split := strings.Split(line, ":")
	if len(split) < 2 {
		panic("split length too short")
	}
	itemsStrings := strings.Split(strings.ReplaceAll(split[1], " ", ""), ",")
	items := []int{}
	for i := range itemsStrings {
		item, err := strconv.Atoi(itemsStrings[i])
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}

// simulate a round of play
func round(monkeys map[int]*Monkey) {
	for i := 0; i < len(monkeys); i++ {
		monkey := monkeys[i]
		for _, item := range monkey.items {
			monkey.inspects++
			item = monkey.operation(item)
			item = item / 3
			if monkey.test(item) {
				targetMonkey := monkeys[monkey.throwTrue]
				targetMonkey.items = append(targetMonkey.items, item)
			} else {
				targetMonkey := monkeys[monkey.throwFalse]
				targetMonkey.items = append(targetMonkey.items, item)
			}
		}
		monkey.items = []int{}
	}
}

// return the number of a string in the form of "Monkey 1", "Monkey 2" etc
func parseID(line string) int {
	var id int
	_, err := fmt.Sscanf(line, "Monkey %d", &id)
	if err != nil {
		fmt.Println(line)
		panic(err)
	}
	return id
}

func readFile() int {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	monkeys := map[int]*Monkey{}
	var currentMonkey *Monkey
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 4 && line[:4] == "    " {
			if line[:11] == "    If true" {
				currentMonkey.throwTrue = parseThrow(line)
			} else if line[:12] == "    If false" {
				currentMonkey.throwFalse = parseThrow(line)
			}
		} else if len(line) >= 2 && line[:2] == "  " {
			if line[:16] == "  Starting items" {
				currentMonkey.items = parseItems(line)
			} else if line[:11] == "  Operation" {
				currentMonkey.operation = parseOperation(line)
			} else if line[:6] == "  Test" {
				currentMonkey.test = parseTest(line)
			}
		} else if line != "" {
			currentMonkey = &Monkey{
				id: parseID(line),
			}
			monkeys[currentMonkey.id] = currentMonkey
		}
	}

	for i := 0; i < 20; i++ {
		round(monkeys)
	}

	max := []int{}
	for k, monkey := range monkeys {
		fmt.Println(k, "inspects ", monkey.inspects)
		max = append(max, monkey.inspects)
	}

	sort.Slice(max, func(i, j int) bool {
		return max[i] < max[j]
	})

	return max[len(max)-1] * max[len(max)-2]
}

func main() {
	result := readFile()
	fmt.Println("Result is:", result)
}
