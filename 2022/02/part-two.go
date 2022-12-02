package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Shape int

// Shape enums
const (
	Rock     Shape = iota
	Paper    Shape = iota
	Scissors Shape = iota
)

// convert A,B,C to the shape
func convert(s string) Shape {
	if s == "A" {
		return Rock
	}
	if s == "B" {
		return Paper
	}
	if s == "C" {
		return Scissors
	}
	return Rock
}

// return points for the shape
func points(shape Shape) int {
	if shape == Rock {
		return 1
	}
	if shape == Paper {
		return 2
	}
	if shape == Scissors {
		return 3
	}
	return 0
}

// returns a beats b for values of a and b
func beats(a, b, l, r Shape) bool {
	if r == a && l == b {
		return true
	}
	return false
}

// calculate the points for the outcome (win, draw or loss)
func win(l, r Shape) int {
	if l == r {
		return 3
	}
	if beats(Rock, Scissors, l, r) || beats(Scissors, Paper, l, r) || beats(Paper, Rock, l, r) {
		return 6
	}
	return 0
}

func chooseShape(neededWin string, shapeOpponent Shape) Shape {

	// lose
	if neededWin == "X" {
		if shapeOpponent == Rock {
			return Scissors
		}
		if shapeOpponent == Paper {
			return Rock
		}
		if shapeOpponent == Scissors {
			return Paper
		}
	}

	// draw
	if neededWin == "Y" {
		return shapeOpponent
	}

	// win
	if neededWin == "Z" {
		if shapeOpponent == Rock {
			return Paper
		}
		if shapeOpponent == Paper {
			return Scissors
		}
		if shapeOpponent == Scissors {
			return Rock
		}
	}

	return Rock
}

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(inputFile)
	reader.Comma = ' '
	csvLines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	score := 0
	for _, line := range csvLines {
		left := convert(line[0])
		neededWin := line[1]
		right := chooseShape(neededWin, left)
		score += points(right)
		score += win(left, right)
	}
	return score
}

func main() {
	finalScore := readFile()
	fmt.Println(finalScore)
}
