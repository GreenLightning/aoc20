package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	publicKeys := readNumbers("input.txt")
	loopSizes := make([]int, len(publicKeys))
	for i, key := range publicKeys {
		loopSizes[i] = calculateLoopSize(key)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(transform(publicKeys[0], loopSizes[1]))
		fmt.Println(transform(publicKeys[1], loopSizes[0]))
	}
}

func calculateLoopSize(key int) int {
	iteration := 0
	for value := 7; value != key; iteration++ {
		value = (value * 7) % 20201227
	}
	return iteration
}

func transform(subjectNumber int, loopSize int) int {
	value := subjectNumber
	for iteration := 0; iteration < loopSize; iteration++ {
		value = (value * subjectNumber) % 20201227
	}
	return value
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
