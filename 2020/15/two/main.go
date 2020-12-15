package main

import "fmt"

func main() {
	start := []int{6, 19, 0, 5, 7, 13, 1}
	// start := []int{3, 1, 2}
	m := make(map[int]int, 0)
	for i, num := range start {
		m[num] = i + 1
	}
	num := 0
	for i := len(start) + 1; i < 30000000; i++ {
		// fmt.Println(i, num, m[num])
		if sl, ok := m[num]; ok {
			m[num] = i
			num = i - sl
		} else {
			m[num] = i
			num = 0
		}
	}
	fmt.Println()
	fmt.Println(num)
}
