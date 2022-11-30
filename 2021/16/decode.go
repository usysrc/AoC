package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE = "test.txt"

func readFile() string {
	var m = map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	input := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for _, v := range line {
			input = input + m[string(v)]
		}
	}
	return input
}

func get(s string, a, b int) (int64, string) {
	bits := s[a:b]
	number, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return number, bits[b:]
}

func main() {
	input := readFile()
	// fmt.Println(input)
	for input != "" {
		version, input := get(input, 0, 3)
		typeid, input := get(input, 0, 3)
		if typeid == 4 { // literal value
			fmt.Println(version, typeid)
		} else {
			// typeid := get(input, 3, 6)
			// fmt.Println(typeid)
		}
	}
}
