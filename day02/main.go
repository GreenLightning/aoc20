package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	regex := regexp.MustCompile(`^(\d+)-(\d+) (.): (.+)$`)

	{
		fmt.Println("--- Part One ---")
		valid := 0
		for _, line := range lines {
			matches := regex.FindStringSubmatch(line)
			min, max := toInt(matches[1]), toInt(matches[2])
			letter, password := matches[3], matches[4]
			count := strings.Count(password, letter)
			if count >= min && count <= max {
				valid++
			}
		}
		fmt.Println(valid)
	}

	{
		fmt.Println("--- Part Two ---")
		valid := 0
		for _, line := range lines {
			matches := regex.FindStringSubmatch(line)
			a, b := toInt(matches[1])-1, toInt(matches[2])-1
			letter, password := matches[3], matches[4]
			if (password[a:a+1] == letter) != (password[b:b+1] == letter) {
				valid++
			}
		}
		fmt.Println(valid)
	}
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
