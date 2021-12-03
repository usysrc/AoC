package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []string {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(inputFile).ReadAll()
	if err != nil {
		panic(err)
	}

	var data []string

	for _, line := range csvLines {
		for _, num := range line {
			if err != nil {
				panic(err)
			}
			data = append(data, num)
		}

	}
	return data
}

func main() {
	data := readFile()
	gamma := ""
	epsilon := ""
	for k := 0; k < len(data[0]); k++ {
		zeros := 0
		ones := 0
		for _, v := range data {
			if v[k:k+1] == "1" {
				ones++
			}
			if v[k:k+1] == "0" {
				zeros++
			}
		}
		if zeros > ones {
			gamma += "0"
		} else {
			gamma += "1"
		}
		if zeros < ones {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	gamma_dec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	epsilon_dec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	fmt.Println(gamma_dec * epsilon_dec)
}
