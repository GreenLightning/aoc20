package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	var maxID int
	var taken [1024]bool
	for _, pass := range lines {
		row := binarySearch(pass[0:7], 'F')
		seat := binarySearch(pass[7:10], 'L')
		id := 8*row + seat
		maxID = max(maxID, id)
		taken[id] = true
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(maxID)
	}

	{
		fmt.Println("--- Part Two ---")
		for id := 1; id < len(taken)-1; id++ {
			if taken[id-1] && !taken[id] && taken[id+1] {
				fmt.Println(id)
			}
		}
	}
}

func binarySearch(str string, lower rune) int {
	low, high := 0, (1<<len(str))-1
	for _, c := range str {
		if c == lower {
			high = (low + high) / 2
		} else {
			low = (low+high)/2 + 1
		}
	}
	return low
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

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
