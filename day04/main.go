package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func processPasports(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	data := string(content)
	passports := strings.Split(data, "\n\n")

	return passports, nil
}

func validateStrict(passports []string) int {
	requiredFileds := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`byr:(?P<byr>\d{4})\b`),
		"iyr": regexp.MustCompile(`iyr:(?P<iyr>\d{4})\b`),
		"eyr": regexp.MustCompile(`eyr:(?P<eyr>\d{4})\b`),
		"hgt": regexp.MustCompile(`hgt:(?P<hgt>\d+)(?P<unit>cm|in)\b`),
		"hcl": regexp.MustCompile(`hcl:#[a-f0-9]{6}\b`),
		"ecl": regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)\b`),
		"pid": regexp.MustCompile(`pid:\d{9}\b`),
	}

	passportCount := 0
	for _, passport := range passports {
		fieldCount := 0

		if hcl := requiredFileds["hcl"].FindString(passport); hcl != "" {
			fieldCount++
		}

		if ecl := requiredFileds["ecl"].FindString(passport); ecl != "" {
			fieldCount++
		}

		if byr := requiredFileds["byr"].FindStringSubmatch(passport); byr != nil {
			year, _ := strconv.Atoi(byr[1])
			if year >= 1920 && year <= 2002 {
				fieldCount++
			}
		}

		if iyr := requiredFileds["iyr"].FindStringSubmatch(passport); iyr != nil {
			year, _ := strconv.Atoi(iyr[1])
			if year >= 2010 && year <= 2020 {
				fieldCount++
			}
		}

		if eyr := requiredFileds["eyr"].FindStringSubmatch(passport); eyr != nil {
			year, _ := strconv.Atoi(eyr[1])
			if year >= 2020 && year <= 2030 {
				fieldCount++
			}
		}

		if hgt := requiredFileds["hgt"].FindStringSubmatch(passport); hgt != nil {
			height, _ := strconv.Atoi(hgt[1])
			unit := hgt[2]
			if unit == "cm" && height >= 150 && height <= 193 {
				fieldCount++
			} else if unit == "in" && height >= 59 && height <= 76 {
				fieldCount++
			}
		}

		if requiredFileds["pid"].FindString(passport) != "" {
			fieldCount++
		}

		if fieldCount == 7 {
			passportCount++
		}
	}

	return passportCount
}

func validate(passports []string) int {
	requiredFileds := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`byr:`),
		"iyr": regexp.MustCompile(`iyr:`),
		"eyr": regexp.MustCompile(`eyr:`),
		"hgt": regexp.MustCompile(`hgt:`),
		"hcl": regexp.MustCompile(`hcl:`),
		"ecl": regexp.MustCompile(`ecl:`),
		"pid": regexp.MustCompile(`pid:`),
	}

	passportCount := 0
	for _, passport := range passports {
		fieldCount := 0

		for _, re := range requiredFileds {
			value := re.FindString(passport)
			if value != "" {
				fieldCount++
			}
		}

		if fieldCount == 7 {
			passportCount++
		}
	}

	return passportCount
}

func main() {
	passports, err := processPasports("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	count := validate(passports)
	fmt.Println("#1 Valid passports:", count)

	count = validateStrict(passports)
	fmt.Println("#2 Valid passports:", count)
}
