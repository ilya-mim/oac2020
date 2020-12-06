package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readGroups(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(content), "\n\n")
	return groups, nil
}

func count(people []string, all bool) int {
	uniqueAnswers := make(map[byte]int)
	peopleCount := len(people)
	for _, answers := range people {
		az := []byte(answers)
		for _, v := range az {
			uniqueAnswers[v]++
		}
	}

	count := 0
	for _, v := range uniqueAnswers {
		if all && v == peopleCount {
			count++
		} else if all == false {
			count++
		}
	}

	return count
}

func calcSum(groups []string, all bool) int {
	sum := 0
	for _, group := range groups {
		sum += count(strings.Split(group, "\n"), all)
	}

	return sum
}

func main() {
	groups, err := readGroups("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sum := calcSum(groups, false)
	fmt.Println("Sum of anyone counts:", sum)

	sum = calcSum(groups, true)
	fmt.Println("Sum of everyone counts:", sum)
}
