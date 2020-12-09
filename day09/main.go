package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readData(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		data = append(data, number)
	}
	return data, scanner.Err()
}

func isValid(number int, preamble []int) bool {
	tmp := make([]int, len(preamble))
	copy(tmp, preamble)
	sort.Ints(tmp)

	for i, j := 0, len(tmp)-1; i < j; {
		if tmp[i]+tmp[j] == number {
			return true
		} else if tmp[i]+tmp[j] < number {
			i++
		} else {
			j--
		}
	}
	return false
}

func findNumber(data []int, preamble int) (int, bool) {
	for i := 0; i < len(data)-preamble; i++ {
		number := data[i+preamble]
		if !isValid(number, data[i:i+preamble]) {
			return number, true
		}
	}

	return 0, false
}

func findRange(number int, data []int) (int, int, bool) {
	j := 0
	i := 1
	sum := data[j] + data[i]
	for {
		if sum == number && j != i {
			return j, i, true
		} else if sum < number || j == i {
			i++
			if i == len(data) {
				return 0, 0, false
			}
			sum += data[i]
		} else {
			sum -= data[j]
			j++
		}
	}
}

func findWeakness(number int, data []int) (int, bool) {
	j, i, ok := findRange(number, data)

	if ok {
		tmp := make([]int, i+1-j)
		copy(tmp, data[j:i+1])
		sort.Ints(tmp)
		return tmp[0] + tmp[i-j], true
	}

	return 0, false
}

func main() {
	data, err := readData("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	preamble := 25
	number, _ := findNumber(data, preamble)
	fmt.Println("Number:", number)

	weakness, _ := findWeakness(number, data)
	fmt.Println("Weakness:", weakness)
}
