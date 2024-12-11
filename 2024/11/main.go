package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Day 10: Plutonian Pebbles")
	fmt.Println("Using file: ", os.Args[1])

	stonerow := ""
	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stonerow = line
	}

	// Part 1 & 2
	type CacheKey struct {
		stone string
		i     int
	}
	cache := map[CacheKey]int64{}
	var countStones func(string, int, int) int64
	countStones = func(stone string, i int, max int) int64 {
		if num, found := cache[CacheKey{stone, i}]; found {
			return num
		}
		count := int64(0)
		if i == max {
			return int64(1)
		}
		if stone == "0" {
			count += countStones("1", i+1, max)
		} else if len(stone)%2 == 0 && len(stone) > 1 {
			left := stone[:len(stone)/2]
			right := stone[len(stone)/2:]
			for left[0] == '0' && len(left) > 1 {
				left = left[1:]
			}
			for right[0] == '0' && len(right) > 1 {
				right = right[1:]
			}
			count += countStones(left, i+1, max)
			count += countStones(right, i+1, max)
		} else {
			num, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			count += countStones(strconv.Itoa(num*2024), i+1, max)
		}
		cache[CacheKey{stone, i}] = count
		return count
	}
	// Part 1
	stonesA := int64(0)
	line := strings.Split(stonerow, " ")
	for i := 0; i < len(line); i++ {
		stonesA += countStones(line[i], 0, 25)
	}
	fmt.Println("Part 1, number of stones: ", stonesA)

	// Part 2
	cache = map[CacheKey]int64{}
	stonesB := int64(0)
	for i := 0; i < len(line); i++ {
		stonesB += countStones(line[i], 0, 75)
	}
	fmt.Println("Part 2, number of stones: ", stonesB)
}
