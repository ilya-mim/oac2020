package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bagRecord struct {
	color string
	count int
}

func readRules(path string) (map[string][]bagRecord, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	rules := make(map[string][]bagRecord)
	reMain := regexp.MustCompile(`(?P<color>[a-z ]+) bags contain (?P<nobags>no other bags.)?`)
	reSub := regexp.MustCompile(`(?P<count>\d+) (?P<color>[a-z ]+) bags?[,|.]`)

	for _, line := range lines {
		matchesMain := reMain.FindStringSubmatch(line)
		matchesSub := reSub.FindAllStringSubmatch(line, -1)

		if matchesMain[2] != "" {
			rules[matchesMain[1]] = nil
		} else {
			rules[matchesMain[1]] = make([]bagRecord, len(matchesSub))
			for i := 0; i < len(matchesSub); i++ {
				count, _ := strconv.Atoi(matchesSub[i][1])
				rules[matchesMain[1]][i] = bagRecord{
					color: matchesSub[i][2],
					count: count,
				}
			}
		}
	}

	return rules, nil
}

// TODO: ugly
func calcOutside(rules map[string][]bagRecord, color string) int {
	count := 0
	colorMap := make(map[string]bool)
	colors := list.New()
	colors.PushBack(color)

	for front := colors.Front(); front != nil; front = front.Next() {
		for key, bags := range rules {
			for _, bag := range bags {
				if bag.color == front.Value {
					if _, ok := colorMap[key]; !ok {
						colorMap[key] = true
						colors.PushBack(key)
						count++
					}
					break
				}
			}
		}
	}

	return count
}

func calcInside(rules map[string][]bagRecord, color string) int {
	count := 0
	colors := list.New()
	colors.PushBack(bagRecord{
		color: color,
		count: 1,
	})

	for front := colors.Front(); front != nil; front = front.Next() {
		for color, bags := range rules {
			if color == front.Value.(bagRecord).color {
				for _, bag := range bags {
					count += bag.count * front.Value.(bagRecord).count
					colors.PushBack(bagRecord{
						color: bag.color,
						count: bag.count * front.Value.(bagRecord).count,
					})
					fmt.Println("Bag color:", bag.color, bag.count*front.Value.(bagRecord).count)
				}
				break
			}
		}
	}

	return count
}

func main() {
	rules, err := readRules("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	count := calcOutside(rules, "shiny gold")
	fmt.Println("Bag count outside:", count)

	count = calcInside(rules, "shiny gold")
	fmt.Println("Bag count inside:", count)
}
