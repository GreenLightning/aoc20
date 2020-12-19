package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Symbol struct {
	rule int
	text string
}

type Run []Symbol
type Rule []Run

func main() {
	lines := readLines("input.txt")

	rules := make(map[int]Rule)

	idRegex := regexp.MustCompile(`^(\d+):`)
	literalRegex := regexp.MustCompile(`^"(.)"`)
	ruleRegex := regexp.MustCompile(`^(\d+)`)

	index := 0
	for ; index < len(lines); index++ {
		line := lines[index]
		if line == "" {
			break
		}

		idMatches := idRegex.FindStringSubmatch(line)
		line = strings.TrimSpace(strings.TrimPrefix(line, idMatches[0]))
		id := toInt(idMatches[1])

		var rule Rule
		var run Run
		for line != "" {
			if strings.HasPrefix(line, "|") {
				line = strings.TrimSpace(strings.TrimPrefix(line, "|"))
				rule = append(rule, run)
				run = nil
			}
			if matches := literalRegex.FindStringSubmatch(line); matches != nil {
				line = strings.TrimSpace(strings.TrimPrefix(line, matches[0]))
				run = append(run, Symbol{rule: -1, text: matches[1]})
			} else if matches := ruleRegex.FindStringSubmatch(line); matches != nil {
				line = strings.TrimSpace(strings.TrimPrefix(line, matches[0]))
				run = append(run, Symbol{rule: toInt(matches[1])})
			} else {
				panic(line)
			}
		}
		rules[id] = append(rule, run)
	}

	{
		fmt.Println("--- Part One ---")
		var count int
		for _, line := range lines[index:] {
			ok, text := matches(rules, line, 0)
			if ok && text == "" {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")

		// This assumes that rule 0 is "0: 8 11" and all other rules do not
		// reference rules 0, 8 and 11 (and therefore still do not contain any
		// loops). Together with the new rules 8 and 11 from part 2 we have:
		//
		// 0:  8 11
		// 8:  42 | 42 8
		// 11: 42 31 | 42 11 31
		//
		// which can be written as 42^x 42^y 31^y ; x >= 1, y >= 1
		// or equivalently: 42^a 31^b ; a >= 2, b >= 1, a > b

		var count int
		for _, line := range lines[index:] {
			text, a, b := line, -1, -1
			for ok := true; ok; a++ {
				ok, text = matches(rules, text, 42)
			}
			for ok := true; ok; b++ {
				ok, text = matches(rules, text, 31)
			}
			if a >= 2 && b >= 1 && a > b && text == "" {
				count++
			}
		}
		fmt.Println(count)
	}
}

func matches(rules map[int]Rule, input string, id int) (bool, string) {
	for _, run := range rules[id] {
		ok, text := true, input
		for _, symbol := range run {
			if symbol.text != "" {
				if strings.HasPrefix(text, symbol.text) {
					text = strings.TrimPrefix(text, symbol.text)
				} else {
					ok = false
				}
			} else {
				ok, text = matches(rules, text, symbol.rule)
			}
			if !ok {
				break
			}
		}
		if ok {
			return true, text
		}
	}
	return false, input
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
