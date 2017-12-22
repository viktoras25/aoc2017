package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	commands := strings.Split(utils.FileGetContents("input16.txt"), ",")
	letters := "abcdefghijklmnop"

	fmt.Println("Part 1", dance(letters, commands))

	// Period from one value to the same value is experimentally found to be 60
	for i := 0; i < 10000000%60; i++ {
		letters = dance(letters, commands)
	}

	fmt.Println("Part 2", letters)
}

func dance(letters string, commands []string) string {
	for _, command := range commands {
		switch command[0] {
		case 's':
			offset, _ := strconv.Atoi(command[1:])
			letters = rotate(letters, offset)
		case 'x':
			placesStr := strings.Split(command[1:], "/")
			placesInt := utils.StringsToInts(placesStr)
			letters = swap(letters, placesInt[0], placesInt[1])
		case 'p':
			letters = partner(letters, string(command[1]), string(command[3]))
		}
	}

	return letters
}

func rotate(str string, offset int) string {
	offset = len(str) - offset
	return str[offset:] + str[:offset]
}

func swap(str string, a, b int) string {
	if a > b {
		a, b = b, a
	}

	return string(str[:a]) + string(str[b]) + string(str[a+1:b]) + string(str[a]) + string(str[b+1:])
}

func partner(str, a, b string) string {
	str = strings.Replace(str, a, "#", -1)
	str = strings.Replace(str, b, a, -1)
	str = strings.Replace(str, "#", b, -1)
	return str
}
