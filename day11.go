package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

func main() {
	input := strings.Split(utils.FileGetContents("input11.txt"), ",")

	x, y := 0, 0

	distances := []int{}

	for _, v := range input {
		switch v {
		case "n":
			y -= 2
		case "ne":
			y -= 1
			x += 1
		case "nw":
			y -= 1
			x -= 1
		case "s":
			y += 2
		case "se":
			y += 1
			x += 1
		case "sw":
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
