package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseEntries(path string) ([]int, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	entries := strings.Fields(string(content))
	output := make([]int, len(entries))

	for i, v := range entries {
		entry, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		output[i] = entry
	}

	return output, nil
}

func sumOfTwo(entries []int, sum int) (bool, int, int) {
	sort.Ints(entries)

	for i, j := 0, len(entries)-1; i < j; {
		if entries[i]+entries[j] == sum {
			return true, entries[i], entries[j]
		} else if entries[i]+entries[j] < sum {
			i++
		} else {
			j--
		}
	}
	return false, 0, 0
}

func sumOfThree(entries []int, sum int) (bool, int, int, int) {
	for i := 0; i < len(entries)-2; i++ {
		for j := i + 1; j < len(entries)-1; j++ {
			for l := j + 1; l < len(entries); l++ {
				if entries[i]+entries[j]+entries[l] == sum {
					return true, entries[i], entries[j], entries[l]
				}
			}
		}
	}
	return false, 0, 0, 0
}

func main() {
	entries, err := parseEntries("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	isFound, v1, v2 := sumOfTwo(entries, 2020)
	if isFound {
		fmt.Println(v1, "*", v2, "=", v1*v2)
	} else {
		fmt.Println("No two matching entries found!")
	}

	isFound, v1, v2, v3 := sumOfThree(entries, 2020)
	if isFound {
		fmt.Println(v1, "*", v2, "*", v3, "=", v1*v2*v3)
	} else {
		fmt.Println("No three matching entries found!")
	}
}
