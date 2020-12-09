package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {

	lines := make([]int, 0)

	for _, line := range readFile() {
		val, err := strconv.Atoi(line)
		if err != nil{
			log.Fatal(err)
		}
		lines = append(lines, val)
	}

	p := 0
	q := 25
	cache := make(map[int]bool)

	for i := 0; i < q; i++ {
		cache[lines[i]] = true
	}

	for q < len(lines) {

		found := false
		for k,_ := range cache{
			fmt.Println(q, k, lines[q], lines[q] - k, cache[lines[q] - k])
			if cache[lines[q] - k]{
				found = true
			}
		}
		if !found {
			fmt.Println(lines[q])
			return
		}
		fmt.Println("")
		// remove value at the end
		delete(cache, lines[p])
		p++

		cache[lines[q]] = true
		q++

	}
}