package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	anyCount := 0
	allCount := 0
	persons := 0
	answers := make(map[string]int)

	countGroup := func() {
		anyCount += len(answers)
		for _, count := range answers {
			if count == persons {
				allCount++
			}
		}
		answers = make(map[string]int)
		persons = 0
	}

	for _, line := range lines {
		if line == "" {
			countGroup()
			continue
		}
		for _, answer := range strings.Split(line, "") {
			answers[answer]++
		}
		persons++
	}
	countGroup()

	fmt.Println("--- Part One ---")
	fmt.Println(anyCount)
	fmt.Println("--- Part Two ---")
	fmt.Println(allCount)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
