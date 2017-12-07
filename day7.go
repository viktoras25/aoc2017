package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

type node struct {
	Name        string
	Weight      int
	Parent      *node
	Subnodes    []*node
	TotalWeight int
}

func (n *node) addChild(m *node) {
	n.Subnodes = append(n.Subnodes, m)
	m.Parent = n
}

func (n *node) calcTotalWeight() int {
	n.TotalWeight = n.Weight

	for _, subnode := range n.Subnodes {
		n.TotalWeight += subnode.calcTotalWeight()
	}

	return n.TotalWeight
}

func (n *node) isBalanced() bool {
	if len(n.Subnodes) == 0 {
		return true
	}

	weight := n.Subnodes[0].TotalWeight

	for _, v := range n.Subnodes {
		if weight != v.TotalWeight {
			return false
		}
	}

	return true
}

func main() {
	lines := utils.FileGetLines("input7.txt")

	nodes := map[string]*node{}

	// Read nodes
	for _, line := range lines {
		parts := strings.Fields(line)
		weight, _ := strconv.Atoi(parts[1][1 : len(parts[1])-1])
		nodes[parts[0]] = &node{parts[0], weight, nil, make([]*node, 0), 0}
	}

	for _, line := range lines {
		parts := strings.Fields(line)
		currentNode := parts[0]
		childNodes := parts[2:]
		for i := 1; i < len(childNodes); i++ {
			name := strings.Trim(childNodes[i], ",")
			nodes[currentNode].addChild(nodes[name])
		}
	}

	current := nodes[strings.Fields(lines[0])[0]]

	for current.Parent != nil {
		current = current.Parent
	}

	fmt.Println("Part 1", current.Name)

	current.calcTotalWeight()

	for {
		balanced := true

		for _, v := range current.Subnodes {
			if v.isBalanced() {
				continue
			}

			balanced = false
			current = v
		}

		// All subnodes themselves are balanced, then it's one of them to have the wrong weight
		if balanced {

			weights := []int{}
			for _, s := range current.Subnodes {
				weights = append(weights, s.TotalWeight)
			}

			// Works only if the node is heavier
			maxWeight := utils.MaxInt(weights)
			minWeight := utils.MinInt(weights)
			diff := maxWeight - minWeight

			for _, s := range current.Subnodes {
				if s.TotalWeight == maxWeight {
					fmt.Println("Part 2", s.Weight-diff)
					return
				}
			}

			return
		}
	}

	return
}
