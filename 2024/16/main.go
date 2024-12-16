package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Tile string

type Pos struct {
	X, Y int
}

const (
	WALL  Tile = "#"
	EMPTY Tile = "."
	START Tile = "S"
	END   Tile = "E"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mazeMap := map[Pos]Tile{}
	scanner := bufio.NewScanner(file)

	w, h := 0, 0
	j := 0
	start := Pos{0, 0}
	end := Pos{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			tile := Tile(string(c))
			pos := Pos{i, j}
			mazeMap[pos] = tile
			if tile == START {
				start = pos
			}
			if tile == END {
				end = pos
			}
			w = i + 1
		}
		j++
		h = j
	}
	printMaze(mazeMap, w, h)
	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	left := Pos{-1, 0}
	right := Pos{1, 0}
	up := Pos{0, -1}
	down := Pos{0, 1}

	type State struct {
		pos     Pos
		dir     Pos
		cost    int
		visited map[Pos]bool
	}
	type CostKey struct {
		pos Pos
		dir Pos
	}

	found := false
	queue := []State{{start, right, 0, map[Pos]bool{}}}
	minCostToEnd := math.MaxInt
	costs := map[CostKey]int{}
	pathsToEnd := []State{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// if there is a wall here, skip
		if tile, ok := mazeMap[current.pos]; ok && tile == WALL {
			continue
		}

		key := CostKey{current.pos, current.dir}
		if prevCost, ok := costs[key]; ok {
			if prevCost < current.cost {
				continue
			}
		}
		costs[key] = current.cost

		visited := map[Pos]bool{}
		for k, v := range current.visited {
			visited[k] = v
		}
		visited[current.pos] = true

		if current.pos == end {
			if current.cost < minCostToEnd {
				minCostToEnd = current.cost
				pathsToEnd = []State{current}
			} else if current.cost == minCostToEnd {
				pathsToEnd = append(pathsToEnd, current)
			}
			found = true
			continue
		}

		// straight first
		nextPos := Pos{current.pos.X + current.dir.X, current.pos.Y + current.dir.Y}
		if tile, ok := mazeMap[nextPos]; ok && tile != WALL {
			queue = append(queue, State{nextPos, current.dir, current.cost + 1, visited})
		}

		// try turns
		turns := []Pos{}
		if current.dir == right || current.dir == left {
			turns = []Pos{up, down}
		} else {
			turns = []Pos{right, left}
		}

		for _, turn := range turns {
			nextPos := Pos{current.pos.X + turn.X, current.pos.Y + turn.Y}
			if tile, ok := mazeMap[nextPos]; ok && tile != WALL {
				newVisited := map[Pos]bool{}
				for k, v := range visited {
					newVisited[k] = v
				}
				newVisited[nextPos] = true
				queue = append(queue, State{nextPos, turn, current.cost + 1001, newVisited})
			}
		}
	}

	fmt.Println("Found", found, "-", "Cost:", minCostToEnd)
	vistedMap := map[Pos]bool{}
	for _, path := range pathsToEnd {
		fmt.Println("Path:", path.cost)
		for pos := range path.visited {
			vistedMap[pos] = true
		}
	}
	// print maze with visited tiles
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pos := Pos{i, j}
			if _, ok := vistedMap[pos]; ok {
				print("x")
			} else if tile, ok := mazeMap[pos]; ok {
				print(string(tile))
			} else {
				print(" ")
			}
		}
		println()
	}
	fmt.Println("Visited tiles:", len(vistedMap)+1)
}

func printMaze(mazeMap map[Pos]Tile, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pos := Pos{i, j}
			if tile, ok := mazeMap[pos]; ok {
				print(string(tile))
			} else {
				print(" ")
			}
		}
		println()
	}
}
