package main

// this is broken!

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func run(starttime int, ids []int) {
	time := starttime
	step := ids[0]
	for i, id := range ids {
		if id < 0 && i > 0 {
			continue
		}
		j := 0
		for {
			j++
			fmt.Println(j)
			k := time + j*step
			if (k+i)%id == 0 {
				time = k
				step *= id
				break
			}
		}
		// fmt.Println(id)
	}
	fmt.Println(time)
}

func main() {
	lines := readFile()
	notes := lines[1]
	notesarr := strings.Split(notes, ",")

	ids := make([]int, 0)
	for _, note := range notesarr {
		if note != "x" {
			id, err := strconv.Atoi(string(note))
			if err != nil {
				log.Fatal(err)
			}
			ids = append(ids, id)
		} else if note == "x" {
			ids = append(ids, -1)
		}
	}

	run(0, ids)
}
