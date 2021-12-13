package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE = "input.txt"

type Point struct {
	x int
	y int
}

func CreatePoint(i, j int) Point {
	return Point{x: i, y: j}
}

func (p Point) GetIndex() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

type Board struct {
	field map[string]*Point
}

func CreateBoard() Board {
	b := Board{}
	b.field = make(map[string]*Point)
	return b
}

func (b *Board) Get(x, y int) (*Point, bool) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	v, ok := b.field[index]
	return v, ok
}

func (b *Board) Set(x, y int, p *Point) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	b.field[index] = p
}

func (b *Board) Remove(x, y int) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	delete(b.field, index)
}

type Instruction struct {
	dir string
	val int
}

var instructions []*Instruction

func readFile() Board {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	board := CreateBoard()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "fold along y") {
			ys := line[13:]
			y, err := strconv.Atoi(ys)
			if err != nil {
				fmt.Errorf("error: %v", err)
			}
			instruction := &Instruction{dir: "y", val: y}
			instructions = append(instructions, instruction)
		} else if strings.Contains(line, "fold along x") {
			xs := line[13:]
			x, err := strconv.Atoi(xs)
			if err != nil {
				fmt.Errorf("error: %v", err)
			}
			instruction := &Instruction{dir: "x", val: x}
			instructions = append(instructions, instruction)
		} else {
			points := strings.Split(line, ",")
			if len(points) > 1 {
				is := string(points[0])
				i, err := strconv.Atoi(is)
				if err != nil {
					fmt.Errorf("error: %v", err)
				}
				js := string(points[1])
				j, err := strconv.Atoi(js)
				if err != nil {
					fmt.Errorf("error: %v", err)
				}
				board.Set(i, j, &Point{x: i, y: j})
				// fmt.Println(i, j, len(board.field), points)
			}
		}
	}
	return board
}

func main() {
	board := readFile()
	for _, instruction := range instructions {
		if instruction.dir == "x" {
			for _, v := range board.field {
				if v.x > instruction.val {
					x := instruction.val - (v.x - instruction.val)
					board.Set(x, v.y, &Point{x: x, y: v.y})
					board.Remove(v.x, v.y)
				}
			}
		} else if instruction.dir == "y" {
			for _, v := range board.field {
				if v.y > instruction.val {
					y := instruction.val - (v.y - instruction.val)
					board.Set(v.x, y, &Point{x: v.x, y: y})
					board.Remove(v.x, v.y)
				}
			}
		}

		fmt.Println(len(board.field))
	}

	for j := 0; j < 10; j++ {
		line := ""
		for i := 0; i < 40; i++ {
			if _, ok := board.Get(i, j); ok {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

}
