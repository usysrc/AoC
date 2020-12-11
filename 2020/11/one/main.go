package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func seating(a [][]rune, b [][]rune, w int, h int) [][]rune {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			r := a[j][i]
			if r == 'L' {
				found := false
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if (j+y >= 0) && (j+y < h) && (i+x >= 0) && (i+x < w) {
							if a[j+y][i+x] == '#' {
								found = true
							}
						}
					}
				}
				if !found {
					// fmt.Println(found, i, j)
					b[j][i] = '#'
				}
			} else if r == '#' {
				found := 0
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if (x != 0 || y != 0) && j+y >= 0 && j+y < h && i+x >= 0 && i+x < w {
							if a[j+y][i+x] == '#' {
								found++
							}
						}
					}
				}
				if found >= 4 {
					// fmt.Println(found, i, j)
					b[j][i] = 'L'
				}
			}

		}
	}
	return b
}

func count(a [][]rune, r rune) int {
	counter := 0
	for _, k := range a {
		for _, seat := range k {
			if seat == r {
				counter++
			}
		}
	}
	return counter
}

func different(a [][]rune, b [][]rune, w int, h int) bool {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			if a[j][i] != b[j][i] {
				return true
			}
		}

	}
	return false
}

func cp(a [][]rune, w int, h int) [][]rune {
	b := make([][]rune, 0)
	for j := 0; j < h; j++ {
		b = append(b, make([]rune, 0))
		for i := 0; i < w; i++ {
			b[j] = append(b[j], a[j][i])
		}

	}
	return b
}

func main() {
	a := make([][]rune, 0)
	b := make([][]rune, 0)
	w := 0
	h := 0

	for j, line := range readFile() {
		a = append(a, make([]rune, 0))
		b = append(b, make([]rune, 0))
		for i, seat := range line {
			a[j] = append(a[j], seat)
			b[j] = append(b[j], seat)
			w = i + 1
		}
		h = j + 1
	}
	for {
		b = seating(a, b, w, h)
		fmt.Println("")
		for k := 0; k < h; k++ {
			fmt.Println(string(b[k]))
		}

		if different(a, b, w, h) {
			a = cp(b, w, h)
		} else {
			break
		}
	}
	fmt.Println(count(b, '#'))
	fmt.Println(w, h)

}
