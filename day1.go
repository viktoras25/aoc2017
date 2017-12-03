package main

import "fmt"

import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	digits := stringToDigits(utils.FileGetContents("input1.txt"))

	fmt.Println(task1(digits))
	fmt.Println(task2(digits))
}

func stringToDigits(digits string) []int {

	result := []int{}
	for _, letter := range digits {
		digit, _ := strconv.Atoi(string(letter))
		result = append(result, digit)
	}

	return result
}

func task1(digits []int) int {
	sum := 0

	length := len(digits)
	for k, v := range digits {
		if v == digits[(k+1)%length] {
			sum += v
		}
	}

	return sum
}

func task2(digits []int) int {
	sum := 0

	length := len(digits)
	for k, v := range digits {
		if v == digits[(k+length/2)%length] {
			sum += v
		}
	}

	return sum
}
