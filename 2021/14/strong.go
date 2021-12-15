package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE = "test.txt"

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

func main() {
	subs, start := readFile()
	cur := start
	counts := make(map[string]int)
	countChar := make(map[string]int)

	for i := 0; i < len(cur)-1; i++ {
		c := cur[i : i+2]
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c]++
	}
	fmt.Println(counts)

	for i := 0; i < 10; i++ {
		newcounts := make(map[string]int)
		for k, num := range counts {
			newcounts[k] = num
		}
		for k, num := range counts {
			c1 := string(k[0:1]) + subs[k]
			newcounts[c1] += num
			c2 := subs[k] + string(k[1:2])
			newcounts[c2] += num
			// fmt.Println(c1, c2)
			newcounts[k] -= num
		}
		counts = newcounts
		fmt.Println(i, counts)
	}

	for k, num := range counts {
		for _, c := range k {
			countChar[string(c)] += num
		}
	}

	min := 10000000000
	max := 0
	max_c := ""
	min_c := ""
	for k, v := range countChar {
		if v > max {
			max = v
			max_c = k
		}
		if v < min {
			min = v
			min_c = k
		}
	}
	fmt.Println(countChar)
	fmt.Println(max_c, min_c, max, min, max-min)
}
