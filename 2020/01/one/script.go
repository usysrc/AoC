package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func readStdin() []int {
	var arr = []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var i, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println(err)
		}
		arr = append(arr, i)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return arr
}

func main() {
	arr := readStdin();

	var m = make(map[int]int)
	for _, num := range arr {
		if val, ok := m[2020 - num]; ok {
			fmt.Println(num * val);
		}
		m[num] = num;
	}
}

