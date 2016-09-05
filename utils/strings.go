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

func uniqueStringsOnLeft(arr1 []string, arr2 []string) ([]string, error) {
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
	return newStrings, nil
}

func DiffStringSlices(arr1 []string, arr2 []string) ([]string, []string, error) {
	var newStrings1 []string
	var newStrings2 []string
	var err error
	if newStrings1, err = uniqueStringsOnLeft(arr1, arr2); err != nil {
		return nil, nil, err
	}
	if newStrings2, err = uniqueStringsOnLeft(arr2, arr1); err != nil {
		return nil, nil, err
	}
	return newStrings1, newStrings2, nil
}
