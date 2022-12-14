package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Stone int = iota
	Sand  int = iota
)

type Point struct {
	x     int
	y     int
	value int
}

// Grid type, might not be fastest but at least I understand how it works and it can be of arbitrary size
type Grid struct {
	field        map[string]int
	lowestY      int
	freezeLowest bool
}

func NewGrid() *Grid {
	return &Grid{
		field:        map[string]int{},
		lowestY:      0,
		freezeLowest: false,
	}
}

func (g *Grid) get(i, j int) (int, bool) {
	val, ok := g.field[strconv.Itoa(i)+","+strconv.Itoa(j)]
	return val, ok
}

func (g *Grid) set(i, j int, v int) {
	if !g.freezeLowest && j > g.lowestY {
		g.lowestY = j
	}
	g.field[strconv.Itoa(i)+","+strconv.Itoa(j)] = v
}

func swap(a, b Point) (Point, Point) {
	return b, a
}

func (g *Grid) drawLine(a, b Point) {
	if a.x == b.x {
		if a.y > b.y {
			a, b = swap(a, b)
		}
		for y := a.y; y <= b.y; y++ {
			g.set(a.x, y, Stone)
		}
	}
	if a.y == b.y {
		if a.x > b.x {
			a, b = swap(a, b)
		}
		for x := a.x; x <= b.x; x++ {
			g.set(x, a.y, Stone)
		}
	}
}

func readFile() *Grid {
	inputFile, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	g := NewGrid()
	for scanner.Scan() {
		line := scanner.Text()
		points := []Point{}
		rawPoints := strings.Split(line, "->")
		for _, p := range rawPoints {
			pointString := strings.TrimSpace(p)
			coordinates := strings.Split(pointString, ",")
			x, errX := strconv.Atoi(coordinates[0])
			if errX != nil {
				panic(errX)
			}
			y, errY := strconv.Atoi(coordinates[1])
			if errY != nil {
				panic(errY)
			}
			points = append(points, Point{
				x: x,
				y: y,
			})
		}
		for i := 0; i < len(points)-1; i++ {
			g.drawLine(points[i], points[i+1])
		}
	}
	return g
}

// used for debugging
func drawMap(g *Grid) {
	fmt.Println("lowest y:", g.lowestY)
	for j := 0; j < 15; j++ {
		for i := 485; i < 515; i++ {
			val, ok := g.get(i, j)
			if ok {
				if val == Stone {
					fmt.Print("#")
				}
				if val == Sand {
					fmt.Print("o")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func simulateSand(g *Grid) int {
	g.freezeLowest = true
	source := Point{
		x: 500,
		y: 0,
	}
	particleCount := 0
	for true {
		particleCount++
		sandParticle := Point{
			x: source.x,
			y: source.y,
		}
		for true {
			_, okDown := g.get(sandParticle.x, sandParticle.y+1)
			if !okDown {
				sandParticle.y++
				if sandParticle.y == g.lowestY+1 {
					break
				}
				continue
			}
			_, okDownAndLeft := g.get(sandParticle.x-1, sandParticle.y+1)
			if !okDownAndLeft {
				sandParticle.x--
				sandParticle.y++
				if sandParticle.y == g.lowestY+1 {
					break
				}
				continue
			}
			_, okDownAndRight := g.get(sandParticle.x+1, sandParticle.y+1)
			if !okDownAndRight {
				sandParticle.x++
				sandParticle.y++
				if sandParticle.y == g.lowestY+1 {
					break
				}
				continue
			}
			break
		}
		g.set(sandParticle.x, sandParticle.y, Sand)
		if sandParticle.x == source.x && sandParticle.y == source.y {
			break
		}
	}
	drawMap(g)

	return particleCount
}

func main() {
	grid := readFile()
	units := simulateSand(grid)
	fmt.Println("How many units: :", units)
}
