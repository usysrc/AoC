package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
)

type Vector2 struct {
	x, y int
}

type Robot struct {
	sp Vector2 // starting position
	p  Vector2 // position
	v  Vector2 // velocity
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	robots := []Robot{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// parse the robot line, example: p=0,4 v=2,0
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.sp.x, &robot.sp.y, &robot.v.x, &robot.v.y)
		robot.p.x, robot.p.y = robot.sp.x, robot.sp.y
		fmt.Println(robot)
		robots = append(robots, robot)
	}
	// printRobots(7, 11, robots)
	fmt.Println()
	// simulate the movement of the robots for 100 seconds
	width := 101
	height := 103
	for i := 0; i < 100; i++ {
		for k, robot := range robots {
			robot.p.x += robot.v.x
			robot.p.y += robot.v.y
			// wrap around the edges
			if robot.p.x < 0 {
				robot.p.x += width
			}
			if robot.p.x >= width {
				robot.p.x -= width
			}
			if robot.p.y < 0 {
				robot.p.y += height
			}
			if robot.p.y >= height {
				robot.p.y -= height
			}
			robots[k] = robot
		}
		// print the positions of the robots
		// printRobots(height, width, robots)
		fmt.Println()
	}

	// count the number of robots in each quadrant
	quadrants := make([]int, 4)
	for i := 0; i < 4; i++ {
		quadrants[i] = 0
	}
	for _, robot := range robots {
		if robot.p.x >= 0 && robot.p.x < width && robot.p.y >= 0 && robot.p.y < height {
			fmt.Println(robot.p)
			quadrant := -1
			if robot.p.x < width/2 {
				if robot.p.y < height/2 {
					quadrant = 0
				}
			}
			if robot.p.x > width/2 {
				if robot.p.y < height/2 {
					quadrant = 1
				}
			}
			if robot.p.x < width/2 {
				if robot.p.y > height/2 {
					quadrant = 2
				}
			}
			if robot.p.x > width/2 {
				if robot.p.y > height/2 {
					quadrant = 3
				}
			}
			if quadrant >= 0 {
				quadrants[quadrant]++
			}
		}
	}

	factorA := 1
	for i := 0; i < 4; i++ {
		fmt.Println("Quadrant", i, ":", quadrants[i])
		factorA *= quadrants[i]
	}
	fmt.Println("Safety factor of Part A:", factorA)

	// Part B
	// reset the robots to their starting positions
	for k, robot := range robots {
		robot.p.x, robot.p.y = robot.sp.x, robot.sp.y
		robots[k] = robot
	}
	// simulate the movement until they are in the shape of a christmas tree
	tMap := map[Vector2][]Robot{}
	for k := 0; k < 10000; k++ {
		tMap = map[Vector2][]Robot{}
		for i, robot := range robots {
			robot.p.x += robot.v.x
			robot.p.y += robot.v.y
			// wrap around the edges
			if robot.p.x < 0 {
				robot.p.x += width
			}
			if robot.p.x >= width {
				robot.p.x -= width
			}
			if robot.p.y < 0 {
				robot.p.y += height
			}
			if robot.p.y >= height {
				robot.p.y -= height
			}
			robots[i] = robot
			tMap[robot.p] = append(tMap[robot.p], robot)
		}
		found := false
		for _, robot := range robots {
			a := tMap[Vector2{robot.p.x - 1, robot.p.y + 1}]
			b := tMap[Vector2{robot.p.x, robot.p.y + 1}]
			c := tMap[Vector2{robot.p.x + 1, robot.p.y + 1}]
			if len(a) > 0 && len(b) > 0 && len(c) > 0 {
				found = true
				break
			}
		}
		if found {
			// save a png of the map
			img := image.NewRGBA(image.Rect(0, 0, width, height))
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					if len(tMap[Vector2{j, i}]) > 0 {
						img.Set(j, i, image.White)
					} else {
						img.Set(j, i, image.Black)
					}
				}
			}
			outputFile, err := os.Create(fmt.Sprintf("output%03d.png", k))
			if err != nil {
				panic(err)
			}
			defer outputFile.Close()

			err = png.Encode(outputFile, img)
			if err != nil {
				panic(err)
			}
		}

	}
}

func printRobots(height int, width int, robots []Robot) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			found := false
			num := 0
			for _, robot := range robots {
				if robot.p.x == j && robot.p.y == i {
					num++
					found = true
				}
			}
			if !found {
				fmt.Print(".")
			} else {
				fmt.Print(num)
			}
		}
		fmt.Println()
	}
}
