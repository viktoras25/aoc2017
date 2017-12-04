package main

import "fmt"
import "math"

type grid [][]int

func main() {
	number := 289326

	fmt.Println(task1(number))
	fmt.Println(task2(number))
}

func task1(number int) int {
	// 9 is a lower right corner of a square with side 3. Given a number,
	// find the closest odd square root. This number will be on one of the
	// sides of the next layer/square
	sideSize := int(math.Sqrt(float64(number)))

	// In case of 23, sqrt is 4.7, and we need an odd number
	if sideSize%2 == 0 {
		sideSize -= 1
	}

	// Lower right corner of a square with a side of sideSize
	lowerRight := sideSize * sideSize

	// How many numbers are left, after we "subtract" sides
	shiftFromCorner := (number - lowerRight) % (sideSize + 1)

	// Distance from the center of the side
	shiftFromMiddle := shiftFromCorner - (sideSize+1)/2

	// Distance from 1
	distance := shiftFromMiddle + sideSize/2

	if sideSize%2 != 0 {
		distance++
	}

	return distance
}

func calcCell(grid grid, r, c int) {
	grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]
}

func task2(number int) int {
	grid := grid{}
	for i := 0; i < 12; i++ {
		grid = append(grid, make([]int, 12))
	}

	r := 5
	c := 5

	grid[r][c] = 1

	c++

	for i := 0; i < 4; i++ {

		// Up
		for grid[r][c-1] != 0 {
			calcCell(grid, r, c)
			r--
		}

		// Top right corner
		calcCell(grid, r, c)

		for grid[r+1][c] != 0 {
			c--
			calcCell(grid, r, c)
		}

		// Top left corner
		calcCell(grid, r, c)

		for grid[r][c+1] != 0 {
			r++
			calcCell(grid, r, c)
		}

		// Bottom left corner
		calcCell(grid, r, c)

		for grid[r-1][c] != 0 {
			c++
			calcCell(grid, r, c)
		}
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%7d\t", grid[i][j])
		}
		fmt.Println(" ")
	}

	return 0
}
