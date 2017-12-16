package main

import "fmt"
import "strings"
import "strconv"
import "viktoras.de/aoc2017/utils"

func main() {
	commands := strings.Split(utils.FileGetContents("input16.txt"), ",")

	letters := "abcdefghijklmnop"
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

	fmt.Println(letters)

	fmt.Println("Part 1", 0)
	fmt.Println("Part 2", 0)
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
