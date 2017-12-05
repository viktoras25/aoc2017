package main

import "fmt"
import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	lines := utils.FileGetLines("input5.txt")

	numbers := []int{}
	for _, v := range lines {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}

	fmt.Println(task1(append([]int{}, numbers...)))
	fmt.Println(task2(append([]int{}, numbers...)))
}

func task1(numbers []int) int {
	result := 0
	position := 0
	jump := 0

	for position >= 0 && position < len(numbers) {
		jump = numbers[position]
		numbers[position]++
		position += jump
		result++
	}

	return result
}

func task2(numbers []int) int {
	result := 0
	position := 0
	jump := 0

	for position >= 0 && position < len(numbers) {
		jump = numbers[position]
		if jump >= 3 {
			numbers[position]--
		} else {
			numbers[position]++
		}
		position += jump
		result++
	}

	return result
}
