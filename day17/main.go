package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// This solution has been super overengineered to support an arbitrary number
// of dimensions. Unfortunately, the computational complexity means that for 3
// dimensions the code finishes almost instantly, 4 dimensions takes 2
// seconds, 5 dimensions takes 50 seconds, 6 dimensions takes a few hours and
// I did not wait for 7 dimensions to finish.
func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(simulate(lines, [3]int{}))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(simulate(lines, [4]int{}))
	}
}

func simulate(lines []string, prototype interface{}) int {
	// In Go slices cannot be used as map keys, so we have to use arrays. Since
	// the exact type varies depending on the size of the array (number of
	// dimensions) and there are no generics in Go, we have to use the reflect
	// package. Because there are no type literals, you have to create an empty
	// array and pass it as prototype and we get its type using the
	// reflect.TypeOf() function below.
	posType := reflect.TypeOf(prototype)
	dim := reflect.ValueOf(prototype).Len()

	// This map represent the playing field. The key is an array containing
	// the coordinates and the value says whether the cube is active or not.
	active := make(map[interface{}]bool)

	// Copy the input into the active map.
	for y, line := range lines {
		for x, rune := range line {
			if rune == '#' {
				pos := reflect.Indirect(reflect.New(posType))
				pos.Index(0).SetInt(int64(x))
				pos.Index(1).SetInt(int64(y))
				active[pos.Interface()] = true
			}
		}
	}

	// Min and max values (both inclusive) for each dimension.
	// Start with the dimensions of the input.
	min := reflect.Indirect(reflect.New(posType))
	max := reflect.Indirect(reflect.New(posType))
	max.Index(0).SetInt(int64(len(lines[0]) - 1))
	max.Index(1).SetInt(int64(len(lines) - 1))

	for iteration := 0; iteration < 6; iteration++ {
		// Grow region of interest.
		for i := 0; i < dim; i++ {
			min.Index(i).SetInt(min.Index(i).Int() - 1)
			max.Index(i).SetInt(max.Index(i).Int() + 1)
		}

		// Run one cycle.
		next := make(map[interface{}]bool)
		pos := reflect.Indirect(reflect.New(posType))
		update(active, next, min, max, pos, dim, 0)
		active = next
	}

	// Count number of active cells.
	var count int
	for _, a := range active {
		if a {
			count++
		}
	}
	return count
}

func update(active, next map[interface{}]bool, min, max, pos reflect.Value, dim, idx int) {
	if idx == dim {
		// Update cell at pos.
		area := countArea(active, pos, dim, 0)
		if (active[pos.Interface()] && area-1 >= 2 && area-1 <= 3) || (!active[pos.Interface()] && area == 3) {
			next[pos.Interface()] = true
		}
		return
	}

	// Set coordinate at idx and recursively call update with idx+1.
	minValue, maxValue := min.Index(idx).Int(), max.Index(idx).Int()
	for value := minValue; value <= maxValue; value++ {
		pos.Index(idx).SetInt(value)
		update(active, next, min, max, pos, dim, idx+1)
	}
}

func countArea(active map[interface{}]bool, pos reflect.Value, dim, idx int) int {
	if idx == dim {
		// Check neighbor at pos.
		if active[pos.Interface()] {
			return 1
		}
		return 0
	}

	// Update coordinate at idx and recursively call countArea with idx+1.
	count := 0
	value := pos.Index(idx).Int()
	for delta := -1; delta <= 1; delta++ {
		pos.Index(idx).SetInt(value + int64(delta))
		count += countArea(active, pos, dim, idx+1)
	}
	pos.Index(idx).SetInt(value)
	return count
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
