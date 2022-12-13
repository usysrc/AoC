// doesn't work
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x     int
	y     int
	value int
}

// Grid type, might not be fastest but at least I understand how it works and it can be of arbitrary size
type Grid struct {
	field map[string]int
	start Point
	end   Point
}

func NewGrid() Grid {
	return Grid{
		field: map[string]int{},
	}
}

func (g Grid) get(i, j int) (int, bool) {
	val, ok := g.field[strconv.Itoa(i)+","+strconv.Itoa(j)]
	return val, ok
}

func (g Grid) set(i, j int, v int) {
	g.field[strconv.Itoa(i)+","+strconv.Itoa(j)] = v
}

func (g Grid) del(i, j int) {
	delete(g.field, strconv.Itoa(i)+","+strconv.Itoa(j))
}

// returns all left, right, up, down neighbours (if existing, they don't exist on borders)
func (g Grid) getNeighbours(i, j int) []Point {
	neighbours := []Point{}
	if val, ok := g.get(i-1, j); ok {
		neighbours = append(neighbours, Point{x: i - 1, y: j, value: val})
	}
	if val, ok := g.get(i+1, j); ok {
		neighbours = append(neighbours, Point{x: i + 1, y: j, value: val})
	}
	if val, ok := g.get(i, j-1); ok {
		neighbours = append(neighbours, Point{x: i, y: j - 1, value: val})
	}
	if val, ok := g.get(i, j+1); ok {
		neighbours = append(neighbours, Point{x: i, y: j + 1, value: val})
	}
	return neighbours
}

type DistPoint struct {
	p    Point
	dist int
}

// get the distance
func getDistance(g Grid) int {

	visited := NewGrid()
	visited.set(g.end.x, g.end.y, 1)
	queue := []DistPoint{}
	queue = append(queue, DistPoint{p: g.end, dist: 0})

	for len(queue) > 0 {
		currentDistPoint := queue[0]
		queue = queue[1:]
		currentPoint := currentDistPoint.p

		for _, p := range g.getNeighbours(currentPoint.x, currentPoint.y) {
			if p.x == g.start.x && p.y == g.start.y {
				fmt.Println("found the start")
				return currentDistPoint.dist + 1
			}
			if _, ok := visited.get(p.x, p.y); !ok {
				if p.value >= currentPoint.value-1 {
					fmt.Println(p.x, p.y, currentPoint.value, p.value, currentDistPoint.dist+1, string(rune(currentPoint.value)), string(rune(p.value)))
					visited.set(p.x, p.y, 1)
					queue = append(queue, DistPoint{p: p, dist: currentDistPoint.dist + 1})
				}
			}
		}
	}
	return -1
}

// read the file and create a grid of weights
func readFile() Grid {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	g := NewGrid()
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			if v == 'S' {
				g.start = Point{x: i, y: j, value: int('a')}
				g.set(i, j, int('a'))
			} else if v == 'E' {
				g.end = Point{x: i, y: j, value: int('z')}
				g.set(i, j, int('z'))
			} else {
				// fmt.Println(int(v))
				g.set(i, j, int(v))
			}
		}
		j++
	}
	return g
}

//

func main() {
	grid := readFile()
	distance := getDistance(grid)
	fmt.Println("result is:", distance)
}
