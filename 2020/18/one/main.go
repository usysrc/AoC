package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func mul(a int, b int) int {
	return a * b
}
func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

func eval(k int, line string, stop int) int {
	var op func(a int, b int) int
	res := 0
	for {
		tok := line[k]
		if tok == ' ' {

		} else if tok == '*' {
			op = mul
		} else if tok == '+' {
			op = add
		} else if tok == '-' {
			op = sub
		} else if tok != ')' {
			b := int(line[k] - '0')

			if line[k] == '(' {
				str := ""
				x := 1
				c := 0
				for {
					str = str + string(line[k+x])
					if line[k+x] == '(' {
						c++
					} else if line[k+x] == ')' {
						c--
						if c < 0 {
							break
						}
					}
					x++
				}
				b = eval(0, str, len(str))
				k += len(str) + 1
			}
			if op != nil {
				fmt.Println(res, b)
				res = op(res, b)
			} else {
				res = b
			}
		}
		k++
		if k >= stop {
			break
		}

	}
	return res
}

func main() {
	sum := 0
	for _, line := range readFile() {
		res := eval(0, line, len(line))
		sum += res
	}
	fmt.Println(sum)
}
