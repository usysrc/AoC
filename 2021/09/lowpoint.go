// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// const FILE = "input.txt"

// type Map struct {
// 	width  int
// 	height int
// 	field  map[string]int
// }

// func (m *Map) Get(x, y int) (int, bool) {
// 	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
// 	val, ok := m.field[index]
// 	return val, ok
// }

// func (m *Map) Set(x, y, num int) {
// 	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
// 	if x > m.width {
// 		m.width = x + 1
// 	}
// 	if y > m.height {
// 		m.height = y + 1
// 	}
// 	m.field[index] = num
// }

// func (m *Map) GetNeigh(x, y int) []int {
// 	vals := make([]int, 0)
// 	if up, ok := m.Get(x, y-1); ok {
// 		vals = append(vals, up)
// 	}
// 	if down, ok := m.Get(x, y+1); ok {
// 		vals = append(vals, down)
// 	}
// 	if left, ok := m.Get(x-1, y); ok {
// 		vals = append(vals, left)
// 	}
// 	if right, ok := m.Get(x+1, y); ok {
// 		vals = append(vals, right)
// 	}
// 	return vals
// }

// func (m *Map) GetLowPoints() []int {
// 	low := make([]int, 0)
// 	for j := 0; j < m.height+1; j++ {
// 		for i := 0; i < m.width; i++ {
// 			v, ok := m.Get(i, j)
// 			if ok {
// 				lower := true

// 				n := m.GetNeigh(i, j)
// 				for _, p := range n {
// 					if p <= v {
// 						lower = false
// 						break
// 					}
// 				}
// 				if lower {
// 					low = append(low, v)
// 				}
// 			}

// 		}
// 	}
// 	return low
// }

// func (m *Map) GetRiskLevelSum() int {
// 	ps := m.GetLowPoints()
// 	fmt.Println(ps)
// 	sum := 0
// 	for _, p := range ps {
// 		sum += (p + 1)
// 	}
// 	return sum
// }

// func CreateMap() *Map {
// 	m := &Map{}
// 	m.field = make(map[string]int, 0)
// 	return m
// }

// func readFile() *Map {
// 	fileHandle, _ := os.Open(FILE)
// 	defer fileHandle.Close()
// 	fileScanner := bufio.NewScanner(fileHandle)
// 	vents := CreateMap()
// 	j := 0
// 	for fileScanner.Scan() {
// 		line := fileScanner.Text()
// 		for i, char := range line {
// 			number, err := strconv.Atoi(string(char))
// 			if err != nil {
// 				fmt.Errorf("error: %s", err)
// 			}
// 			vents.Set(i, j, number)
// 		}
// 		j++
// 	}
// 	return vents
// }

// func main() {
// 	vents := readFile()
// 	fmt.Println(vents.width, vents.height)
// 	sum := vents.GetRiskLevelSum()
// 	fmt.Printf("sum is %d\n", sum)
// }
