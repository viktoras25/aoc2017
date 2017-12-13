package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

func main() {
	layers := readLayers()

	fmt.Println("Part 1", task1(layers))
	fmt.Println("Part 1", task2(layers))
}

func task1(layers map[int]int) int {
	return calcSeverity(layers, 0)
}

func task2(layers map[int]int) int {
	offset := 0
	for calcSeverity(layers, offset) > 0 {
		offset++
	}
	return offset
}

func readLayers() map[int]int {
	lines := utils.FileGetLines("input13.txt")

	layers := map[int]int{}

	for _, line := range lines {
		values := utils.StringsToInts(strings.Split(line, ": "))
		layers[values[0]] = values[1]
	}

	return layers
}

func calcSeverity(layers map[int]int, offset int) int {
	severity := 0
	for k, v := range layers {
		k += offset

		if k == 0 {
			continue
		}

		if k%((v-1)*2) == 0 {
			severity += k * v
		}
	}

	return severity
}
