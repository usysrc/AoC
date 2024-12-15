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
	printMap(height, width, warehouse, robot)
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
		if isCrate(warehouse[Pos{robot.x + dx, robot.y + dy}].name) {

			// move the crate and all the crates behind it
			var moveBox func(Pos) bool
			moveBox = func(p Pos) bool {
				nextPosA := Pos{p.x + dx, p.y + dy}
				nextPosB := Pos{p.x + dx + 1, p.y + dy}

				if warehouse[nextPosA].name == WALL || warehouse[nextPosB].name == WALL {
					return false
				}
				if isCrate(warehouse[nextPosA].name) || isCrate(warehouse[nextPosB].name) {
					blockingBoxA, foundA := boxes[nextPosA]
					blockingBoxB, foundB := boxes[nextPosB]

					if foundA && !moveBox(blockingBoxA.Pos) {
						return false
					}
					if foundB && !moveBox(blockingBoxB.Pos) {
						return false
					}

					// re-check if the next position is a wall
					if warehouse[nextPosA].name == WALL || warehouse[nextPosB].name == WALL {
						return false
					}
					if isCrate(warehouse[nextPosA].name) || isCrate(warehouse[nextPosB].name) {
						return false
					}
				}
				// move the box now that we know it is safe
				currentTileA := warehouse[p]
				currentTileB := warehouse[Pos{p.x + 1, p.y}]
				warehouse[p] = Tile{p, SPACE}
				warehouse[Pos{p.x + 1, p.y}] = Tile{Pos{p.x + 1, p.y}, SPACE}

				warehouse[nextPosA] = Tile{nextPosA, currentTileA.name}
				warehouse[nextPosB] = Tile{nextPosB, currentTileB.name}

				delete(boxes, p)
				delete(boxes, Pos{p.x + 1, p.y})
				box := boxes[nextPosA]
				boxes[nextPosA] = box
				boxes[nextPosB] = box
				return true
			}
			if dy != 0 {
				box := boxes[Pos{robot.x + dx, robot.y + dy}]
				moveBox(box.Pos)
			} else {
				// move the crate and all the crates behind it
				var move func(Pos)
				move = func(p Pos) {
					nextPos := Pos{p.x + dx, p.y + dy}
					if warehouse[nextPos].name == WALL {
						return
					}

					// check if there is a crate at the next position
					if isCrate(warehouse[p].name) {
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

		}
		if warehouse[Pos{robot.x + dx, robot.y + dy}].name == SPACE {
			robot.x += dx
			robot.y += dy
		}
		fmt.Println(n)
		printMap(height, width, warehouse, robot)
	}

}

func printMap(height int, width int, warehouse map[Pos]Tile, robot Pos) {
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if robot.x == i && robot.y == j {
				fmt.Print("@")
			} else {
				fmt.Print(warehouse[Pos{i, j}].name)
			}
		}
		fmt.Println()
	}
}

func main() {
	partA()
	partB()
}
