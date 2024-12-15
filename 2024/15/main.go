package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x, y int
}

type Tile struct {
	Pos
	name string
}

const (
	MAP = iota
	MOVES
)

const (
	WALL       = "#"
	SPACE      = "."
	CRATE      = "O"
	CRATELEFT  = "["
	CRATERIGHT = "]"
)

func isCrate(s string) bool {
	return s == CRATE || s == CRATELEFT || s == CRATERIGHT
}

func partA() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := MAP
	warehouse := map[Pos]Tile{}
	moves := []string{}
	robot := Pos{0, 0}
	width, height := 0, 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if state == MAP {
			for i, c := range line {
				if string(c) == "@" {
					robot = Pos{i, j}
					warehouse[Pos{i, j}] = Tile{Pos{i, j}, SPACE}
				} else {
					warehouse[Pos{i, j}] = Tile{Pos{i, j}, string(c)}
				}
				width = i + 1
			}
			j++
			height = j
		}

		if state == MOVES {
			for _, c := range line {
				moves = append(moves, string(c))
			}
		}
		if line == "" {
			state = MOVES
		}
	}

	// s can be ^,v,<,>
	for _, s := range moves {
		dx, dy := 0, 0
		switch s {
		case "^":
			dy--
		case "v":
			dy++
		case "<":
			dx--
		case ">":
			dx++
		}
		if warehouse[Pos{robot.x + dx, robot.y + dy}].name == WALL {
			continue
		}
		// check if there is a crate at the next position
		if warehouse[Pos{robot.x + dx, robot.y + dy}].name == CRATE {
			// move the crate and all the crates behind it
			var move func(Pos)
			move = func(p Pos) {
				nextPos := Pos{p.x + dx, p.y + dy}
				if warehouse[nextPos].name == WALL {
					return
				}

				// check if there is a crate at the next position
				if warehouse[p].name == CRATE {
					move(nextPos)
				}
				// store the current tile and move the crate
				currentTile := warehouse[p]
				nextTile := warehouse[nextPos]
				if nextTile.name != SPACE {
					return
				}
				warehouse[p] = Tile{p, SPACE}
				warehouse[nextPos] = Tile{nextPos, currentTile.name}
			}
			move(Pos{robot.x + dx, robot.y + dy})
		}
		if warehouse[Pos{robot.x + dx, robot.y + dy}].name == SPACE {
			robot.x += dx
			robot.y += dy
		}

	}
	fmt.Println(width, height)

	sumA := 0
	for _, v := range warehouse {
		if v.name == CRATE {
			sumA += v.y*100 + v.x
		}
	}
	fmt.Println("Sum of Part A:", sumA)

}

