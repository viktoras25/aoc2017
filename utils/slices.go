package utils

import "strconv"

func MinInt(list []int) int {
	if len(list) == 0 {
		panic("Empty list")
	}

	min := list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}

	return min
}

func MaxInt(list []int) int {
	if len(list) == 0 {
		panic("Empty list")
	}

	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}

	return max
}

func StringsToInts(strings []string) []int {
	result := []int{}

	for _, string := range strings {
		number, _ := strconv.Atoi(string)
		result = append(result, number)
	}

	return result
}

func ReverseInts(ints []int) {
	length := len(ints) - 1
	for i := 0; i < len(ints)/2; i++ {
		ints[i], ints[length-i] = ints[length-i], ints[i]
	}
}
