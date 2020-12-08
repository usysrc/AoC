package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func splitInto(str string, sep string) (string, string) {
	s := strings.Split(str, sep)
	return s[0], s[1]
}

type instruction struct {
	operation string
	argument int
}


func main() {
	program := make([]instruction, 0)

	for _, ln := range readFile() {
		left, right := splitInto(ln, " ")
		arg, err := strconv.Atoi(right)
		if err != nil {
			log.Fatal(err)
		}

		program = append(program, instruction{operation:left, argument:arg})
	}
	instructionTrace := make(map[int]bool)
	p := 0
	accumulator := 0

	for {
		if instructionTrace[p] {
			fmt.Println(accumulator)
			break
		}
		instructionTrace[p] = true
		instruction := program[p]
		if instruction.operation == "acc"{
			accumulator += instruction.argument
			p++
		}
		if instruction.operation == "jmp"{
			p += instruction.argument
		}
		if instruction.operation == "nop"{
			p++
		}

	}
}