func partB() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := MAP
	warehouse := map[Pos]Tile{}
	moves := []string{}
	robot := Pos{0, 0}
	boxes := map[Pos]Tile{}
	width, height := 0, 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if state == MAP {
			for k, c := range line {
				i := k * 2
				if string(c) == "@" {
					robot = Pos{i, j}
					warehouse[Pos{i, j}] = Tile{Pos{i, j}, SPACE}
					warehouse[Pos{i + 1, j}] = Tile{Pos{i + 1, j}, SPACE}
				} else {
					if string(c) == CRATE {
						box := Tile{Pos{i, j}, CRATE}
						boxes[Pos{i, j}] = box
						boxes[Pos{i + 1, j}] = box
						warehouse[Pos{i, j}] = Tile{Pos{i, j}, CRATELEFT}
						warehouse[Pos{i + 1, j}] = Tile{Pos{i + 1, j}, CRATERIGHT}
					}
					if string(c) == SPACE {
						warehouse[Pos{i, j}] = Tile{Pos{i, j}, SPACE}
						warehouse[Pos{i + 1, j}] = Tile{Pos{i + 1, j}, SPACE}
					}
					if string(c) == WALL {
						warehouse[Pos{i, j}] = Tile{Pos{i, j}, WALL}
						warehouse[Pos{i + 1, j}] = Tile{Pos{i + 1, j}, WALL}
					}
				}
				width = (i + 2)
			}
			j++
			height = j
		}

		if state == MOVES {
			for _, c := range line {
				moves = append(moves, string(c))
			}
		}
		if line == "" {
			state = MOVES
		}
	}

	// s can be ^,v,<,>
	for n, s := range moves {
		dx, dy := 0, 0
		switch s {
		case "^":
			dy--
		case "v":
			dy++
		case "<":
			dx--
		case ">":
			dx++
		}
		if warehouse[Pos{robot.x + dx, robot.y + dy}].name == WALL {
			continue
		}
		// check if there is a crate at the next position
		_, found := boxes[Pos{robot.x + dx, robot.y + dy}]
		if found {

			// move the crate and all the crates behind it
			moved := map[Pos]bool{}
			moves := []func(){}
			var moveBox func(Pos) bool
			moveBox = func(p Pos) bool {
				if moved[p] {
					return true
				}
				crateA, foundA := boxes[Pos{p.x + dx, p.y + dy}]
				crateB, foundB := boxes[Pos{p.x + dx + 1, p.y + dy}]
				if warehouse[Pos{p.x + dx, p.y + dy}].name == WALL {
					return false
				}
				if warehouse[Pos{p.x + dx + 1, p.y + dy}].name == WALL {
					return false
				}
				if foundA && crateA.Pos != p {

					if !moveBox(crateA.Pos) {
						return false
					}
				}
				if foundB && crateB.Pos != p {

					if !moveBox(crateB.Pos) {
						return false
					}
				}
				if warehouse[Pos{p.x + dx, p.y + dy}].name == WALL {
					return false
				}
				if warehouse[Pos{p.x + dx + 1, p.y + dy}].name == WALL {
					return false
				}
				// move the box
				moves = append(moves, func() {
					delete(boxes, Pos{p.x, p.y})
					delete(boxes, Pos{p.x + 1, p.y})
					box := Tile{Pos{p.x + dx, p.y + dy}, CRATE}
					boxes[Pos{box.x, box.y}] = box
					boxes[Pos{box.x + 1, box.y}] = box

					// leftTile := warehouse[Pos{p.x, p.y}]
					// rightTile := warehouse[Pos{p.x + 1, p.y}]
					warehouse[Pos{p.x, p.y}] = Tile{Pos{p.x, p.y}, SPACE}
					warehouse[Pos{p.x + 1, p.y}] = Tile{Pos{p.x + 1, p.y}, SPACE}
					warehouse[Pos{box.x, box.y}] = Tile{Pos{box.x, box.y}, CRATELEFT}
					warehouse[Pos{box.x + 1, box.y}] = Tile{Pos{box.x + 1, box.y}, CRATERIGHT}
				})
				moved[p] = true
				return true
			}
			box := boxes[Pos{robot.x + dx, robot.y + dy}]
			if moveBox(box.Pos) {
				for _, m := range moves {
					m()
				}
			}

		}
		_, isABox := boxes[Pos{robot.x + dx, robot.y + dy}]
		if !isABox {
			robot.x += dx
			robot.y += dy
		}
		fmt.Println(n)
		// wait for input
		sumBCrateLeft := 0
		sumBCrateRight := 0
		// printMap(height, width, warehouse, robot, boxes)
		for _, v := range warehouse {
			if v.name == CRATELEFT {
				sumBCrateLeft += v.y*100 + v.x
			}
			if v.name == CRATERIGHT {
				sumBCrateRight += v.y*100 + (v.x - 1)
			}
		}
		if sumBCrateLeft != sumBCrateRight {

			fmt.Scanln()
		}
	}
	printMap(height, width, warehouse, robot, boxes)
	sumB := 0
	visited := map[Pos]bool{}
	for _, v := range boxes {
		if vis, ok := visited[v.Pos]; !ok || !vis {
			sumB += v.y*100 + v.x
		}
		visited[v.Pos] = true
	}
	sumBCrateLeft := 0
	sumBCrateRight := 0
	for _, v := range warehouse {
		if v.name == CRATELEFT {
			sumBCrateLeft += v.y*100 + v.x
		}
		if v.name == CRATERIGHT {
			sumBCrateRight += v.y*100 + (v.x - 1)
		}
	}

	fmt.Println("Sum of Part B:", sumB, sumBCrateLeft, sumBCrateRight)
}

func printMap(height int, width int, warehouse map[Pos]Tile, robot Pos, boxes map[Pos]Tile) {
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if robot.x == i && robot.y == j {
				fmt.Print("@")
			} else if _, ok := boxes[Pos{i, j}]; ok {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	partA()
	partB()
}
