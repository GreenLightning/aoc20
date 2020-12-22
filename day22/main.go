package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	partOne := make([][]int, 2)
	partTwo := make([][]int, 2)

	index := 1
	for player := 0; player < 2; player++ {
		for ; index < len(lines) && lines[index] != ""; index++ {
			value := toInt(lines[index])
			partOne[player] = append(partOne[player], value)
			partTwo[player] = append(partTwo[player], value)
		}
		index += 2
	}

	{
		fmt.Println("--- Part One ---")
		winner := play(partOne, false)
		fmt.Println(score(partOne, winner))
	}

	{
		fmt.Println("--- Part Two ---")
		winner := play(partTwo, true)
		fmt.Println(score(partTwo, winner))
	}
}

func play(cards [][]int, recursive bool) int {
	history := make(map[string]bool)
	for len(cards[0]) != 0 && len(cards[1]) != 0 {
		if recursive {
			var builder strings.Builder
			for player := range cards {
				fmt.Fprintf(&builder, "player%d", player)
				for _, card := range cards[player] {
					fmt.Fprintf(&builder, ",%d", card)
				}
			}
			if ok := history[builder.String()]; ok {
				return 0
			}
			history[builder.String()] = true
		}

		var top [2]int
		top[0], top[1] = cards[0][0], cards[1][0]
		cards[0], cards[1] = cards[0][1:], cards[1][1:]

		winner := 0
		if top[1] > top[0] {
			winner = 1
		}

		if recursive && len(cards[0]) >= top[0] && len(cards[1]) >= top[1] {
			copies := make([][]int, len(cards))
			for player := range copies {
				copies[player] = make([]int, top[player])
				copy(copies[player], cards[player])
			}
			winner = play(copies, recursive)
		}

		cards[winner] = append(cards[winner], top[winner], top[1-winner])
	}

	for player := range cards {
		if len(cards[player]) != 0 {
			return player
		}
	}

	return -1
}

func score(cards [][]int, winner int) (score int) {
	for i, card := range cards[winner] {
		score += card * (len(cards[winner]) - i)
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
