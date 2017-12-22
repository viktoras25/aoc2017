package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

func main() {
	grid := readGrid()

	x, y := 0, 0
	xd, yd := 0, 1

	result := ""
	steps := 0

	for grid[y][x] != "|" {
		x++
	}

	for {
		for grid[y][x] != "+" && grid[y][x] != " " {
			x += xd
			y += yd

			steps++

			if grid[y][x] != "+" && grid[y][x] != "-" && grid[y][x] != "|" {
				result += grid[y][x]
			}
		}

		if grid[y][x] == " " {
			break
		}

		if xd != 0 {
			xd = 0
			yd = 0

			if y < len(grid)-1 && grid[y+1][x] == "|" {
				yd = 1
			}

			if y > 0 && grid[y-1][x] == "|" {
				yd = -1
			}

			if y == len(grid)-1 || (y < len(grid)-1 && grid[y+1][x] == " ") {
				yd = -1
			}

			if y == 0 || (y > 0 && grid[y-1][x] == " ") {
				yd = 1
			}

		} else {
			xd = 0
			yd = 0

			if x < len(grid[y]) && grid[y][x+1] == "-" {
				xd = 1
			}

			if x > 0 && grid[y][x-1] == "-" {
				xd = -1
			}

			if x >= len(grid[y])-1 || (x < len(grid[y])-1 && grid[y][x+1] == " ") {
				xd = -1
			}

			if x == 0 || (x > 0 && grid[y][x-1] == " ") {
				xd = 1
			}
		}

		if xd == 0 && yd == 0 {
			fmt.Println("Stopped")
			break
		}

		x += xd
		y += yd
		steps++
	}

	fmt.Println("Part 1", result)
	fmt.Println("Part 2", steps)
}

func readGrid() [][]string {
	grid := [][]string{}

	lines := utils.FileGetLines("input19.txt")

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
