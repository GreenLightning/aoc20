package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(solveHomework(lines, parseExpression))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(solveHomework(lines, parseMultiplication))
	}
}

func solveHomework(lines []string, parseExpression func(reader io.RuneScanner) int) (sum int) {
	for _, line := range lines {
		reader := strings.NewReader(line)
		sum += parseExpression(reader)
		if reader.Len() != 0 {
			panic(line)
		}
	}
	return
}

func parseExpression(reader io.RuneScanner) (result int) {
	result = parseValue(reader, parseExpression)
	for {
		eatSpaces(reader)
		operator, _, err := reader.ReadRune()
		if err == io.EOF || (operator != '*' && operator != '+') {
			reader.UnreadRune()
			return
		}

		eatSpaces(reader)
		value := parseValue(reader, parseExpression)

		if operator == '*' {
			result *= value
		} else {
			result += value
		}
	}
}

func parseMultiplication(reader io.RuneScanner) (result int) {
	result = parseAddition(reader)
	for {
		eatSpaces(reader)
		operator, _, err := reader.ReadRune()
		if err == io.EOF || operator != '*' {
			reader.UnreadRune()
			return
		}
		eatSpaces(reader)
		result *= parseAddition(reader)
	}
}

func parseAddition(reader io.RuneScanner) (result int) {
	result = parseValue(reader, parseMultiplication)
	for {
		eatSpaces(reader)
		operator, _, err := reader.ReadRune()
		if err == io.EOF || operator != '+' {
			reader.UnreadRune()
			return
		}
		eatSpaces(reader)
		result += parseValue(reader, parseMultiplication)
	}
}

func parseValue(reader io.RuneScanner, parseExpression func(reader io.RuneScanner) int) int {
	peek, _, _ := reader.ReadRune()
	if peek == '(' {
		value := parseExpression(reader)
		peek, _, _ = reader.ReadRune()
		if peek != ')' {
			panic("missing parenthesis")
		}
		return value
	}
	reader.UnreadRune()
	return parseNumber(reader)
}

func parseNumber(reader io.RuneScanner) int {
	var text strings.Builder
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF || !(char >= '0' && char <= '9') {
			reader.UnreadRune()
			return toInt(text.String())
		}
		text.WriteRune(char)
	}
}

func eatSpaces(reader io.RuneScanner) {
	for {
		rune, _, err := reader.ReadRune()
		if err == io.EOF || rune != ' ' {
			reader.UnreadRune()
			return
		}
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
