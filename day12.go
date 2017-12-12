package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	connections := readConnections()

	groups := [][]int{}

	nodesToVisit := []int{0}
	visitedNodes := []int{}

	for len(connections) > 0 {

		for len(nodesToVisit) > 0 {
			currentNode := nodesToVisit[0]
			nodesToVisit = nodesToVisit[1:]

			visitedNodes = append(visitedNodes, currentNode)
			nodesToVisit = append(nodesToVisit, sliceDiff(connections[currentNode], visitedNodes)...)
		}

		groups = append(groups, visitedNodes)

		for _, v := range visitedNodes {
			delete(connections, v)
		}

		nodesToVisit = []int{firstKey(connections)}
		visitedNodes = []int{}
	}

	fmt.Println("Part 1", len(groups[0]))
	fmt.Println("Part 1", len(groups))
}

func readConnections() map[int][]int {
	lines := utils.FileGetLines("input12.txt")

	connections := map[int][]int{}

	for _, line := range lines {
		fields := strings.Fields(line)
		for k, field := range fields {
			fields[k] = strings.Trim(field, ", ")
		}

		i, _ := strconv.Atoi(fields[0])
		connections[i] = utils.StringsToInts(fields[2:])
	}

	return connections
}

func sliceDiff(a, b []int) []int {
	result := []int{}

	for _, va := range a {
		add := true
		for _, vb := range b {
			if va == vb {
				add = false
			}
		}

		if add {
			result = append(result, va)
		}
	}

	return result
}

func firstKey(m map[int][]int) int {
	for k, _ := range m {
		return k
	}

	return 0
}
