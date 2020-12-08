package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	Kind  string
	Value int
}

func main() {
	var program []Instruction

	lines := readLines("input.txt")
	regex := regexp.MustCompile(`^(acc|jmp|nop) (.\d+)$`)
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		kind, value := matches[1], toInt(matches[2])
		program = append(program, Instruction{kind, value})
	}

	{
		fmt.Println("--- Part One ---")
		acc, _ := run(program)
		fmt.Println(acc)
	}

	{
		fmt.Println("--- Part Two ---")
		for i, inst := range program {
			if inst.Kind == "acc" {
				continue
			}
			if inst.Kind != "jmp" {
				program[i].Kind = "jmp"
			} else {
				program[i].Kind = "nop"
			}
			acc, result := run(program)
			if result == ResultTerminating {
				fmt.Println(acc)
				break
			}
			program[i].Kind = inst.Kind
		}
	}
}

const (
	ResultLooping     = 0
	ResultTerminating = 1
	ResultError       = 2
)

func run(program []Instruction) (acc int, result int) {
	var ip int

	visited := make(map[int]bool)
	for {
		if ip == len(program) {
			return acc, ResultTerminating
		}
		if ip < 0 || ip >= len(program) {
			return acc, ResultError
		}
		if visited[ip] {
			return acc, ResultLooping
		}

		visited[ip] = true
		inst := program[ip]
		switch inst.Kind {
		case "acc":
			acc += inst.Value
			ip++
		case "jmp":
			ip += inst.Value
		case "nop":
			ip++
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
