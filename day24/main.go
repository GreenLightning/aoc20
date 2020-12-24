package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	east      int
	southeast int
}

var dirs = []Position{Position{-1, 0}, Position{0, -1}, Position{1, -1}, Position{1, 0}, Position{0, 1}, Position{-1, 1}}

func main() {
	lines := readLines("input.txt")

	floor := make(map[Position]bool)
	for _, line := range lines {
		var pos Position
		for line != "" {
			if strings.HasPrefix(line, "w") {
				line = strings.TrimPrefix(line, "w")
				pos.east--
			} else if strings.HasPrefix(line, "nw") {
				line = strings.TrimPrefix(line, "nw")
				pos.southeast--
			} else if strings.HasPrefix(line, "ne") {
				line = strings.TrimPrefix(line, "ne")
				pos.southeast--
				pos.east++
			} else if strings.HasPrefix(line, "e") {
				line = strings.TrimPrefix(line, "e")
				pos.east++
			} else if strings.HasPrefix(line, "se") {
				line = strings.TrimPrefix(line, "se")
				pos.southeast++
			} else if strings.HasPrefix(line, "sw") {
				line = strings.TrimPrefix(line, "sw")
				pos.southeast++
				pos.east--
			} else {
				panic(line)
			}
		}
		floor[pos] = !floor[pos]
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(count(floor))
	}

	for i := 0; i < 100; i++ {
		next := make(map[Position]bool)
		update := func(pos Position) {
			var neighbors int
			for _, dir := range dirs {
				if floor[Position{pos.east + dir.east, pos.southeast + dir.southeast}] {
					neighbors++
				}
			}
			if floor[pos] && neighbors >= 1 && neighbors <= 2 {
				next[pos] = true
			}
			if !floor[pos] && neighbors == 2 {
				next[pos] = true
			}
		}
		for pos := range floor {
			update(pos)
			for _, dir := range dirs {
				update(Position{pos.east + dir.east, pos.southeast + dir.southeast})
			}
		}
		floor = next
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(count(floor))
	}
}

func count(floor map[Position]bool) (count int) {
	for _, black := range floor {
		if black {
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
