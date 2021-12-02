package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type mov struct {
	command string
	amount  int
}

func readFile() []mov {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(inputFile)
	reader.Comma = ' '
	csvLines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var data []mov

	for _, line := range csvLines {

		number, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		data = append(data, mov{
			command: line[0],
			amount:  number,
		})

	}
	return data
}

func main() {
	data := readFile()
	aim := 0
	horizontal := 0
	vertical := 0
	for _, v := range data {
		if v.command == "up" {
			aim -= v.amount
		}
		if v.command == "down" {
			aim += v.amount
		}
		if v.command == "forward" {
			horizontal += v.amount
			vertical += v.amount * aim
		}
	}
	fmt.Println(horizontal * vertical)

}
