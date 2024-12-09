package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	arr := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Do something with the line
		for i, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			for k := 0; k < num; k++ {
				if i%2 == 0 {
					arr = append(arr, strconv.Itoa(int(i/2)))
				} else {
					arr = append(arr, ".")
				}

			}
		}
	}
	// now defrag the array
	p := 0
	k := len(arr) - 1
	for {
		if p >= k || p >= len(arr) {
			break
		}
		if arr[p] == "." {
			if arr[k] != "." {
				arr[p] = arr[k]
				arr[k] = "."
			}
			k--
		} else {
			p++
		}
	}

	// Print the result
	for _, v := range arr {
		fmt.Print(v)
	}
	fmt.Println()

	// calculate checksum
	sum := int64(0)
	for i, v := range arr {
		if v != "." {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sum += int64(i * num)
		}
	}
	fmt.Println("Sum Part A: ", sum)
}
