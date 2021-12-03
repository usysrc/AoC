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

func oxygen(data []string) string {
	for k := 0; k < len(data[0]); k++ {
		zeros := 0
		ones := 0
		new := []string{}
		for _, v := range data {
			if v[k:k+1] == "1" {
				ones++
			}
			if v[k:k+1] == "0" {
				zeros++
			}
		}
		for _, v := range data {
			if zeros == ones && v[k:k+1] == "1" {
				new = append(new, v)
			}
			if zeros > ones && v[k:k+1] == "0" {
				new = append(new, v)
			}
			if ones > zeros && v[k:k+1] == "1" {
				new = append(new, v)
			}
		}
		data = new
		if len(data) == 1 {
			break
		}
	}
	return data[0]
}

func co(data []string) string {
	for k := 0; k < len(data[0]); k++ {
		zeros := 0
		ones := 0
		new := []string{}
		for _, v := range data {
			if v[k:k+1] == "1" {
				ones++
			}
			if v[k:k+1] == "0" {
				zeros++
			}
		}
		for _, v := range data {
			if zeros == ones && v[k:k+1] == "0" {
				new = append(new, v)
			}
			if zeros < ones && v[k:k+1] == "0" {
				new = append(new, v)
			}
			if ones < zeros && v[k:k+1] == "1" {
				new = append(new, v)
			}
		}
		data = new
		if len(data) == 1 {
			break
		}
	}
	return data[0]
}

func main() {
	data := readFile()
	oxygen_generator_rating := oxygen(data)
	co_scrubber_rating := co(data)
	fmt.Println(oxygen_generator_rating, co_scrubber_rating)
	oxygen_dec, err := strconv.ParseInt(oxygen_generator_rating, 2, 64)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	co_dec, err := strconv.ParseInt(co_scrubber_rating, 2, 64)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	fmt.Println(oxygen_dec * co_dec)
}
