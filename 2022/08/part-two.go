package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() int {
	w, h := 99, 99
	var height [99][99]int
	var visible [99][99]bool
	inputFile, err := os.Open("./input")

	// w, h := 5, 5
	// var height [5][5]int
	// var visible [5][5]bool
	// inputFile, err := os.Open("./test")
	if err != nil {
		panic(err)
	}

	j := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for k, v := range line {
			height[j][k], err = strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			visible[j][k] = false
		}
		j++
	}

	max := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			val := height[j][i]
			right := 0
			for ii := i + 1; ii < w; ii++ {
				right++
				if height[j][ii] >= val {
					break
				}
			}

			left := 0
			for ii := i - 1; ii >= 0; ii-- {
				left++
				if height[j][ii] >= val {
					break
				}
			}

			down := 0
			for jj := j + 1; jj < h; jj++ {
				down++
				if height[jj][i] >= val {
					break
				}
			}

			up := 0
			for jj := j - 1; jj >= 0; jj-- {
				up++
				if height[jj][i] >= val {
					break
				}
			}
			fmt.Println(left, right, up, down)
			if left*right*up*down > max {
				max = left * right * up * down
			}
		}
	}

	return max
}

func main() {
	count := readFile()
	fmt.Println("result is: ", count)
}
