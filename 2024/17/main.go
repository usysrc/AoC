package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	A                  int
	B                  int
	C                  int
	Program            []int
	InstructionPointer int
}

func main() {
	filename := "input"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := Machine{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// if this line is a register line
		// then parse the register value
		// else parse the program
		if len(line) == 0 {
			continue
		}
		if line[:8] == "Register" {
			var registerValue int
			var register string
			fmt.Println(line)
			fmt.Sscanf(line, "Register %s %d", &register, &registerValue)
			fmt.Println(register, registerValue)
			switch register {
			case "A:":
				m.A = registerValue
			case "B:":
				m.B = registerValue
			case "C:":
				m.C = registerValue
			}
		} else if line[:7] == "Program" {
			program := []int{}
			programString := line[9:]
			numbers := strings.Split(programString, ",")
			for _, number := range numbers {
				num, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				program = append(program, num)
			}
			m.Program = program
		}
	}

	fmt.Println("Machine before starting:", m)

	machine := Machine{
		A:                  0,
		B:                  m.B,
		C:                  m.C,
		Program:            m.Program,
		InstructionPointer: 0,
	}
	a := 0
	for i := len(m.Program) - 1; i >= 0; i-- {
		a <<= 3
		machine.A = a
		machine.B = m.B
		machine.C = m.C
		machine.Program = m.Program
		machine.InstructionPointer = 0

		for !slices.Equal(run(machine), m.Program[i:]) {
			a++
			machine.A = a
			machine.B = m.B
			machine.C = m.C
			machine.Program = m.Program
			machine.InstructionPointer = 0
			// fmt.Println(machine, run(machine))
		}
	}

	fmt.Println("Result:", a)
}

func run(m Machine) []int {
	output := []string{}
	for m.InstructionPointer < len(m.Program) {
		opcode := m.Program[m.InstructionPointer]
		operand := m.Program[m.InstructionPointer+1]
		literalOperand := operand
		comboOperand := operand
		if operand == 4 {
			comboOperand = m.A
		} else if operand == 5 {
			comboOperand = m.B
		} else if operand == 6 {
			comboOperand = m.C
		} else if operand == 7 {
			// nothing to do here
			panic("Invalid operand")
		}

		switch opcode {
		case 0:
			// adv opcode, performs division, register A = register A / (2^comboOperand)
			m.A >>= comboOperand
		case 1:
			// bxl opcode, calculate bitwise XOR of register B, register B = register B XOR operand
			m.B ^= literalOperand
		case 2:
			// bst opcode, register B = comboOperand%8
			m.B = comboOperand % 8
		case 3:
			// jnz opcode, jump to literalOperand if register A is not zero
			if m.A != 0 {
				m.InstructionPointer = literalOperand - 2
			}
		case 4:
			// bxc, bitwise XOR of register B and C, register B = register B XOR register C, operand is ignored
			m.B ^= m.C
		case 5:
			// out, calculate comboOperand%8 and print it
			output = append(output, strconv.Itoa(comboOperand%8))
		case 6:
			// bdv, like adv but for register B
			m.B = m.A >> comboOperand
		case 7:
			// cdv, like adv but for register C
			m.C = m.A >> comboOperand
		}
		m.InstructionPointer += 2
	}
	// fmt.Println(m)
	// fmt.Println(strings.Join(output, ","))
	// fmt.Println(strings.Join(output, ""))

	out := []int{}
	for _, o := range output {
		num, err := strconv.Atoi(o)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}
	return out
}
