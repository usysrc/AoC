package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// create 2d array
	lines := make([][]string, 0)

	for scanner.Scan() {
		line := make([]string, 0)
		for _, c := range scanner.Text() {
			char := string(c)
			line = append(line, char)
		}
		lines = append(lines, line)
	}

	countXMAS(lines)
	countXshapedMAS(lines)
}

// count all MAS in the shape of an X in the array
// M.S
// .A.
// M.S
// each MAS can be written forwards or backwards
func countXshapedMAS(lines [][]string) {
	count := 0
	for i, line := range lines {
		for j, char := range line {
			if char == "A" {
				// check left upper diagonal for M and right lower diagonal for S
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(lines) && j+1 < len(line) {
					if lines[i-1][j-1] == "M" && lines[i+1][j+1] == "S" {
						// check left lower diagonal for M and right upper diagonal for S
						if lines[i+1][j-1] == "M" && lines[i-1][j+1] == "S" {
							count++
						}
						// check left lower diagonal for S and right upper diagonal for M
						if lines[i+1][j-1] == "S" && lines[i-1][j+1] == "M" {
							count++
						}
					}
				}
				// check left upper diagonal for S and right lower diagonal for M
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(lines) && j+1 < len(line) {
					if lines[i-1][j-1] == "S" && lines[i+1][j+1] == "M" {
						// check left lower diagonal for M and right upper diagonal for S
						if lines[i+1][j-1] == "M" && lines[i-1][j+1] == "S" {
							count++
						}
						// check left lower diagonal for S and right upper diagonal for M
						if lines[i+1][j-1] == "S" && lines[i-1][j+1] == "M" {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println("Count X shaped MAS:", count)
}

func countXMAS(lines [][]string) {
	// iterate over the 2d array
	count := 0
	for i, line := range lines {
		for j, char := range line {
			// count the number of times XMAS appears in the 2d array in any direction, including diagonals
			if char == "X" {
				// check right
				if j+3 < len(line) {
					if line[j+1] == "M" && line[j+2] == "A" && line[j+3] == "S" {
						count++
					}
				}
				// check left
				if j-3 >= 0 {
					if line[j-1] == "M" && line[j-2] == "A" && line[j-3] == "S" {
						count++
					}
				}
				// check down
				if i+3 < len(lines) {
					if lines[i+1][j] == "M" && lines[i+2][j] == "A" && lines[i+3][j] == "S" {
						count++
					}
				}
				// check up
				if i-3 >= 0 {
					if lines[i-1][j] == "M" && lines[i-2][j] == "A" && lines[i-3][j] == "S" {
						count++
					}
				}
				// check diagonal right down
				if i+3 < len(lines) && j+3 < len(line) {
					if lines[i+1][j+1] == "M" && lines[i+2][j+2] == "A" && lines[i+3][j+3] == "S" {
						count++
					}
				}
				// check diagonal left down
				if i+3 < len(lines) && j-3 >= 0 {
					if lines[i+1][j-1] == "M" && lines[i+2][j-2] == "A" && lines[i+3][j-3] == "S" {
						count++
					}
				}
				// check diagonal right up
				if i-3 >= 0 && j+3 < len(line) {
					if lines[i-1][j+1] == "M" && lines[i-2][j+2] == "A" && lines[i-3][j+3] == "S" {
						count++
					}
				}
				// check diagonal left up
				if i-3 >= 0 && j-3 >= 0 {
					if lines[i-1][j-1] == "M" && lines[i-2][j-2] == "A" && lines[i-3][j-3] == "S" {
						count++
					}
				}
			}
		}
	}
	fmt.Println("Count XMAS:", count)
}
