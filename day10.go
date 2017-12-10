package main

import "fmt"
import "strings"
import "viktoras.de/aoc2017/utils"

type list []int

func main() {

	numbers := list{}
	for i := 0; i < 256; i++ {
		numbers = append(numbers, i)
	}

	input := utils.FileGetContents("input10.txt")

	fmt.Println("Part 1", task1(append([]int{}, numbers...), input))
	fmt.Println("Part 2", task2(append([]int{}, numbers...), input))
}

func task1(numbers list, input string) int {
	skipSize := 0
	currentPosition := 0

	lengths := utils.StringsToInts(strings.Split(input, ","))

	for _, length := range lengths {
		numbers.twist(currentPosition, length)
		currentPosition += length + skipSize
		skipSize++
	}

	return numbers[0] * numbers[1]
}

func task2(numbers list, input string) string {

	lengths := list{}
	for _, b := range []byte(input) {
		lengths = append(lengths, int(b))
	}

	lengths = append(lengths, 17, 31, 73, 47, 23)

	skipSize := 0
	currentPosition := 0

	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			numbers.twist(currentPosition, length)
			currentPosition += length + skipSize
			skipSize++
		}
	}

	return hashize(numbers)
}

func (l list) slice(from, to int) list {
	length := len(l)

	if to < from {
		to += length
	}

	result := list{}

	for i := from; i < to; i++ {
		result = append(result, l[i%length])
	}

	return result
}

func (l list) reverse() {
	length := len(l) - 1
	for i := 0; i < len(l)/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}

func (l list) twist(position, length int) {
	slice := l.slice(position, position+length)
	slice.reverse()

	for i := 0; i < length; i++ {
		l[(position+i)%len(l)] = slice[i]
	}
}

func hashize(numbers []int) string {
	result := ""

	for i := 0; i < 16; i++ {
		tmp := numbers[i*16]
		for _, n := range numbers[i*16+1 : i*16+16] {
			tmp ^= n
		}
		result += fmt.Sprintf("%x", tmp)
	}

	return result
}
