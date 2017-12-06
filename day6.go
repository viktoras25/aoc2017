package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

func main() {
	solve(utils.StringsToInts(strings.Fields(utils.FileGetContents("input6.txt"))))
}

func glue(digits []int) int {
	result := 0

	for _, d := range digits {
		if result > 0 {
			result *= 10
		}

		result += d
	}

	return result
}

func reallocate(banks []int) {
	blocks := 0
	i := 0

	// Find the greatest value
	maxValue := utils.MaxInt(banks)

	// And it's index
	for i = 0; i < len(banks) && banks[i] < maxValue; i++ {
	}

	blocks = banks[i]
	banks[i] = 0

	// Distribute
	for blocks > 0 {
		i++
		banks[i%len(banks)]++
		blocks--
	}
}

func solve(banks []int) int {
	steps := 0
	part1steps := 0

	seen := map[int]int{}

	for {
		steps++

		reallocate(banks)

		if _, ok := seen[glue(banks)]; ok {
			if part1steps == 0 {
				fmt.Println("Part 1", steps)
				part1steps = steps

				steps = 0
				seen = map[int]int{}
			} else {
				fmt.Println("Part 2", steps)
				break
			}
		}

		seen[glue(banks)]++
	}

	return steps
}
