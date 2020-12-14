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

	maskRegex := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRegex := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	{
		fmt.Println("--- Part One ---")

		mem := make(map[uint64]uint64)
		var zeroMask, oneMask uint64

		for _, line := range lines {
			if matches := maskRegex.FindStringSubmatch(line); matches != nil {
				mask := matches[1]
				zeroMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
				oneMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
			} else if matches := memRegex.FindStringSubmatch(line); matches != nil {
				addr, _ := strconv.ParseUint(matches[1], 10, 64)
				value, _ := strconv.ParseUint(matches[2], 10, 64)
				mem[addr] = (value & zeroMask) | oneMask
			} else {
				panic(line)
			}
		}

		var sum uint64
		for _, value := range mem {
			sum += value
		}
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")

		mem := make(map[uint64]uint64)
		var zeroMask, oneMask uint64
		var bits []int

		for _, line := range lines {
			if matches := maskRegex.FindStringSubmatch(line); matches != nil {
				mask := matches[1]
				zeroMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
				oneMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
				bits = nil
				for i, rune := range mask {
					if rune == 'X' {
						bits = append(bits, 35-i)
					}
				}
			} else if matches := memRegex.FindStringSubmatch(line); matches != nil {
				addr, _ := strconv.ParseUint(matches[1], 10, 64)
				value, _ := strconv.ParseUint(matches[2], 10, 64)
				for counter := uint64(0); counter < (1 << len(bits)); counter++ {
					newAddr := (addr & ^zeroMask) | oneMask
					for i, offset := range bits {
						bit := (counter >> i) & 1
						newAddr |= bit << offset
					}
					mem[newAddr] = value
				}
			} else {
				panic(line)
			}
		}

		var sum uint64
		for _, value := range mem {
			sum += value
		}
		fmt.Println(sum)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
