package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARDWIDTH = 5
const BOARDHEIGHT = 5

type BoardNumber struct {
	number int
	marked bool
}

type Board struct {
	field map[string]BoardNumber
}

func (b *Board) Get(x, y int) BoardNumber {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	return b.field[index]
}

func (b *Board) Set(x, y, num int) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	b.field[index] = BoardNumber{
		number: num,
		marked: false,
	}
}

func (b *Board) Mark(num int) {
	for k, v := range b.field {
		if v.number == num {
			if boardNumber, ok := b.field[k]; ok {
				boardNumber.marked = true
				b.field[k] = boardNumber
			}
		}
	}
}

func (b *Board) Won() bool {
	// check horizontal
	for j := 0; j < BOARDHEIGHT; j++ {
		count := 0
		for i := 0; i < BOARDWIDTH; i++ {
			if b.Get(i, j).marked {
				count++
			}
		}
		if count == BOARDWIDTH {
			return true
		}
	}

	// check vertical
	for i := 0; i < BOARDWIDTH; i++ {
		count := 0
		for j := 0; j < BOARDHEIGHT; j++ {
			if b.Get(i, j).marked {
				count++
			}
		}
		if count == BOARDHEIGHT {
			return true
		}
	}
	return false
}

func (b *Board) UnmarkedSum() int {
	sum := 0
	for _, v := range b.field {
		if !v.marked {
			sum += v.number
		}
	}
	return sum
}

func CreateBoard() *Board {
	b := &Board{}
	b.field = make(map[string]BoardNumber)
	return b
}

var gameNumbers []int
var boards []*Board

func readGameNumbers(line string) {
	s := strings.Split(line, ",")
	for _, num := range s {
		number, err := strconv.Atoi(num)
		if err != nil {
			fmt.Errorf("error: %s", err)
		}
		gameNumbers = append(gameNumbers, number)
	}
}

func readFile() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	k := 0
	j := 0
	// board := CreateBoard()
	var board *Board
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		// fmt.Println("=======")
		if k == 0 {
			readGameNumbers(line)
		} else {
			if line == "" {
				if board != nil {
					boards = append(boards, board)
				}
				board = CreateBoard()
				j = 0
			} else {
				for i, num := range strings.Fields(line) {
					number, err := strconv.Atoi(num)
					if err != nil {
						fmt.Errorf("error: %v", err)
					}
					board.Set(i, j, number)
				}
				j++
			}

		}
		k++
	}
	boards = append(boards, board)
}

func runGame() int {
	var stack []int
	for _, v := range gameNumbers {
		var new_boards []*Board
		for _, board := range boards {
			board.Mark(v)
			if board.Won() {
				fmt.Println(v, board.UnmarkedSum())
				stack = append(stack, v*board.UnmarkedSum())
			} else {
				new_boards = append(new_boards, board)
			}
		}
		boards = new_boards
	}
	return stack[len(stack)-1]
}

func main() {
	readFile()
	fmt.Println(len(boards))
	res := runGame()
	fmt.Println(res)
}
