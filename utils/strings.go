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

func uniqueStringsOnLeft(arr1 []string, arr2 []string) []string {
	var thisStrFound bool
	var newStrings []string
	for _, str1 := range arr1 {
		thisStrFound = false
		for _, str2 := range arr2 {
			if str1 == str2 {
				thisStrFound = true
			}
		}
		if !thisStrFound {
			newStrings = append(newStrings, str1)
		}
	}
	return newStrings
}

func DiffStringSlices(arr1 []string, arr2 []string) ([]string, []string) {
	var newStrings1 []string
	var newStrings2 []string
	newStrings1 = uniqueStringsOnLeft(arr1, arr2)
	newStrings2 = uniqueStringsOnLeft(arr2, arr1)
	return newStrings1, newStrings2
}
