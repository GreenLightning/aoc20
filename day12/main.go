package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DirNorth, DirWest, DirSouth, DirEast = 0, 1, 2, 3

var table = []struct{ x, y int }{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

type Instruction struct {
	action string
	value  int
}

func main() {
	var instructions []Instruction

	{
		lines := readLines("input.txt")
		for _, line := range lines {
			instructions = append(instructions, Instruction{line[:1], toInt(line[1:])})
		}
	}

	{
		fmt.Println("--- Part One ---")
		x, y := 0, 0
		dir := DirEast
		for _, inst := range instructions {
			moveDir, moveDist := dir, 0
			switch inst.action {
			case "N":
				moveDir, moveDist = DirNorth, inst.value
			case "S":
				moveDir, moveDist = DirSouth, inst.value
			case "E":
				moveDir, moveDist = DirEast, inst.value
			case "W":
				moveDir, moveDist = DirWest, inst.value
			case "L":
				dir = (dir + inst.value/90) % 4
			case "R":
				dir = (dir + 4 - inst.value/90) % 4
			case "F":
				moveDir, moveDist = dir, inst.value
			}
			x += table[moveDir].x * moveDist
			y += table[moveDir].y * moveDist
		}
		fmt.Println(abs(x) + abs(y))
	}

	{
		fmt.Println("--- Part Two ---")
		x, y := 0, 0
		wx, wy := 10, 1 // waypoint
		for _, inst := range instructions {
			switch inst.action {
			case "N":
				wx += table[DirNorth].x * inst.value
				wy += table[DirNorth].y * inst.value
			case "S":
				wx += table[DirSouth].x * inst.value
				wy += table[DirSouth].y * inst.value
			case "E":
				wx += table[DirEast].x * inst.value
				wy += table[DirEast].y * inst.value
			case "W":
				wx += table[DirWest].x * inst.value
				wy += table[DirWest].y * inst.value
			case "L":
				dirX := (DirEast + inst.value/90) % 4
				dirY := (dirX + 1) % 4
				nx := table[dirX].x*wx + table[dirY].x*wy
				ny := table[dirX].y*wx + table[dirY].y*wy
				wx, wy = nx, ny
			case "R":
				dirX := (DirEast + 4 - inst.value/90) % 4
				dirY := (dirX + 1) % 4
				nx := table[dirX].x*wx + table[dirY].x*wy
				ny := table[dirX].y*wx + table[dirY].y*wy
				wx, wy = nx, ny
			case "F":
				x += wx * inst.value
				y += wy * inst.value
			}
		}
		fmt.Println(abs(x) + abs(y))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
