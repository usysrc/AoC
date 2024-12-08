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
}

func (m Map) copy() Map {
	newMap := Map{data: make(map[Pos]string), width: m.width, height: m.height}
	for k, v := range m.data {
		newMap.data[k] = v
	}
	return newMap
}

func PartA(karte Map, antennas map[string][]Pos) {
	count := 0
	for _, antenna := range antennas {
		for _, pos := range antenna {
			// for each antenna that has the same label
			for _, otherpos := range antenna {
				// if the two antennas are not the same
				if pos != otherpos {
					// project the line between the two antennas and extend it once by the distance of the two antennas to each other, put a # at the projected end position
					dx := otherpos.X - pos.X
					dy := otherpos.Y - pos.Y
					i := 2
					pos := Pos{pos.X + i*dx, pos.Y + i*dy}
					if pos.X < 0 || pos.X >= karte.width || pos.Y < 0 || pos.Y >= karte.height {
						continue
					}
					if karte.data[pos] == "." {
						karte.data[pos] = "#"
						count++
					} else if karte.data[pos] != "." && karte.data[pos] != "#" {
						count++
					}
				}
			}
		}
	}

	// Print the map
	for j := 0; j < karte.width; j++ {
		for i := 0; i < karte.height; i++ {
			fmt.Print(karte.data[Pos{i, j}])
		}
		fmt.Println()
	}

	fmt.Println("Count A:", count)
}

func PartB(karte Map, antennas map[string][]Pos) {
	count := 0
	for _, antenna := range antennas {
		for _, pos := range antenna {
			// for each antenna that has the same label
			for _, otherpos := range antenna {
				// if the two antennas are not the same
				if pos != otherpos {
					// project the line between the two antennas and extend it once by the distance of the two antennas to each other, put a # at the projected end position
					dx := otherpos.X - pos.X
					dy := otherpos.Y - pos.Y
					for i := 1; i < 1000; i++ {
						pos := Pos{pos.X + i*dx, pos.Y + i*dy}
						if pos.X < 0 || pos.X >= karte.width || pos.Y < 0 || pos.Y >= karte.height {
							continue
						}
						if karte.data[pos] == "." {
							karte.data[pos] = "#"
							count++
						} else if karte.data[pos] != "." && karte.data[pos] != "#" {
							karte.data[pos] = "#"
							count++
						}
					}
				}
			}
		}
	}

	// Print the map
	fmt.Println()
	for j := 0; j < karte.width; j++ {
		for i := 0; i < karte.height; i++ {
			fmt.Print(karte.data[Pos{i, j}])
		}
		fmt.Println()
	}

	fmt.Println("Count B:", count)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	karte := Map{data: make(map[Pos]string)}
	antennas := map[string][]Pos{}

	// Read the file
	j := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			karte.data[Pos{i, j}] = string(c)
			if i > karte.width {
				karte.width = i
			}
			str := string(c)
			if str != "." {
				antennas[str] = append(antennas[str], Pos{i, j})
			}
		}
		j++
		karte.height = j
	}
	karte.width++

	karteA := karte.copy()
	PartA(karteA, antennas)

	karteB := karte.copy()
	PartB(karteB, antennas)

}
