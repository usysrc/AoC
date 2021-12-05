package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE = "input.txt"

type GetError struct{}

func (err GetError) Error() string {
	return "Does not exist"
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func CreateLine(x1, y1, x2, y2 string) Line {
	l := Line{}
	x, err := strconv.Atoi(x1)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}
	y, err := strconv.Atoi(y1)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}
	tx, err := strconv.Atoi(x2)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}
	ty, err := strconv.Atoi(y2)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}

	l.x1 = x
	l.y1 = y
	l.x2 = tx
	l.y2 = ty
	return l
}

type Floor struct {
	field map[string]int
	lines []Line
}

func (b *Floor) Get(x, y int) (int, error) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	value, ok := b.field[index]
	if !ok {
		return 0, GetError{}
	}
	return value, nil
}

func (b *Floor) Set(x, y, num int) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	b.field[index] = num
}

func (b *Floor) Increment(x, y int) {
	g, err := b.Get(x, y)
	if err != nil {
		g = 0
	}
	g += 1
	b.Set(x, y, g)
}

func CreateFloor() *Floor {
	b := &Floor{}
	b.field = make(map[string]int)
	return b
}

func sign(a, b int) int {
	if a < b {
		return 1
	}
	return -1
}

func (b *Floor) Line(x1, y1, x2, y2 int) {
	if x1 == x2 && y1 == y2 {
		b.Increment(x1, y1)
		return
	}
	if y1 == y2 {
		sign := sign(x1, x2)
		x := x1
		b.Increment(x, y1)

		for x != x2 {
			x += sign
			b.Increment(x, y1)
		}
		return
	}
	if x1 == x2 {
		sign := sign(y1, y2)
		y := y1
		b.Increment(x1, y)
		for y != y2 {
			y += sign
			b.Increment(x1, y)
		}
		return
	}

}

func drawLines() {
	for _, v := range floor.lines {
		floor.Line(v.x1, v.y1, v.x2, v.y2)
	}
}

func countOverlap() int {
	count := 0
	for _, v := range floor.field {
		if v >= 2 {
			count++
		}
	}
	return count
}

var floor Floor

func readFile() {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	floor = *CreateFloor()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		p1 := strings.Split(fields[0], ",")
		p2 := strings.Split(fields[2], ",")
		lineObj := CreateLine(p1[0], p1[1], p2[0], p2[1])
		floor.lines = append(floor.lines, lineObj)
	}
}

func main() {
	readFile()
	drawLines()
	fmt.Println(countOverlap())
}
