package main

import "fmt"
import "strconv"
import "strings"
import "viktoras.de/aoc2017/utils"

func main() {
	lines := utils.FileGetLines("input2.txt")

	fmt.Println(task1(lines))
	fmt.Println(task2(lines))
}

func task1(lines []string) int {
	sum := 0

	for _, line := range lines {
		numbers := []int{}

		for _, v := range strings.Split(line, "\t") {
			number, _ := strconv.Atoi(string(v))
			numbers = append(numbers, number)
		}

		sum += (utils.MaxInt(numbers) - utils.MinInt(numbers))
	}

	return sum
}

func task2(lines []string) int {
	sum := 0

	for _, line := range lines {
		numbers := []int{}

		for _, v := range strings.Split(line, "\t") {
			number, _ := strconv.Atoi(string(v))
			numbers = append(numbers, number)
		}

		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if numbers[i]%numbers[j] == 0 {
					sum += int(numbers[i] / numbers[j])
					break
				}

				if numbers[j]%numbers[i] == 0 {
					sum += int(numbers[j] / numbers[i])
					break
				}
			}
		}
	}

	return sum
}
