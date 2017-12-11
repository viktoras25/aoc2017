package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

const (
	D_N  = "n"
	D_NE = "ne"
	D_NW = "nw"
	D_S  = "s"
	D_SE = "se"
	D_SW = "sw"
)

type direction string
type directions []direction

func main() {
	solve(strings.Split(utils.FileGetContents("input11.txt"), ","))
}

func solve(input []string) {
	x, y := 0, 0

	distances := []int{}

	for _, v := range input {
		switch v {
		case D_N:
			y -= 2
		case D_NE:
			y -= 1
			x += 1
		case D_NW:
			y -= 1
			x -= 1
		case D_S:
			y += 2
		case D_SE:
			y += 1
			x += 1
		case D_SW:
			y += 1
			x -= 1
		}

		distances = append(distances, getDistance(x, y))
	}

	fmt.Println("Part 1", distances[len(distances)-1])
	fmt.Println("Part 2", utils.MaxInt(distances))
}

func getDistance(x, y int) (distance int) {
	distance = 0

	if y > x {
		x, y = y, x
	}

	return y + (x-y)/2
}
