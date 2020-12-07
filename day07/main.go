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

	mainRegex := regexp.MustCompile(`^(.*) bags contain (.*)\.$`)
	subRegex := regexp.MustCompile(`^(\d+) (.*) bags?$`)

	rules := make(map[string]map[string]int)
	for _, line := range lines {
		mainMatches := mainRegex.FindStringSubmatch(line)
		mainColor := mainMatches[1]
		rule := make(map[string]int)
		if mainMatches[2] != "no other bags" {
			for _, part := range strings.Split(mainMatches[2], ", ") {
				subMatches := subRegex.FindStringSubmatch(part)
				count := toInt(subMatches[1])
				color := subMatches[2]
				rule[color] = count
			}
		}
		rules[mainColor] = rule
	}

	{
		fmt.Println("--- Part One ---")
		count := 0
		for color, _ := range rules {
			if color != "shiny gold" && canCarry(rules, color) {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(countChildren(rules, "shiny gold"))
	}
}

func canCarry(rules map[string]map[string]int, color string) bool {
	if color == "shiny gold" {
		return true
	}
	for containedColor, _ := range rules[color] {
		if canCarry(rules, containedColor) {
			return true
		}
	}
	return false
}

func countChildren(rules map[string]map[string]int, color string) (count int) {
	for containedColor, multiplier := range rules[color] {
		count += multiplier * (1 + countChildren(rules, containedColor))
	}
	return
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
