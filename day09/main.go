package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readNumbers("input.txt")

	var target int

	{
		fmt.Println("--- Part One ---")
		for i := 25; i < len(input); i++ {
			window := input[i-25 : i]
			number := input[i]

			ok := false
			for j := 0; !ok && j < 25; j++ {
				for k := j + 1; !ok && k < 25; k++ {
					if number == window[j]+window[k] {
						ok = true
					}
				}
			}

			if !ok {
				target = number
				break
			}
		}
		fmt.Println(target)
	}

	{
		fmt.Println("--- Part Two ---")
		for length := 2; length <= len(input); length++ {
			for start := 0; start+length <= len(input); start++ {
				span := input[start : start+length]

				sum := 0
				for i := 0; i < length; i++ {
					sum += span[i]
				}

				if sum == target {
					min, max := span[0], span[0]
					for i := 1; i < length; i++ {
						if value := span[i]; value < min {
							min = value
						} else if value > max {
							max = value
						}
					}
					fmt.Println(min + max)
				}
			}
		}
	}
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
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
