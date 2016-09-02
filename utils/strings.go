package utils

import (
	"strings"
)

func SplitStringToSlice(s string) []string {
	rStr := strings.Split(s, "\n")
	if rStr[len(rStr)-1] == "" {
		rStr = rStr[:len(rStr)-1]
	}
	return rStr
}

func IsStringEmpty(element string) bool {
	return len(element) == 0
}
