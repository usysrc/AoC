package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcA(arr []string) {

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

func calcB(arr []string) {
	// print the array
	k := len(arr) - 1
	for {
		if k < 0 {
			break
		}
		// count the number of items at k
		count := 0
		val := arr[k]
		if val == "." {
			k--
			continue
		}
		for n := k; n > 0; n-- {
			if arr[n] != val {
				break
			}
			count++
		}
		// try to find a . space big enought to fit in count
		for n := 0; n < k; n++ {
			if arr[n] != "." {
				continue
			}
			found := true
			for i := 0; i < count && n+i < len(arr); i++ {
				if arr[n+i] != "." {
					found = false
					break
				}
			}
			if found {
				// move the items
				for i := 0; i < count && n+i < len(arr); i++ {
					arr[n+i] = val
					arr[k-i] = "."
				}
				break
			}
		}
		if count == 0 {
			k--
		}
		k -= count
	}
	// print the arr
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
	fmt.Println("Sum Part B: ", sum)
}

func copy(arr []string) []string {
	newArr := make([]string, len(arr))
	for i, v := range arr {
		newArr[i] = v
	}
	return newArr
}

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

	calcA(copy(arr))
	calcB(copy(arr))
}
