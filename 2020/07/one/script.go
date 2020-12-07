package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func splitInto(str string, sep string) (string, string) {
	s := strings.Split(str, sep)
	return s[0], s[1]
}

func run(outsideBag string, m map[string][] string) bool {
	if outsideBag == "shiny gold" {
		return true
	}

	for _, bag := range m[outsideBag] {
		if run(bag, m) {
			return true
		}
	}

	return false
}

func main() {
	m := make(map[string][]string)

	for _, ln := range readFile() {
		left, right := splitInto(ln, " contain" )
		leftSlice := strings.Split(left, " ")
		outsideBag := strings.Join(leftSlice[0:2], " ")
		rightSlice := strings.Split(right, ",")

		var arr []string
		for _, insideBag := range rightSlice {
			bagSlice := strings.Split(insideBag, " ")
			bagName := strings.Join(bagSlice[2:4], " ")
			arr = append(arr, bagName)
		}
		m[outsideBag] = arr
	}
	count := 0
	for outside, _ := range m {
		if outside != "shiny gold" && run(outside, m) {
			count++
		}
	}
	fmt.Println(count)
}