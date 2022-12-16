package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x        int
	y        int
	distance int
	beacon   *Point
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// get manhattan distance
func getDistance(ax, ay, bx, by int) int {
	return abs(ax-bx) + abs(ay-by)
}

// parse the input
func readFile() []Point {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	// g := NewGrid()
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			// fmt.Println(line)
			panic(err)
		}
		points = append(points, Point{
			x:        sx,
			y:        sy,
			distance: getDistance(sx, sy, bx, by),
			beacon: &Point{
				x: bx,
				y: by,
			},
		})
	}
	return points
}

func calculateBeaconCount(points []Point) int {
	lowest := 0
	highest := 4000000
	// highest := 20
	i := lowest
	j := lowest
	for true {
		found := false
		for _, point := range points {
			if getDistance(i, j, point.x, point.y) <= point.distance {
				found = true
				dx := abs(point.x - i)
				dy := abs(point.y - j)
				i += point.distance - dy + dx + 1
				break
			}
		}
		if !found {
			fmt.Println("found in:", i, j)
			return i*4000000 + j
		}
		if i > highest {
			i = 0
			j++
		}
	}
	return 0
}

func main() {
	points := readFile()
	noBeaconCount := calculateBeaconCount(points)
	fmt.Println("result is: ", noBeaconCount)
}
