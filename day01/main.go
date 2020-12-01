package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	entries := readNumbers("input.txt")

	{
		fmt.Println("--- Part One ---")
	one:
		for _, a := range entries {
			for _, b := range entries {
				if a+b == 2020 {
					fmt.Println(a * b)
					break one
				}
			}
		}
	}

	{
		fmt.Println("--- Part Two ---")
	two:
		for _, a := range entries {
			for _, b := range entries {
				for _, c := range entries {
					if a+b+c == 2020 {
						fmt.Println(a * b * c)
						break two
					}
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
