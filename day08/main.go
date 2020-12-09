package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

func readProgram(path string) ([]instruction, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		operation := parts[0]
		argument, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{
			operation: operation,
			argument:  argument,
		}
	}

	return instructions, nil
}

func findLoop(instructions []instruction, start int, accumulator int, fix bool) (bool, map[int]int, int) {
	execCommands := make(map[int]int)
	isLoop := false

	for i := start; ; {
		if i == len(instructions) {
			break
		}
		if _, ok := execCommands[i]; ok {
			isLoop = true
			break
		}

		execCommands[i] = accumulator
		op := instructions[i].operation
		arg := instructions[i].argument

		if i == start && fix {
			if op == "nop" {
				op = "jmp"
			} else if op == "jmp" {
				op = "nop"
			}
		}

		switch op {
		case "nop":
			i++
		case "acc":
			accumulator += arg
			i++
		case "jmp":
			i += arg
		}
	}

	return isLoop, execCommands, accumulator
}

func findCorrupted(instructions []instruction, loop map[int]int) int {
	for key, value := range loop {
		if instructions[key].operation == "acc" {
			continue
		}

		isLoop, _, accumulator := findLoop(instructions, key, value, true)

		if !isLoop {
			return accumulator
		}
	}

	return -1
}

func main() {
	instructions, err := readProgram("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	_, execCommands, accumulator := findLoop(instructions, 0, 0, false)
	fmt.Println("Accumulator:", accumulator)

	accumulator = findCorrupted(instructions, execCommands)
	fmt.Println("Accumulator after the program terminates:", accumulator)
}
