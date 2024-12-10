package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pos struct {
	x, y int
}

const (
	TRAILHEAD = 0
	TRAILEND  = 9
)

// Part A
func countTrailEnds(m map[Pos]int, pos Pos, score int, visited map[Pos]bool, w, h int) int {
	num, ok := m[pos]
	if !ok {
		panic("invalid position")
	}
	if num == TRAILEND {
		if !visited[pos] {
			score++
		}
		visited[pos] = true
		return score
	}
	visited[pos] = true
	defer func() {
		visited[pos] = false
	}()
	neighbors := []Pos{
		{pos.x - 1, pos.y},
		{pos.x + 1, pos.y},
		{pos.x, pos.y - 1},
		{pos.x, pos.y + 1},
	}
	for _, n := range neighbors {
		if n.x < 0 || n.x >= w || n.y < 0 || n.y >= h {
			continue
		}
		if m[n] == num+1 {
			score = countTrailEnds(m, n, score, visited, w, h)
		}
	}
	return score
}

func countTrails(m map[Pos]int, pos Pos, score int, visited map[Pos]bool, w, h int) int {
	num, ok := m[pos]
	if !ok {
		panic("invalid position")
	}
	if num == TRAILEND {
		score++
		return score
	}

	neighbors := []Pos{
		{pos.x - 1, pos.y},
		{pos.x + 1, pos.y},
		{pos.x, pos.y - 1},
		{pos.x, pos.y + 1},
	}
	for _, n := range neighbors {
		if n.x < 0 || n.x >= w || n.y < 0 || n.y >= h {
			continue
		}
		if m[n] == num+1 {
			score = countTrails(m, n, score, visited, w, h)
		}
	}
	return score
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file

	m := map[Pos]int{}
	// array of positions
	trailheads := []Pos{}

	scanner := bufio.NewScanner(file)
	j := 0
	w, h := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		// Do something with the line
		for i, r := range line {
			c := string(r)
			k, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			m[Pos{i, j}] = k
			if i >= w {
				w = i + 1
			}
			if k == TRAILHEAD {
				trailheads = append(trailheads, Pos{i, j})
			}
		}
		j++
		h = j
	}

	// print map
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			print(m[Pos{i, j}])
		}
		println()
	}

	// find all possible endpoints from trailheads
	sumA := 0
	for _, pos := range trailheads {
		score := countTrailEnds(m, pos, 0, map[Pos]bool{}, w, h)
		sumA += score
	}

	fmt.Println("Score Part A:", sumA)

	// find all possible trails from trailheads
	sumB := 0
	for _, pos := range trailheads {
		score := countTrails(m, pos, 0, map[Pos]bool{}, w, h)
		sumB += score
	}

	fmt.Println("Score Part A:", sumB)

}
