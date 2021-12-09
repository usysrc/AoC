package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE = "input.txt"

func readFile() (signal, digits []string) {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	signal = make([]string, 0)
	digits = make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Split(line, "|")
		fmt.Println(fields[1])
		signal = append(signal, strings.Split(fields[0], " ")...)
		digits = append(digits, strings.Split(fields[1], " ")...)
	}

	return signal, digits
}

func main() {
	_, digits := readFile()
	count := 0
	for _, v := range digits {
		if len(v) == 2 || len(v) == 4 || len(v) == 3 || len(v) == 7 {
			fmt.Println(v)

			count++
		}
	}
	fmt.Println(count)
}
