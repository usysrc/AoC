package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	X, Y int64
}

type Machine struct {
	ButtonA Pos
	ButtonB Pos
	Prize   Pos
}

func minimizeCost(m Machine, offset int64) int64 {
	prize := Pos{m.Prize.X + offset, m.Prize.Y + offset}
	det := m.ButtonA.X*m.ButtonB.Y - m.ButtonA.Y*m.ButtonB.X
	a := (prize.X*m.ButtonB.Y - prize.Y*m.ButtonB.X) / det
	b := (m.ButtonA.X*prize.Y - m.ButtonA.Y*prize.X) / det
	if m.ButtonA.X*a+m.ButtonB.X*b == prize.X && m.ButtonA.Y*a+m.ButtonB.Y*b == prize.Y {
		return a*3 + b
	}
	return -1
}

func main() {
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := []Machine{}
	machine := Machine{}
	for scanner.Scan() {
		line := scanner.Text()
		// Do something with the line

		// parse Button A: X+69, Y+23
		fmt.Sscanf(line, "Button A: X+%d, Y+%d", &machine.ButtonA.X, &machine.ButtonA.Y)

		// parse Button B: X+69, Y+23
		fmt.Sscanf(line, "Button B: X+%d, Y+%d", &machine.ButtonB.X, &machine.ButtonB.Y)
		// parse Prize: X+69, Y+23
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &machine.Prize.X, &machine.Prize.Y)

		// empty line indicates end of machine
		if line == "" {
			machines = append(machines, machine)
			machine = Machine{}
		}
	}
	machines = append(machines, machine)

	// find the cheapest number of buttonA presses and buttonB presses to get to the prize position
	sumA := int64(0)
	sumB := int64(0)
	for _, machine := range machines {
		costA := minimizeCost(machine, 0)
		if costA > 0 {
			fmt.Println(costA)
			sumA += costA
		}

		costB := minimizeCost(machine, 10000000000000)
		if costB > 0 {
			fmt.Println(costB)
			sumB += costB
		}
	}
	fmt.Println("Sum of all costs, Part A: ", sumA)
	fmt.Println("Sum of all costs, Part B: ", sumB)
}
