package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type policy func(PwdRecord) bool

type PwdRecord struct {
	min      int
	max      int
	char     string
	password string
}

func readLines(path string) ([]string, error) {
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

func parseEntry(entry string) PwdRecord {
	re := regexp.MustCompile(`(?P<first>\d+)-(?P<second>\d+) (?P<third>\w): (?P<forth>[a-z]+)`)
	matches := re.FindStringSubmatch(entry)

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])

	return PwdRecord{
		min:      min,
		max:      max,
		char:     matches[3],
		password: matches[4],
	}
}

func policyRuleOne(pwdRecord PwdRecord) bool {
	count := strings.Count(pwdRecord.password, pwdRecord.char)

	return count >= pwdRecord.min && count <= pwdRecord.max
}

func policyRuleTwo(pwdRecord PwdRecord) bool {
	chars := []byte(pwdRecord.password)
	chr := []byte(pwdRecord.char)[0]

	return (chars[pwdRecord.min-1] == chr || chars[pwdRecord.max-1] == chr) && chars[pwdRecord.min-1] != chars[pwdRecord.max-1]
}

func count(entries []string, isValid policy) int {
	sum := 0
	for _, v := range entries {
		pwdRecord := parseEntry(v)
		if isValid(pwdRecord) {
			sum++
		}
	}

	return sum
}

func main() {
	entries, err := readLines("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sum := count(entries, policyRuleOne)
	fmt.Println("Valid passwords: ", sum)

	sum = count(entries, policyRuleTwo)
	fmt.Println("Valid passwords: ", sum)
}
