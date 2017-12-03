package utils

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
