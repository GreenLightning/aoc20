package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type bus struct{ index, id int }

func main() {
	lines := readLines("input.txt")
	ready := toInt(lines[0])

	var list []bus
	for i, entry := range strings.Split(lines[1], ",") {
		if entry != "x" {
			list = append(list, bus{i, toInt(entry)})
		}
	}

	{
		fmt.Println("--- Part One ---")
		earliestDelay, earliestID := math.MaxInt32, -1
		for _, bus := range list {
			delay := bus.id - ready%bus.id
			if delay < earliestDelay {
				earliestDelay, earliestID = delay, bus.id
			}
		}
		fmt.Println(earliestDelay * earliestID)
	}

	{
		fmt.Println("--- Part Two ---")
		multiplier, offset := 1, 1
		for _, bus := range list {
			for k := 0; ; k++ {
				t := k*multiplier + offset
				if (t+bus.index)%bus.id == 0 {
					multiplier *= bus.id
					offset = t
					break
				}
			}
		}
		fmt.Println(offset)
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
