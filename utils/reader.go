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
