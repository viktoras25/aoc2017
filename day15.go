package main

import "fmt"

type list []int

func main() {

	genA := 722
	genB := 354
	judge := 0

	for i := 0; i < 40000000; i++ {
		genA = nextA(genA)
		genB = nextB(genB)

		if genA&65535 == genB&65535 {
			judge++
		}
	}

	fmt.Println("Part 1", judge)

	genA = 722
	genB = 354
	judge = 0

	for i := 0; i < 5000000; i++ {
		genA = nextA2(genA)
		genB = nextB2(genB)

		if genA&65535 == genB&65535 {
			judge++
		}
	}

	fmt.Println("Part 2", judge)
}

func nextA(value int) int {
	return value * 16807 % 2147483647
}

func nextB(value int) int {
	return value * 48271 % 2147483647
}

func nextA2(value int) int {
	value = nextA(value)
	for value%4 != 0 {
		value = nextA(value)
	}
	return value
}

func nextB2(value int) int {
	value = nextB(value)
	for value%8 != 0 {
		value = nextB(value)
	}
	return value
}
