package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	lines := utils.FileGetLines("input18.txt")
	registers := map[string]int{}

	var dst, src string
	var srcValue, dstValue int
	var played int

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Fields(line)
		dst = parts[1]
		dstValue = getValue(registers, dst)

		if len(parts) > 2 {
			src = parts[2]
			srcValue = getValue(registers, src)
		}

		switch parts[0] {
		case "set":
			registers[dst] = srcValue
		case "add":
			registers[dst] += srcValue
		case "mul":
			registers[dst] *= srcValue
		case "mod":
			registers[dst] %= srcValue
		case "rcv":
			if dstValue > 0 {
				fmt.Println("Recover", dstValue)
				i = len(lines)
			}
		case "jgz":
			if dstValue > 0 {
				i += srcValue
				i--
				if i < 0 || i > len(lines) {
					break
				}
			}
		case "snd":
			played = dstValue
			fmt.Println("Play ", dstValue)
		}

		fmt.Println(line, registers)
		// fmt.Scanln()
	}

	fmt.Println("Part 1", played)
	fmt.Println("Part 2", 0)
}

func getValue(registers map[string]int, op string) int {
	value := 0

	if utils.IsNumeric(op) {
		value, _ = strconv.Atoi(op)
	} else {
		value, _ = registers[op]
	}

	return value
}
