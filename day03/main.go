package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(countTrees(lines, 3, 1))
	}

	{
		fmt.Println("--- Part Two ---")
		dxs := []int{1, 3, 5, 7, 1}
		dys := []int{1, 1, 1, 1, 2}
		result := 1
		for i := range dxs {
			result *= countTrees(lines, dxs[i], dys[i])
		}
		fmt.Println(result)
	}
}

func countTrees(lines []string, dx, dy int) (count int) {
	for x, y := 0, 0; y < len(lines); x, y = x+dx, y+dy {
		if lines[y][x%len(lines[y])] == '#' {
			count++
		}
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
