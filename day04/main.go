package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

var (
	yearRegex = regexp.MustCompile(`^\d{4}$`)
	hairRegex = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	eyeRegex  = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegex  = regexp.MustCompile(`^\d{9}$`)
)

func main() {
	var passports []Passport

	{
		lines := readLines("input.txt")
		current := make(Passport)
		for _, line := range lines {
			if line == "" {
				passports = append(passports, current)
				current = make(Passport)
				continue
			}
			for _, field := range strings.Split(line, " ") {
				parts := strings.Split(field, ":")
				current[parts[0]] = parts[1]
			}
		}
		if len(current) != 0 {
			passports = append(passports, current)
		}
	}

	{
		fmt.Println("--- Part One ---")
		count := 0
		required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		for _, passport := range passports {
			valid := true
			for _, field := range required {
				if _, ok := passport[field]; !ok {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		count := 0
		for _, passport := range passports {
			if validateYear(passport, "byr", 1920, 2002) &&
				validateYear(passport, "iyr", 2010, 2020) &&
				validateYear(passport, "eyr", 2020, 2030) &&
				validateHeight(passport) &&
				hairRegex.MatchString(passport["hcl"]) &&
				eyeRegex.MatchString(passport["ecl"]) &&
				pidRegex.MatchString(passport["pid"]) {
				count++
			}
		}
		fmt.Println(count)
	}
}

func validateYear(passport Passport, name string, min, max int) bool {
	field := passport[name]
	if !yearRegex.MatchString(field) {
		return false
	}
	value := toInt(field)
	return (value >= min && value <= max)
}

func validateHeight(passport Passport) bool {
	field := passport["hgt"]
	if strings.HasSuffix(field, "cm") {
		value := toInt(strings.TrimSuffix(field, "cm"))
		return (value >= 150 && value <= 193)
	}
	if strings.HasSuffix(field, "in") {
		value := toInt(strings.TrimSuffix(field, "in"))
		return (value >= 59 && value <= 76)
	}
	return false
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
