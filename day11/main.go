package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	lines := readLines("input.txt")
	h, w := len(lines), len(lines[0])

	{
		fmt.Println("--- Part One ---")
		var current, next [][]byte
		for _, line := range lines {
			current = append(current, []byte(line))
			next = append(next, []byte(line))
		}

		for {
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					area := 0 // includes the center seat itself
					for yy := max(y-1, 0); yy <= y+1 && yy < h; yy++ {
						for xx := max(x-1, 0); xx <= x+1 && xx < w; xx++ {
							if current[yy][xx] == '#' {
								area++
							}
						}
					}
					if current[y][x] == 'L' && area == 0 {
						next[y][x] = '#'
					} else if current[y][x] == '#' && area-1 >= 4 {
						next[y][x] = 'L'
					} else {
						next[y][x] = current[y][x]
					}
				}
			}

			current, next = next, current
			if reflect.DeepEqual(current, next) {
				break
			}
		}

		fmt.Println(countOccupiedSeats(current))
	}

	{
		fmt.Println("--- Part Two ---")
		var current, next [][]byte
		for _, line := range lines {
			current = append(current, []byte(line))
			next = append(next, []byte(line))
		}

		for {
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					neighbors := 0
					for _, d := range []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
						for yy, xx := y+d.y, x+d.x; yy >= 0 && yy < h && xx >= 0 && xx < w; yy, xx = yy+d.y, xx+d.x {
							if current[yy][xx] == '#' {
								neighbors++
								break
							}
							if current[yy][xx] == 'L' {
								break
							}
						}
					}
					if current[y][x] == 'L' && neighbors == 0 {
						next[y][x] = '#'
					} else if current[y][x] == '#' && neighbors >= 5 {
						next[y][x] = 'L'
					} else {
						next[y][x] = current[y][x]
					}
				}
			}

			current, next = next, current
			if reflect.DeepEqual(current, next) {
				break
			}
		}

		fmt.Println(countOccupiedSeats(current))
	}
}

func countOccupiedSeats(plan [][]byte) (count int) {
	for y := 0; y < len(plan); y++ {
		for x := 0; x < len(plan[y]); x++ {
			if plan[y][x] == '#' {
				count++
			}
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
