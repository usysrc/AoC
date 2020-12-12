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

	wx := 10
	wy := -1

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
			wy -= val
		}
		if operator == 'S' {
			wy += val
		}
		if operator == 'E' {
			wx += val
		}
		if operator == 'W' {
			wx -= val
		}
		if operator == 'R' {
			cur := math.Atan2(float64(wy), float64(wx))
			fmt.Println(cur)
			radians := float64(val) * math.Pi / 180.0
			dx := float64(wx)
			dy := float64(wy)
			len := math.Sqrt(dx*dx + dy*dy)
			wx = int(math.Round(len * math.Cos(cur+radians)))
			wy = int(math.Round(len * math.Sin(cur+radians)))
		}
		if operator == 'L' {
			cur := math.Atan2(float64(wy), float64(wx))
			radians := float64(val) * math.Pi / 180.0
			dx := float64(wx)
			dy := float64(wy)
			len := math.Sqrt(dx*dx + dy*dy)
			wx = int(math.Round(len * math.Cos(cur-radians)))
			wy = int(math.Round(len * math.Sin(cur-radians)))
		}
		if operator == 'F' {
			x += val * wx
			y += val * wy
		}

		log.Println(string(operator), val, "x:", x, "y:", y, dir, wx, wy)

	}
	fmt.Println(x, y, dir)
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))

}
