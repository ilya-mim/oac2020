package main

import (
	"bufio"
	"fmt"
	"os"
)

func readForest(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isTree(trees string, position int) bool {
	treesAsBytes := []byte(trees)
	treeAsByte := []byte("#")[0]
	return treesAsBytes[position] == treeAsByte
}

func traverse(forest []string, right int, down int) int {
	sum := 0
	width := len(forest[0])
	height := len(forest)
	position := 0
	for i := 0; i < height; i += down {
		if isTree(forest[i], position) {
			sum++
		}
		position = (position + right) % width
	}

	return sum
}

func main() {
	forest, err := readForest("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// task #1
	trees := traverse(forest, 3, 1)
	fmt.Println("Encountered trees: ", trees)

	// task #2
	trees21 := traverse(forest, 1, 1)
	fmt.Println("Right 1, down 1. Encountered trees: ", trees21)
	trees22 := trees //traverse(forest, 3, 1)
	fmt.Println("Right 3, down 1. Encountered trees: ", trees22)
	trees23 := traverse(forest, 5, 1)
	fmt.Println("Right 5, down 1. Encountered trees: ", trees23)
	trees24 := traverse(forest, 7, 1)
	fmt.Println("Right 7, down 1. Encountered trees: ", trees24)
	trees25 := traverse(forest, 1, 2)
	fmt.Println("Right 1, down 2. Encountered trees: ", trees25)

	fmt.Println("TREES:", trees21*trees22*trees23*trees24*trees25)
}
