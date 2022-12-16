package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x        int
	y        int
	distance int
	beacon   *Point
}

// get manhattan distance
func getDistance(ax, ay, bx, by int) int {
	return int(math.Abs(float64(ax-bx)) + math.Abs(float64(ay-by)))
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
	// find min and max of x
	lowestX := 0
	highestX := 0
	for _, point := range points {
		if point.x-point.distance < lowestX {
			lowestX = point.x - point.distance
		}
		if point.x+point.distance > highestX {
			highestX = point.x + point.distance
		}
	}

	y := 2000000
	// y := 10
	count := 0
	for i := lowestX; i <= highestX; i++ {
		for _, point := range points {
			if i == point.beacon.x && y == point.beacon.y {
				// is sensor
			} else if getDistance(i, y, point.x, point.y) <= point.distance {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	points := readFile()
	noBeaconCount := calculateBeaconCount(points)
	fmt.Println("result is: ", noBeaconCount)
}
