package main

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
	for {
		for _, id := range ids {
			if time%id == 0 {
				fmt.Println(time, id, (time-starttime)*id)
				return
			}
		}
		time++
	}

}

func main() {
	lines := readFile()
	starttime, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}
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
		}
	}

	run(starttime, ids)
}
