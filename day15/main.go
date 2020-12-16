package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	for _, text := range strings.Split(readFile("input.txt"), ",") {
		numbers = append(numbers, toInt(text))
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(play(numbers, 2020))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(play(numbers, 30000000))
	}
}

func play(numbers []int, limit int) int {
	var previous int
	memory := make(map[int]int)
	for turn := 1; turn <= limit; turn++ {
		var current int
		if turn <= len(numbers) {
			current = numbers[turn-1]
		} else {
			if old, ok := memory[previous]; ok {
				current = turn - 1 - old
			} else {
				current = 0
			}
		}
		memory[previous] = turn - 1
		previous = current
	}
	return previous
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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
