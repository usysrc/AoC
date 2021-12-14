package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE = "input.txt"

func readFile() (map[string]string, string) {
	subs := make(map[string]string)
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	j := 0
	start := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if j == 0 {
			start = line
		}
		if j > 1 {
			parts := strings.Split(line, " -> ")
			subs[parts[0]] = parts[1]
			// fmt.Println(subs)

		}
		j++
	}
	return subs, start
}

func step(subs map[string]string, start string) string {
	cur := start
	next := ""
	fmt.Println(subs)
	for i := 0; i < len(cur)-1; i++ {
		c := cur[i : i+2]
		next = next + string(c[0]) + subs[c]
		if i == len(cur)-2 {
			next += string(c[1])
		}
	}
	return next
}

func main() {
	subs, start := readFile()
	cur := start
	for i := 0; i < 10; i++ {
		next := step(subs, cur)
		cur = next
	}
	counts := make(map[string]int)
	min := 1000000
	max := 0
	max_c := ""
	min_c := ""
	for i := 0; i < len(cur); i++ {
		c := string(cur[i])
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c]++
	}
	for k, v := range counts {
		if v > max {
			max = v
			max_c = k
		}
		if v < min {
			min = v
			min_c = k
		}
	}

	fmt.Println(counts[max_c]-counts[min_c], max_c, min_c, counts[max_c], counts[min_c])

}
