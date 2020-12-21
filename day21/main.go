package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Food struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func main() {
	foods := make([]Food, 0)
	ingredients := make(map[string]bool)
	allergens := make(map[string]bool)

	lines := readLines("input.txt")
	allergenRegex := regexp.MustCompile(` \(contains (.*)\)`)
	for _, line := range lines {
		food := Food{make(map[string]bool), make(map[string]bool)}
		matches := allergenRegex.FindStringSubmatch(line)
		line = strings.ReplaceAll(line, matches[0], "")
		for _, ingredient := range strings.Split(line, " ") {
			ingredients[ingredient] = true
			food.ingredients[ingredient] = true
		}
		for _, allergen := range strings.Split(matches[1], ", ") {
			allergens[allergen] = true
			food.allergens[allergen] = true
		}
		foods = append(foods, food)
	}

	// Candidates maps from an ingredient to the set of the potential allergens it can contain.
	candidates := make(map[string]map[string]bool)
	for ingredient := range ingredients {
		candidates[ingredient] = make(map[string]bool)
	}

	// Populate candidates by looking at the allergens of each food an ingredient appears in.
	for _, food := range foods {
		for ingredient := range food.ingredients {
			for allergen := range food.allergens {
				candidates[ingredient][allergen] = true
			}
		}
	}

	// If an allergen is listed, the ingredient has to be in the list.
	// Therefore, if we can find a food that has the allergen, but not the ingredient,
	// we can remove that allergen from the ingredients candidates.
	for ingredient, allergens := range candidates {
		for allergen := range allergens {
			for _, food := range foods {
				if food.allergens[allergen] && !food.ingredients[ingredient] {
					delete(allergens, allergen)
					break
				}
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		count := 0
		for ingredient, allergens := range candidates {
			if len(allergens) == 0 {
				for _, food := range foods {
					if food.ingredients[ingredient] {
						count++
					}
				}
			}
		}
		fmt.Println(count)
	}

	type Mapping struct {
		ingredient, allergen string
	}

	// If an ingredient has only one candidate, we know it must contain that allergen.
	// Then we can remove the allergen from the other candidate lists,
	// as each allergen is found in only one ingredient.
	var dangerous []Mapping
	for {
		changed := false
		for ingredient, allergens := range candidates {
			if len(allergens) == 1 {
				var allergen string
				for a := range allergens {
					allergen = a
				}
				dangerous = append(dangerous, Mapping{ingredient, allergen})
				for _, allergens := range candidates {
					delete(allergens, allergen)
				}
				changed = true
				break
			}
		}
		if !changed {
			break
		}
	}

	sort.Slice(dangerous, func(i, j int) bool {
		return dangerous[i].allergen < dangerous[j].allergen
	})

	{
		fmt.Println("--- Part Two ---")
		for i, item := range dangerous {
			fmt.Print(item.ingredient)
			if i+1 < len(dangerous) {
				fmt.Print(",")
			}
		}
		fmt.Println()
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
