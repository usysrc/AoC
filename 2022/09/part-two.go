package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func moveKnot(head, tail *Point) {
	if math.Abs(float64(head.X-tail.X)) > 1 || math.Abs(float64(head.Y-tail.Y)) > 1 {
		if head.X == tail.X {
			if head.Y < tail.Y {
				tail.Y--
			} else if head.Y > tail.Y {
				tail.Y++
			}
		} else if head.Y == tail.Y {
			if head.X < tail.X {
				tail.X--
			} else if head.X > tail.X {
				tail.X++
			}
		} else {
			// diagonally
			if head.X > tail.X {
				tail.X++
			} else if head.X < tail.X {
				tail.X--
			}
			if head.Y > tail.Y {
				tail.Y++
			} else if head.Y < tail.Y {
				tail.Y--
			}
		}
	}
}

func readFile() int {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	head := Point{
		X: 0,
		Y: 0,
	}
	var tails [9]Point
	for i := 0; i < 9; i++ {
		tails[i] = Point{
			X: 0,
			Y: 0,
		}
	}
	visited := map[string]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		var direction string
		var steps int
		_, err := fmt.Sscanf(line, "%s %d", &direction, &steps)
		if err != nil {
			panic(err)
		}
		for k := 0; k < steps; k++ {
			if direction == "R" {
				head.X += 1
			}
			if direction == "L" {
				head.X -= 1
			}
			if direction == "U" {
				head.Y -= 1
			}
			if direction == "D" {
				head.Y += 1
			}
			moveKnot(&head, &tails[0])
			for i := 1; i < 9; i++ {
				moveKnot(&tails[i-1], &tails[i])
			}

			visited[strconv.Itoa(tails[8].X)+","+strconv.Itoa(tails[8].Y)] = true
		}

	}
	count := 0
	for _, a := range visited {
		if a {
			count++
		}
	}
	return count
}

func main() {
	steps := readFile()
	fmt.Println("result is:", steps)
}
