package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

const (
	INC = "inc"
	DEC = "dec"
)

const (
	EQ  = "=="
	NEQ = "!="
	GT  = ">"
	LT  = "<"
	GE  = ">="
	LE  = "<="
)

type instruction struct {
	src       string
	command   string
	value1    int
	dst       string
	condition string
	value2    int
}

type registers map[string]int

func (r registers) applyInstruction(i instruction) {
	if _, ok := r[i.src]; !ok {
		r[i.src] = 0
	}

	if _, ok := r[i.dst]; !ok {
		r[i.dst] = 0
	}

	result := false

	switch i.condition {
	case EQ:
		result = (r[i.src] == i.value2)
	case NEQ:
		result = (r[i.src] != i.value2)
	case GT:
		result = (r[i.src] > i.value2)
	case GE:
		result = (r[i.src] >= i.value2)
	case LT:
		result = (r[i.src] < i.value2)
	case LE:
		result = (r[i.src] <= i.value2)
	}

	if !result {
		return
	}

	if i.command == INC {
		r[i.dst] += i.value1
	} else {
		r[i.dst] -= i.value1
	}
}

func main() {
	solve(readInstructions())
}

func readInstructions() []instruction {
	instructions := []instruction{}

	for _, line := range utils.FileGetLines("input8.txt") {
		fields := strings.Fields(line)
		newInstruction := instruction{}
		newInstruction.dst = fields[0]
		newInstruction.command = fields[1]
		newInstruction.value1, _ = strconv.Atoi(fields[2])
		newInstruction.src = fields[4]
		newInstruction.condition = fields[5]
		newInstruction.value2, _ = strconv.Atoi(fields[6])
		instructions = append(instructions, newInstruction)
	}

	return instructions
}

func mapToSlice(m map[string]int) []int {
	values := []int{}
	for _, value := range m {
		values = append(values, value)
	}

	return values
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func solve(instructions []instruction) {

	rs := registers{}

	maxValue := 0

	for _, instruction := range instructions {
		rs.applyInstruction(instruction)
		maxValue = max(maxValue, utils.MaxInt(mapToSlice(rs)))
	}

	fmt.Println("Task 1", utils.MaxInt(mapToSlice(rs)))

	fmt.Println("Task 2", maxValue)
}
