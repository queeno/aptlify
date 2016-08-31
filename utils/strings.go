package utils

import (
	"strings"
)

func SplitStringToSlice(s string) []string {
  return strings.Split(s, "\n")
}

func StringToSlice(array ...string) []string {

	var strings_arr []string

	for _, e := range array {
		if ! IsStringEmpty(e) {
			strings_arr = append(strings_arr, e)
		}
	}
	return strings_arr
}

func IsStringEmpty(element string) bool {
	return len(element) == 0
}
