package utils

import "strconv"

func IsNumeric(str string) bool {
	if _, ok := strconv.Atoi(str); ok == nil {
		return true
	}

	return false
}
