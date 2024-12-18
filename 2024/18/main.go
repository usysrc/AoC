package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	memory := map[Pos]string{}
	fallingBytes := []Pos{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, b := 0, 0
		fmt.Sscanf(line, "%d,%d", &a, &b)
		fallingBytes = append(fallingBytes, Pos{a, b})
		i++
	}
	memory = map[Pos]string{}
	for i, pos := range fallingBytes {
		if i < 12 {
			memory[pos] = "#"
		}
	}
	printShortestPath(memory)

	memory = map[Pos]string{}
	for _, pos := range fallingBytes {
		memory[pos] = "#"
		if !hasPath(memory) {
			fmt.Print(pos)
			break
		}
	}
}

// checks if there is a path from start to exit
func hasPath(memory map[Pos]string) bool {
	w, h := 70, 70
	start := Pos{0, 0}
	// w, h := 6, 6
	exit := Pos{w, h}

	visited := map[Pos]bool{}
	queue := []Pos{start}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if visited[pos] {
			continue
		}
		visited[pos] = true
		if pos == exit {
			return true
		}
		for _, dir := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if pos.x+dir.x < 0 || pos.x+dir.x > w || pos.y+dir.y < 0 || pos.y+dir.y > h {
				continue
			}
			next := Pos{pos.x + dir.x, pos.y + dir.y}
			if memory[next] == "#" {
				continue
			}
			queue = append(queue, next)
		}
	}
	return false
}

func printShortestPath(memory map[Pos]string) {
	// w, h := 70, 70
	w, h := 6, 6
	exit := Pos{w, h}
	start := Pos{0, 0}

	dijkstraMap := map[Pos]int{exit: 0}
	visited := map[Pos]bool{}
	queue := []Pos{exit}
	dijkstraMap[exit] = 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if visited[pos] {
			continue
		}
		visited[pos] = true
		for _, dir := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if pos.x+dir.x < 0 || pos.x+dir.x > w || pos.y+dir.y < 0 || pos.y+dir.y > h {
				continue
			}
			next := Pos{pos.x + dir.x, pos.y + dir.y}
			if memory[next] == "#" {
				continue
			}
			if _, ok := dijkstraMap[next]; !ok {
				dijkstraMap[next] = dijkstraMap[pos] + 1
			}
			queue = append(queue, next)
		}
	}
	shortestPath := []Pos{start}
	pos := start
	fmt.Println(dijkstraMap[start])
	for pos != exit {
		for _, dir := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if pos.x+dir.x < 0 || pos.x+dir.x > w || pos.y+dir.y < 0 || pos.y+dir.y > h {
				continue
			}
			next := Pos{pos.x + dir.x, pos.y + dir.y}
			if memory[next] == "#" {
				continue
			}
			if dijkstraMap[next] == dijkstraMap[pos]-1 {
				shortestPath = append(shortestPath, next)
				pos = next
				break
			}
		}
	}

	fmt.Println("Steps:", len(shortestPath)-1)
	// print the memory and the shortest path
	// for y := 0; y < h+1; y++ {
	// 	for x := 0; x < w+1; x++ {
	// 		if memory[Pos{x, y}] == "#" {
	// 			fmt.Print("#")
	// 		} else {
	// 			found := false
	// 			for _, pos := range shortestPath {
	// 				if pos.x == x && pos.y == y {
	// 					fmt.Print("O")
	// 					found = true
	// 					break
	// 				}
	// 			}
	// 			if !found {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
