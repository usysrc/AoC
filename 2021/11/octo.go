package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE = "input.txt"

type Point struct {
	x int
	y int
}

func CreatePoint(i, j int) Point {
	return Point{x: i, y: j}
}

func (p Point) GetValue(b Board) int {
	val, ok := b.Get(p.x, p.y)
	if !ok {
		fmt.Errorf("Value not found")
	}
	return val
}

func (p Point) GetIndex() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

type Board struct {
	field map[string]int
}

func CreateBoard() Board {
	b := Board{}
	b.field = make(map[string]int)
	return b
}

func (b *Board) Get(x, y int) (int, bool) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	v, ok := b.field[index]
	return v, ok
}

func (b *Board) Set(x, y, num int) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	b.field[index] = num
}

func (b Board) GetNeigh(x, y int) []Point {
	vals := make([]Point, 0)
	if _, ok := b.Get(x-1, y-1); ok {
		vals = append(vals, CreatePoint(x-1, y-1))
	}
	if _, ok := b.Get(x, y-1); ok {
		vals = append(vals, CreatePoint(x, y-1))
	}
	if _, ok := b.Get(x+1, y-1); ok {
		vals = append(vals, CreatePoint(x+1, y-1))
	}
	if _, ok := b.Get(x-1, y); ok {
		vals = append(vals, CreatePoint(x-1, y))
	}
	if _, ok := b.Get(x+1, y); ok {
		vals = append(vals, CreatePoint(x+1, y))
	}
	if _, ok := b.Get(x-1, y+1); ok {
		vals = append(vals, CreatePoint(x-1, y+1))
	}
	if _, ok := b.Get(x, y+1); ok {
		vals = append(vals, CreatePoint(x, y+1))
	}
	if _, ok := b.Get(x+1, y+1); ok {
		vals = append(vals, CreatePoint(x+1, y+1))
	}
	return vals
}

func (b Board) Flash(i, j int) int {
	stack := []Point{}
	stack = append(stack, CreatePoint(i, j))
	flashes := 0
	for len(stack) > 0 {
		fmt.Println("flash", len(stack))
		flashes++
		n := len(stack) - 1
		cur := stack[n]
		stack = stack[:n]

		for _, p := range b.GetNeigh(cur.x, cur.y) {
			if p.GetValue(b) != 0 {
				b.Set(p.x, p.y, p.GetValue(b)+1)
			}
			if p.GetValue(b) > 9 {
				b.Set(p.x, p.y, 0)
				stack = append(stack, p)
			}
		}
	}
	return flashes
}

func (b Board) Step() int {
	for k, v := range b.field {
		b.field[k] = v + 1
	}
	flashes := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			v, ok := b.Get(i, j)
			if ok {
				if v > 9 {
					b.Set(i, j, 0)
					flashes += b.Flash(i, j)
				}
			}
		}
	}
	return flashes
}

func (b Board) AllFlash() bool {
	count := 0
	for _, v := range b.field {
		if v == 0 {
			count++
		}
	}
	return count == len(b.field)
}

func readFile() Board {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	board := CreateBoard()
	j := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i := 0; i < len(line); i++ {
			num := string(line[i])
			number, err := strconv.Atoi(num)
			if err != nil {
				fmt.Errorf("error: %v", err)
			}
			board.Set(i, j, number)
		}
		j++
	}
	return board
}

func main() {
	board := readFile()
	flashes := 0
	i := 0
	for !board.AllFlash() {
		i++
		flashes += board.Step()
	}
	fmt.Println(i, flashes)
}
