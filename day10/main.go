package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// This is a very hacky solution :)
func main() {
	input := readNumbers("input.txt")

	// Add socket.
	input = append(input, 0)

	// Sort input, because an adapter can only increase the number of jolts.
	sort.Ints(input)

	// Add device.
	input = append(input, input[len(input)-1]+3)

	{
		fmt.Println("--- Part One ---")
		d1, d3 := 0, 0
		for i := 0; i+1 < len(input); i++ {
			diff := input[i+1] - input[i]
			if diff == 1 {
				d1++
			} else if diff == 3 {
				d3++
			} else {
				// The code for part two requires the differences to be ones and threes only.
				panic("the elves must have stolen one of your adapters (cannot handle input)")
			}
		}
		fmt.Println(d1 * d3)
	}

	{
		fmt.Println("--- Part Two ---")

		// Count the number of consecutive adapters that have a difference of one on either side.
		// Figure out how many combinations there are for each run.
		// Multiply everything together to get the total number of arrangements.

		// If there are 0 adapters in a run, then there is only 1 combination.
		// If there is  1 adapter  in a run, then you can (optionally) remove it, so there are 2 combinations.
		// If there are 2 adapters in a run, then you can use or remove either one independently, so there are 4 combinations in total.
		// If there are 3 adapters in a run, theoretically, there are 8 combinations,
		// however, if you remove all three adapters, then the joltage difference is 4,
		// which is too much for the adapter, so there are only 7 valid combinations.
		// Longer runs are not handled yet.

		combinations := []int{1, 2, 4, 7}

		arrangements := 1
		run := 0
		for i := 1; i+1 < len(input); i++ {
			if input[i]-input[i-1] == 1 && input[i+1]-input[i] == 1 {
				run++
				continue
			}
			if run >= len(combinations) {
				panic("you have too many adapters (cannot handle input)")
			}
			arrangements *= combinations[run]
			run = 0
		}
		fmt.Println(arrangements)
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
