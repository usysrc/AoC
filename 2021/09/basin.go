package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const FILE = "input.txt"

type Map struct {
	width  int
	height int
	field  map[string]int
}

func (m *Map) Get(x, y int) (int, bool) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	val, ok := m.field[index]
	return val, ok
}

func (m *Map) Set(x, y, num int) {
	index := strconv.Itoa(x) + "," + strconv.Itoa(y)
	if x > m.width {
		m.width = x + 1
	}
	if y > m.height {
		m.height = y + 1
	}
	m.field[index] = num
}

func (m *Map) GetNeigh(x, y int) []Point {
	vals := make([]Point, 0)
	if _, ok := m.Get(x, y-1); ok {
		vals = append(vals, CreatePoint(x, y-1))
	}
	if _, ok := m.Get(x, y+1); ok {
		vals = append(vals, CreatePoint(x, y+1))
	}
	if _, ok := m.Get(x-1, y); ok {
		vals = append(vals, CreatePoint(x-1, y))
	}
	if _, ok := m.Get(x+1, y); ok {
		vals = append(vals, CreatePoint(x+1, y))
	}
	return vals
}

type Point struct {
	x int
	y int
}

func CreatePoint(i, j int) Point {
	return Point{x: i, y: j}
}

func (p Point) GetValue(m *Map) int {
	val, ok := m.Get(p.x, p.y)
	if !ok {
		fmt.Errorf("Value not found")
	}
	return val
}

func (p Point) GetIndex() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func (m *Map) GetLowPoints() []Point {
	low := make([]Point, 0)
	for j := 0; j < m.height+1; j++ {
		for i := 0; i < m.width; i++ {
			v, ok := m.Get(i, j)
			if ok {
				lower := true
				n := m.GetNeigh(i, j)
				for _, p := range n {
					if p.GetValue(m) <= v {
						lower = false
						break
					}
				}
				if lower {
					low = append(low, CreatePoint(i, j))
				}
			}

		}
	}
	return low
}

func (m *Map) GetBasin(x, y int) []Point {
	marked := make(map[string]bool, 0)
	stack := make([]Point, 0)
	startPoint := CreatePoint(x, y)
	stack = append(stack, startPoint)
	basin := make([]Point, 0)
	for len(stack) > 0 {
		cur := stack[0]
		basin = append(basin, cur)
		newStack := make([]Point, 0)
		// restore the stack
		for i, v := range stack {
			if i > 0 {
				newStack = append(newStack, v)
			}
		}
		// check out the neighbours
		n := m.GetNeigh(cur.x, cur.y)
		for _, p := range n {
			_, ok := marked[p.GetIndex()]
			if !ok && p.GetValue(m) > cur.GetValue(m) && p.GetValue(m) < 9 {
				newStack = append(newStack, p)
				marked[p.GetIndex()] = true
			}
		}
		stack = newStack
	}
	return basin
}

func (m *Map) GetBasinProduct() int {
	lowpoints := m.GetLowPoints()
	fmt.Println("lowpoints:", lowpoints)
	basins := make([]int, 0)
	for _, p := range lowpoints {
		b := m.GetBasin(p.x, p.y)
		basins = append(basins, len(b))
	}
	sort.Ints(basins)
	fmt.Println(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

// func (m *Map) GetRiskLevelSum() int {
// 	ps := m.GetLowPoints()
// 	fmt.Println(ps)
// 	sum := 0
// 	for _, p := range ps {
// 		sum += (p + 1)
// 	}
// 	return sum
// }

func CreateMap() *Map {
	m := &Map{}
	m.field = make(map[string]int, 0)
	return m
}

func readFile() *Map {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	vents := CreateMap()
	j := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i, char := range line {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Errorf("error: %s", err)
			}
			vents.Set(i, j, number)
		}
		j++
	}
	return vents
}

func main() {
	vents := readFile()
	fmt.Println(vents.width, vents.height)
	product := vents.GetBasinProduct()
	fmt.Printf("product is %d\n", product)
}
