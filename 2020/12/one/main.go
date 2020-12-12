package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	x := 0
	y := 0
	dir := 0.0

	for _, line := range readFile() {
		operator := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if operator == 'N' {
			y -= val
		}
		if operator == 'S' {
			y += val
		}
		if operator == 'E' {
			x += val
		}
		if operator == 'W' {
			x -= val
		}
		if operator == 'R' {
			radians := float64(val) * math.Pi / 180.0
			dir += radians
		}
		if operator == 'L' {
			radians := float64(val) * math.Pi / 180.0
			dir -= radians
		}
		if operator == 'F' {
			x += int(float64(val) * math.Cos(float64(dir)))
			y += int(float64(val) * math.Sin(float64(dir)))
		}
	}
	fmt.Println(x, y, dir)
	fmt.Println(x + y)

}
