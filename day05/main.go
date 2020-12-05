package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readPasses(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	passes := strings.Split(string(content), "\n")
	return passes, nil
}

func spacePartitioning(code []byte, move byte, start int, end int) int {
	i := 0
	for _, c := range code {
		if c == move {
			end = start + (end-start)/2
			i = end
		} else {
			start = start + (end-start)/2 + 1
			i = start
		}
	}

	return i
}

func processID(code []byte) int {
	// F: 70
	// B: 66
	// L: 76
	// R: 82

	var rowCode []byte = code[0:7]
	row := spacePartitioning(rowCode, 70, 0, 127)

	var columnCode []byte = code[7:10]
	column := spacePartitioning(columnCode, 76, 0, 7)

	return row*8 + column
}

func processIDs(passes []string) []int {
	ids := make([]int, len(passes))
	for i, pass := range passes {
		id := processID([]byte(pass))
		ids[i] = id
	}

	return ids
}

func maxID(ids []int) int {
	max := 0
	for _, id := range ids {
		if id >= max {
			max = id
		}
	}

	return max
}

func findMyID(ids []int) int {
	sort.Ints(ids)
	myID := 0
	for i := 1; i < len(ids); i++ {
		if ids[i] == ids[i-1]+2 {
			myID = ids[i-1] + 1
		}
	}

	return myID
}

func main() {
	passes, err := readPasses("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	ids := processIDs(passes)
	id := maxID(ids)
	fmt.Println("Max id:", id)

	myID := findMyID(ids)
	fmt.Println("My id:", myID)
}
