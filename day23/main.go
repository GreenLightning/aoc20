package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cup struct {
	value int32
	next  int32
}

func main() {
	input := readFile("input.txt")

	{
		fmt.Println("--- Part One ---")

		var cups []Cup
		for _, text := range strings.Split(input, "") {
			value := toInt(text)
			cups = append(cups, Cup{int32(value), int32(len(cups) + 1)})
		}
		cups[len(cups)-1].next = 0

		play(cups, 100)

		for i, cup := range cups {
			if cup.value == 1 {
				for j := cups[i].next; j != int32(i); j = cups[j].next {
					fmt.Printf("%d", cups[j].value)
				}
				fmt.Println()
				break
			}
		}
	}

	{
		fmt.Println("--- Part Two ---")

		var cups []Cup
		for _, text := range strings.Split(input, "") {
			value := toInt(text)
			cups = append(cups, Cup{int32(value), int32(len(cups) + 1)})
		}
		for len(cups) < 1_000_000 {
			cups = append(cups, Cup{int32(len(cups) + 1), int32(len(cups) + 1)})
		}
		cups[len(cups)-1].next = 0

		play(cups, 10_000_000)

		for i, cup := range cups {
			if cup.value == 1 {
				a := cups[i].next
				b := cups[a].next
				fmt.Println(int(cups[a].value) * int(cups[b].value))
				break
			}
		}
	}
}

func play(cups []Cup, moves int) {
	var current int32 = 0
	for move := 0; move < moves; move++ {
		// Pick up the 3 cups after the current cup.
		pickup := cups[current].next
		after := pickup
		for i := 0; i < 3; i++ {
			after = cups[after].next
		}
		cups[current].next = after

		// Find the value of the destination cup.
		currentValue := cups[current].value
		var destinationValue int32
		for offset := int32(1); ; offset++ {
			destinationValue = currentValue - offset
			if destinationValue < 1 { // wrap around
				destinationValue += int32(len(cups))
			}
			// Check if this value belongs to one of the cups we picked up.
			test := pickup
			found := false
			for i := 0; i < 3; i++ {
				if cups[test].value == destinationValue {
					found = true
					break
				}
				test = cups[test].next
			}
			if !found {
				break
			}
		}

		// Find the index of the destination cup.
		// For most of the cups we can directly compute the index as they were added to the array in order.
		// For the other cups the search will be short as they are all at the beginning of the array.
		// This is important as we would waste most of the time searching for the destination cup otherwise.
		var destination int32
		if destinationValue >= 10 {
			destination = destinationValue - 1
		} else {
			for i, cup := range cups {
				if cup.value == destinationValue {
					destination = int32(i)
					break
				}
			}
		}

		// Insert cups back into the list.
		last := pickup
		for i := 0; i < 2; i++ {
			last = cups[last].next
		}
		cups[last].next = cups[destination].next
		cups[destination].next = pickup

		// Select new current cup.
		current = cups[current].next
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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
