package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct{ low, high int }

func (r *Range) Contains(value int) bool {
	return r.low <= value && value <= r.high
}

type Field []Range

func (f *Field) Accepts(value int) bool {
	for _, r := range *f {
		if r.Contains(value) {
			return true
		}
	}
	return false
}

type Ticket []int

func parseTicket(line string) (ticket Ticket) {
	for _, part := range strings.Split(line, ",") {
		ticket = append(ticket, toInt(part))
	}
	return
}

func main() {
	lines := readLines("input.txt")

	fields := make(map[string]Field)
	var theTicket Ticket
	var nearbyTickets []Ticket

	// Parse input.
	{
		i := 0

		// Parse field descriptions.
		fieldRegex := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
		for ; i < len(lines); i++ {
			matches := fieldRegex.FindStringSubmatch(lines[i])
			if matches == nil {
				break
			}
			name := matches[1]
			one := Range{toInt(matches[2]), toInt(matches[3])}
			two := Range{toInt(matches[4]), toInt(matches[5])}
			fields[name] = Field{one, two}
		}

		// Skip empty lines.
		for i < len(lines) && lines[i] == "" {
			i++
		}

		// Skip text.
		if lines[i] != "your ticket:" {
			panic(lines[i])
		}
		i++

		// Parse the ticket.
		theTicket = parseTicket(lines[i])
		i++

		// Skip empty lines.
		for i < len(lines) && lines[i] == "" {
			i++
		}

		// Skip text.
		if lines[i] != "nearby tickets:" {
			panic(lines[i])
		}
		i++

		// Parse nearby tickets.
		for ; i < len(lines); i++ {
			nearbyTickets = append(nearbyTickets, parseTicket(lines[i]))
		}
	}

	{
		fmt.Println("--- Part One ---")
		var errorRate int
		for ticketIndex := 0; ticketIndex < len(nearbyTickets); ticketIndex++ {
			ticket := nearbyTickets[ticketIndex]
			ticketOk := true
			for _, value := range ticket {
				valueOk := false
				for _, field := range fields {
					if field.Accepts(value) {
						valueOk = true
						break
					}
				}
				if !valueOk {
					errorRate += value
					ticketOk = false
				}
			}
			if !ticketOk {
				// Remove invalid ticket for part two.
				copy(nearbyTickets[ticketIndex:], nearbyTickets[ticketIndex+1:])
				nearbyTickets = nearbyTickets[:len(nearbyTickets)-1]
				ticketIndex--
			}
		}
		fmt.Println(errorRate)
	}

	{
		fmt.Println("--- Part Two ---")

		// Map from field name to field index.
		fieldIndices := make(map[string]int)

		// Array of indices that we have not assigned to a field yet.
		var openIndices []int
		for index := range theTicket {
			openIndices = append(openIndices, index)
		}

		if len(openIndices) != len(fields) {
			panic("number of values different from number of fields")
		}

		for len(openIndices) != 0 {
			progress := false

			for name, field := range fields {
				if _, ok := fieldIndices[name]; ok {
					continue
				}

				var possibleIndices []int
				for _, index := range openIndices {
					indexOk := true
					for _, ticket := range nearbyTickets {
						if !field.Accepts(ticket[index]) {
							indexOk = false
							break
						}
					}
					if indexOk {
						possibleIndices = append(possibleIndices, index)
					}
				}

				if len(possibleIndices) == 1 {
					fieldIndices[name] = possibleIndices[0]
					for i, index := range openIndices {
						if index == possibleIndices[0] {
							copy(openIndices[i:], openIndices[i+1:])
							openIndices = openIndices[:len(openIndices)-1]
							break
						}
					}
					progress = true
				}
			}

			if !progress {
				panic("no progress")
			}
		}

		var departureProduct int = 1
		for name, index := range fieldIndices {
			if strings.HasPrefix(name, "departure") {
				departureProduct *= theTicket[index]
			}
		}

		fmt.Println(departureProduct)
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
