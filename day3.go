package main

import "fmt"
import "math"

// import "viktoras.de/aoc2017/utils"

func main() {
	number := 289326

	//number = 1024

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

func task2(number int) int {
	// step := 1

	grid := [][]int{}
	for i := 0; i < 100; i++ {
		grid = append(grid, make([]int, 100))
	}

	r := 50
	c := 50

	grid[r][c] = 1

	c++

	for i := 0; i < 4; i++ {

		// Up
		for grid[r][c-1] != 0 {
			grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]
			// fmt.Println(r, c, grid[r][c])
			// fmt.Println(grid[r-1][c-1], grid[r-1][c], grid[r-1][c+1], grid[r][c-1], grid[r][c+1], grid[r+1][c-1], grid[r+1][c+1], grid[r+1][c])
			r--
		}

		// Top right corner
		grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]

		c--

		for grid[r+1][c] != 0 {
			grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]
			c--
		}

		// Top left corner
		grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]

		r++

		for grid[r][c+1] != 0 {
			grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]
			r++
		}

		// Bottom left corner
		grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]

		c++

		for grid[r-1][c] != 0 {
			grid[r][c] = grid[r-1][c-1] + grid[r-1][c] + grid[r-1][c+1] + grid[r][c-1] + grid[r][c+1] + grid[r+1][c-1] + grid[r+1][c+1] + grid[r+1][c]
			c++
		}
	}

	for i := 40; i < 60; i++ {
		for j := 40; j < 60; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println(" ")
	}

	return 0
}
