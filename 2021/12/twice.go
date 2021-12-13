package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE = "test.txt"

type Node struct {
	name    string
	isSmall bool
	visited int
}

func (n Node) Compare(o Node) bool {
	return n.name == o.name
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

var nodePool = map[string]*Node{}

func readFile() (map[*Node][]*Node, *Node) {
	fileHandle, _ := os.Open(FILE)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	g := map[*Node][]*Node{}
	var start *Node
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nodes := strings.Split(line, "-")

		left := &Node{name: nodes[0], isSmall: isLower(nodes[0])}
		if val, ok := nodePool[nodes[0]]; ok {
			left = val
		}
		nodePool[nodes[0]] = left

		right := &Node{name: nodes[1], isSmall: isLower(nodes[1]), visited: 0}
		if val, ok := nodePool[nodes[1]]; ok {
			right = val
		}
		nodePool[nodes[1]] = right
		if left.name == "start" {
			start = left
		}

		if val, ok := g[left]; ok {
			g[left] = append(val, right)
		} else {
			g[left] = []*Node{right}
		}
		if val, ok := g[right]; ok {
			g[right] = append(val, left)
		} else {
			g[right] = []*Node{left}
		}

	}
	return g, start
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverse(n *Node, graph map[*Node][]*Node, visited int) int {
	if n.name == "end" {
		return 1
	}
	p := 0
	targets := graph[n]
	for _, v := range targets {
		if v.visited < visited && v.name != "start" {
			if v.isSmall {
				v.visited++
				if v.visited >= 2 {
					visited = 1
				}
			}
			p += traverse(v, graph, visited)
			if v.isSmall {
				v.visited--
				if v.visited < 2 {
					visited = 2
				}
			}
		}
	}
	return p
}

func main() {
	graph, start := readFile()
	paths := traverse(start, graph, 2)
	fmt.Println(paths)
}
