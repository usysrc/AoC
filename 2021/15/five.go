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
	// if _, ok := b.Get(x, y-1); ok {
	// 	vals = append(vals, CreatePoint(x+1, y))
	// }
	if _, ok := b.Get(x+1, y); ok {
		vals = append(vals, CreatePoint(x+1, y))
	}
	// if _, ok := b.Get(x-1, y); ok {
	// 	vals = append(vals, CreatePoint(x+1, y))
	// }
	if _, ok := b.Get(x, y+1); ok {
		vals = append(vals, CreatePoint(x, y+1))
	}
	return vals
}

func readFile() Board {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	j := 0
	board := CreateBoard()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i, v := range line {
			k, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Errorf("error: %v", err)
			}
			board.Set(i, j, k)
		}
		j++
	}
	return board
}

func sumRisk(board Board, visited Board, i, j int) int {
	if v, ok := visited.Get(i, j); ok {
		return v
	}
	risk := 0
	n := board.GetNeigh(i, j)
	min := 1000000
	for _, p := range n {
		val := p.GetValue(board)
		sum := sumRisk(board, visited, p.x, p.y)
		if val+sum < min {
			min = val + sum
		}
	}

	if min != 1000000 {
		risk += min
	}
	visited.Set(i, j, risk)
	return risk
}

func extendBoard(b Board) Board {
	w := 100
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for ii := 0; ii < w; ii++ {
				for jj := 0; jj < w; jj++ {
					if val, ok := b.Get(ii, jj); ok {
						v := val + i + j
						for v > 9 {
							v -= 9
						}
						b.Set(i*w+ii, j*w+jj, v)
					}
				}
			}
		}
	}
	return b
}

func main() {
	board := readFile()
	board = extendBoard(board)
	risk := sumRisk(board, CreateBoard(), 0, 0)
	fmt.Println(risk)
}
