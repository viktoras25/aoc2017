package main

import "fmt"
import "strconv"
import "viktoras.de/aoc2017/utils"

type list []int

func main() {
	input := "xlqgujun"

	grid := generateGrid(input)

	part1 := 0
	for _, row := range grid {
		for _, value := range row {
			if value == 1 {
				part1++
			}
		}
	}

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", countGroups(grid))
}

func binary(number int) []int {
	result := []int{}

	for number > 0 {
		result = append(result, number&1)
		number >>= 1
	}

	utils.ReverseInts(result)

	// prepend with enough 0's
	result = append(make([]int, 8-len(result)), result...)

	return result
}

func generateGrid(input string) [][]int {
	grid := [][]int{}

	for i := 0; i < 128; i++ {
		row := []int{}
		hashNums := utils.KnotHash(input + "-" + strconv.Itoa(i))
		for _, v := range hashNums {
			row = append(row, binary(v)...)
		}
		grid = append(grid, row)
	}

	return grid
}

func countGroups(grid [][]int) int {
	result := 0

	for y, row := range grid {
		for x, value := range row {
			if value == 1 {
				result++
				trackGroup(grid, x, y)
			}
		}
	}

	return result
}

func trackGroup(grid [][]int, x int, y int) {
	if grid[y][x] == 0 {
		return
	}

	grid[y][x] = 0

	if x >= 1 {
		trackGroup(grid, x-1, y)
	}

	if y >= 1 {
		trackGroup(grid, x, y-1)
	}

	if x < len(grid[y])-1 {
		trackGroup(grid, x+1, y)
	}

	if y < len(grid)-1 {
		trackGroup(grid, x, y+1)
	}
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, value := range row {
			fmt.Print(value)
		}
		fmt.Println("")
	}
}
