package utils

import "os"
import "bufio"

func Fopen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}

	return file
}

func FileGetContents(filename string) string {
	scanner := bufio.NewScanner(Fopen(filename))
	scanner.Scan()
	return scanner.Text()
}

func FileGetLines(filename string) []string {
	result := []string{}

	scanner := bufio.NewScanner(Fopen(filename))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
