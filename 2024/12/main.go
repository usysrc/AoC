package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	X, Y int
}

type Map struct {
	data          map[Pos]string
	width, height int
	count         int
	crop          string
}

func getAreas(crop string, m Map) []Map {
	areas := []Map{}
	currentArea := Map{data: map[Pos]string{}, crop: crop}
	visited := map[Pos]bool{}
	for pos, v := range m.data {
		if !visited[pos] && v == crop {
			// do a flood fill to find all adjacent tiles of the same type
			// and add them to the current area
			var fill func(Pos)
			fill = func(p Pos) {
				if visited[p] {
					return
				}
				if m.data[p] == crop {
					currentArea.count++
					currentArea.data[p] = crop
					visited[p] = true
					for _, d := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
						fill(Pos{p.X + d.X, p.Y + d.Y})
					}
				}
			}
			fill(pos)
			areas = append(areas, currentArea)
			currentArea = Map{data: map[Pos]string{}, crop: crop}
		}
	}
	return areas
}

// CountHull counts the number of tiles around the shape of x
func CountHullAround(crop string, m Map) int {
	// get all areas of the crop
	areas := getAreas(crop, m)
	// count the number of hull tiles around each area
	count := 0
	// iterate over map data
	for _, area := range areas {
		areaHullCount := 0
		for pos, c := range area.data {
			if c == crop {
				// check all 4 directions
				for _, d := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					newPos := Pos{pos.X + d.X, pos.Y + d.Y}
					if k, ok := m.data[newPos]; !ok || k != crop {
						areaHullCount++
					}
				}
			}
		}
		fmt.Println("Area count for", area.crop, area.count, "*", areaHullCount)
		count += areaHullCount * area.count
	}
	return count
}

func countSides(crop string, m Map) int {
	// get all areas of the crop
	areas := getAreas(crop, m)
	// count the number of sides around each area
	count := 0
	visited := map[Pos]bool{}
	for _, area := range areas {
		sides := 0
		// count the number of corners
		for pos, c := range area.data {
			if visited[pos] {
				continue
			}
			if c == crop {
				// check all 4 directions
				for _, d := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					newPos := Pos{pos.X + d.X, pos.Y + d.Y}
					if k, ok := m.data[newPos]; !ok || k != crop {
						// count neighbor that are not crop
						neighbors := 0
						for _, d2 := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
							if k, ok := m.data[Pos{newPos.X + d2.X, newPos.Y + d2.Y}]; !ok || k != crop {
								neighbors++
							}
						}
						if neighbors == 2 {
							sides += 2
						}
						if neighbors == 3 {
							sides += 3
						}
					}
				}
			}
		}

		fmt.Println("Area count for", area.crop, area.count, "*", sides)
		count += sides * area.count
	}
	return count
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	count := map[string]int{}
	m := Map{data: map[Pos]string{}}

	scanner := bufio.NewScanner(file)
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Do something with the line
		for i, r := range line {
			str := string(r)
			count[str]++
			m.data[Pos{i, j}] = str
		}
		j++
	}
	// Part A
	sumA := 0
	for k, _ := range count {
		hullCount := CountHullAround(k, m)
		println(k, hullCount)
		sumA += hullCount
	}
	fmt.Println("Part A:", sumA)

	// Part B
	sumB := 0
	for k, _ := range count {
		sides := countSides(k, m)
		println(k, sides)
		sumB += sides
	}
	fmt.Println("Part B:", sumB)
}
