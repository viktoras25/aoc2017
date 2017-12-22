package main

import "fmt"

const INPUT = 335

type node struct {
	value int
	next  *node
}

func (n *node) append(value int) {
	newNode := node{value, n.next}
	n.next = &newNode
}

func main() {

	n := node{0, nil}
	n.next = &n

	n.append(1)
	n = *n.next
	n.append(2)
	fmt.Println(n, n.next, n.next.next)
	return
	numbers := []int{0}
	for i := 1; i <= 2017; i++ {
		// offset := INPUT % len(numbers)
		// fmt.Println(offset)
		// , numbers[:offset]...
		// fmt.Println(numbers[offset:], i, numbers[:offset])
		// head, tail := numbers[:offset+1], numbers[offset+1:]
		// fmt.Println(offset, head, i, tail)
		// newNumbers := []int{i}
		// newNumbers = append(newNumbers, tail...)
		// newNumbers = append(newNumbers, head...)
		// numbers = newNumbers
		// numbers = append(numbers, numbers[:offset]...)
		//numbers = append(numbers[offset:], i, numbers[:offset]...)
		// fmt.Println(numbers)
		numbers = insert(numbers)

		for j := INPUT; j > 0; j-- {
			n = *n.next
		}

		n.append(i)
	}

	// 50000000

	fmt.Println("Part 1", numbers[1])
	fmt.Println(n.next.value)

	// for len(numbers) < 10000 {
	// 	numbers = insert(numbers)
	// 	fmt.Println(len(numbers))
	// }

	// fmt.Println(numbers)

	// for i := 0; i < len(numbers); i++ {
	// 	if numbers[i] == 0 {
	// 		fmt.Println("Part 2", numbers[(i+1)%len(numbers)])
	// 	}
	// }
}

func insert(numbers []int) []int {
	offset := INPUT % len(numbers)
	value := len(numbers)
	head, tail := numbers[:offset+1], numbers[offset+1:]
	newNumbers := []int{value}
	newNumbers = append(newNumbers, tail...)
	newNumbers = append(newNumbers, head...)
	return newNumbers
}
