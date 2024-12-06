package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction string

type Pos struct {
	x int
	y int
}

type Guard struct {
	x   int
	y   int
	dir Direction
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2d map of the floor with "." being an open space and "#" being a obstacle
	floorMap := map[Pos]string{}
	// guard struct to keep track of the guards position and direction
	guard := Guard{0, 0, "up"}
	start := Pos{0, 0}
	// Read the file and populate the floorMap, set the guards starting position
	scanner := bufio.NewScanner(file)
	mw, mh := 0, 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Do something with the line
		for i, c := range line {
			if string(c) == "^" {
				floorMap[Pos{i, j}] = "."
				start.x = i
				start.y = j
				continue
			}
			floorMap[Pos{i, j}] = string(c)
			mw = i + 1
		}
		j++
		mh = j
	}

	guard.x = start.x
	guard.y = start.y
	visited := map[Pos]bool{}

	// mark the starting position as visited
	visited[Pos{guard.x, guard.y}] = true
	count := 1

	// simulate movement until guard leaves the floor, mark visited spaces
	for {
		gx, gy := guard.x, guard.y
		if guard.dir == "up" {
			gy--
		} else if guard.dir == "down" {
			gy++
		} else if guard.dir == "left" {
			gx--
		} else if guard.dir == "right" {
			gx++
		}

		// check if guard is out of bounds
		if gy < 0 || gy >= mh || gx < 0 || gx >= mw {
			break
		}

		// check if guard has hit a wall
		if floorMap[Pos{gx, gy}] == "#" {
			if guard.dir == "up" {
				guard.dir = "right"
			} else if guard.dir == "down" {
				guard.dir = "left"
			} else if guard.dir == "left" {
				guard.dir = "up"
			} else if guard.dir == "right" {
				guard.dir = "down"
			}
		} else {
			guard.x = gx
			guard.y = gy

			// check if guard has visited this space before
			if !visited[Pos{guard.x, guard.y}] {
				visited[Pos{guard.x, guard.y}] = true
				count++
			}
		}
	}

	// find positions for new obstacles on the map that would create loops for the guard
	loopCount := 0
	dir := Direction("up")
	for pos, _ := range visited {
		tx, ty := pos.x, pos.y
		changed := false
		if floorMap[Pos{tx, ty}] == "." {
			floorMap[Pos{tx, ty}] = "#"
			changed = true
		}
		// simulate movement until guard leaves the floor or lands in a loop, mark visited spaces
		guard.x, guard.y = start.x, start.y
		guard.dir = Direction("up")
		visitedLoop := map[Pos]Direction{}
		visitedLoop[Pos{guard.x, guard.y}] = guard.dir
		for {
			gx, gy := guard.x, guard.y
			if guard.dir == "up" {
				gy--
			} else if guard.dir == "down" {
				gy++
			} else if guard.dir == "left" {
				gx--
			} else if guard.dir == "right" {
				gx++
			}

			// check if guard is out of bounds
			if gy < 0 || gy >= mh || gx < 0 || gx >= mw {
				break
			}

			// check if guard has hit a wall
			if floorMap[Pos{gx, gy}] == "#" {
				if guard.dir == "up" {
					guard.dir = "right"
				} else if guard.dir == "down" {
					guard.dir = "left"
				} else if guard.dir == "left" {
					guard.dir = "up"
				} else if guard.dir == "right" {
					guard.dir = "down"
				}
				continue
			} else {
				guard.x = gx
				guard.y = gy

				// check if guard has visited this space before
				if visitedLoop[Pos{guard.x, guard.y}] == "" || visitedLoop[Pos{guard.x, guard.y}] != guard.dir {
					visitedLoop[Pos{guard.x, guard.y}] = guard.dir
				} else if visitedLoop[Pos{guard.x, guard.y}] == guard.dir {
					// found a loop
					loopCount++
					break
				}
			}
		}
		// restore floormap if changed
		if changed {
			floorMap[Pos{tx, ty}] = "."
		}

		// change direction
		if floorMap[Pos{tx, ty}] == "#" {
			if dir == "up" {
				dir = "right"
			} else if dir == "down" {
				dir = "left"
			} else if dir == "left" {
				dir = "up"
			} else if dir == "right" {
				dir = "down"
			}
			continue
		}
	}

	// draw the floor and the visited spaces
	// for y := 0; y < mh; y++ {
	// 	for x := 0; x < mw; x++ {
	// 		if visited[Pos{x, y}] {
	// 			fmt.Print("X")
	// 		} else {
	// 			fmt.Print(floorMap[Pos{x, y}])
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println("Count A:", count)
	fmt.Println("Count B:", loopCount)
}
