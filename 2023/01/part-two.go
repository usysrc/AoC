package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		l := -1
		r := -1
		word := ""
		for _, c := range line {
			if unicode.IsLetter(c) {
				word += string(c)
			}
			for w, num := range numbers {
				if strings.Contains(word, w) {
					word = string(c)
					if l == -1 {
						l = num
					}
					r = num
				}
			}

			i, err := strconv.Atoi(string(c))
			if err == nil {
				word = ""
				if l == -1 {
					l = i
				}
				r = i
			}
		}
		if r >= 0 && l >= 0 {
			fmt.Println(10*l + r)
			sum += 10*l + r
		}
	}
	fmt.Println(sum)
}
