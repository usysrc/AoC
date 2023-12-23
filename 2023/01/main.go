package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		l := -1
		r := -1
		for _, c := range line {
			i, err := strconv.Atoi(string(c))
			if err == nil {
				if l == -1 {
					l = i
				}
				r = i
			}
		}
		// fmt.Println(10*l + r)
		if r >= 0 && l >= 0 {
			sum += 10*l + r
		}
	}
	fmt.Println(sum)
}
