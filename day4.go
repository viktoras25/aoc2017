package main

import "fmt"
import "strings"
import "sort"
import "viktoras.de/aoc2017/utils"

func main() {
	lines := utils.FileGetLines("input4.txt")

	fmt.Println(task1(lines))
	fmt.Println(task2(lines))
}

func hasDuplicates(strings []string) bool {
	counts := map[string]bool{}

	for _, str := range strings {

		if counts[str] {
			return true
		}

		counts[str] = true
	}

	return false
}

func sortStr(str string) string {
	letters := sort.StringSlice(strings.Split(str, ""))
	letters.Sort()
	return strings.Join(letters, "")
}

func task1(lines []string) int {
	result := 0

	for _, line := range lines {
		if !hasDuplicates(strings.Fields(line)) {
			result++
		}
	}

	return result
}

func task2(lines []string) int {
	result := 0

	for _, line := range lines {

		words := strings.Fields(line)

		for k, word := range words {
			words[k] = sortStr(word)
		}

		if !hasDuplicates(words) {
			result++
		}
	}

	return result
}
