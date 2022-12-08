package main

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name        string
	up          *Directory
	directories map[string]*Directory
	files       map[string]*File
	visited     bool
	size        int
}

func readFile() int {
	inputFile, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	root := &Directory{
		name:        "/",
		up:          nil,
		directories: map[string]*Directory{},
		files:       map[string]*File{},
		size:        0,
	}
	currentDirectory := root

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1] == "$" {
			fmt.Println(line)
			if line[2:4] == "cd" {
				path := line[5:]
				if path == ".." {
					if currentDirectory.up == nil {
						fmt.Println(currentDirectory.name, "has no up, is probably root")
					}
					currentDirectory = currentDirectory.up
					fmt.Println(currentDirectory.name)
				} else if path == "/" {

				} else {
					if _, b := currentDirectory.directories[path]; !b {
						fmt.Printf("could not find %s in %s\n", currentDirectory.name, path)
					}
					currentDirectory = currentDirectory.directories[path]
				}
			} else if line[2:4] == "ls" {

			}
		} else if line[0:3] == "dir" {
			dirName := line[4:]
			// fmt.Println("adding folder ", dirName)
			currentDirectory.directories[dirName] = &Directory{
				name:        dirName,
				up:          currentDirectory,
				directories: map[string]*Directory{},
				files:       map[string]*File{},
			}
		} else {
			var size int
			var filename string
			_, err := fmt.Sscanf(line, "%d %s", &size, &filename)
			if err != nil {
				panic(err)
			}
			if currentDirectory.name != "" {
				currentDirectory.files[filename] = &File{
					name: filename,
					size: size,
				}
			} else {
				fmt.Println(currentDirectory.name, line, ". invalid directory.")
			}
		}
	}

	// fmt.Println(v.size)
	sizeOfRoot := calculateFilesizes(root)
	target := 70000000 - sizeOfRoot
	delta := 30000000 - target
	size := findSmallest(delta, root)
	return size
}

func findSmallest(delta int, root *Directory) int {
	min := 1000000000
	var recurse func(currentDirectory *Directory)
	recurse = func(currentDirectory *Directory) {
		for _, dir := range currentDirectory.directories {
			if dir.size >= delta && dir.size < min {
				min = dir.size
				fmt.Println(dir.name)
			}
			recurse(dir)
		}
	}
	recurse(root)
	return min
}

func calculateFilesizes(root *Directory) int {
	count := 0
	var recurse func(currentDirectory *Directory)
	recurse = func(currentDirectory *Directory) {
		if currentDirectory.size == 0 {
			for _, v := range currentDirectory.directories {
				recurse(v)
				currentDirectory.size += v.size
			}
			for _, v := range currentDirectory.files {
				currentDirectory.size += v.size

			}
		}
		fmt.Println(currentDirectory.name, currentDirectory.size)
		count = currentDirectory.size
	}
	recurse(root)
	return count
}

func main() {
	size := readFile()
	fmt.Println("result: ", size)
}
