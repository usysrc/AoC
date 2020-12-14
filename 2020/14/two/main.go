package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func bit(num int) string {
	return fmt.Sprintf("%036b", num)
}

func dec(num string) int {
	if i, err := strconv.ParseInt(num, 2, 64); err != nil {
		log.Fatal(err)
	} else {
		if i != int64(int(i)) {
			log.Fatal("what")
		}
		return int(i)
	}
	return 0
}

func main() {

	mask := ""
	mem := make(map[int]int)
	for _, line := range readFile() {
		sides := strings.Split(line, " = ")
		left := sides[0]
		right := sides[1]
		if left == "mask" {
			mask = right
			// fmt.Println(nil, mask)
		} else {
			number, err := strconv.Atoi(right)
			if err != nil {
				log.Fatal(err)
			}

			// num := bit(number)
			// for k := 0; k < 36; k++ {
			// 	if mask[k] != 'X' {
			// 		v := num[:k] + string(mask[k]) + num[k+1:]
			// 		num = v
			// 	}
			// }
			// fmt.Println(num)

			r := regexp.MustCompile(`mem\[([[:digit:]]+)\]`)
			target := r.FindStringSubmatch(left)
			if err != nil {
				log.Fatal(err)
			}
			targetnumber, err := strconv.Atoi(target[1])
			if err != nil {
				log.Fatal(err)
			}
			targetbits := bit(targetnumber)
			for k := 0; k < 36; k++ {
				if mask[k] == 'X' {
					v := targetbits[:k] + "X" + targetbits[k+1:]
					targetbits = v
				}
				if mask[k] == '1' {
					v := targetbits[:k] + "1" + targetbits[k+1:]
					targetbits = v
				}
			}
			fmt.Println()
			fmt.Println(bit(targetnumber))
			fmt.Println(targetbits, target[1])
			var check func(k int, s string)
			check = func(k int, s string) {
				if k >= 36 {
					mem[dec(s)] = number
					fmt.Println(s, number, dec(s))
					return
				}
				if targetbits[k] == 'X' {
					check(k+1, s+"0")
					check(k+1, s+"1")
				} else {
					s = s + string(targetbits[k])
					check(k+1, s)
				}
			}
			check(0, "")
		}
	}
	sum := 0
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
}
