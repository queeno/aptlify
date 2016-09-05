package utils

import (
	"github.com/queeno/aptlify/config"
)

func uniqueFiltersOnLeft(filts1 []config.AptlyFilterStruct, filts2 []config.AptlyFilterStruct) ([]config.AptlyFilterStruct, error) {
	var thisOneFound bool
	var newFilters []config.AptlyFilterStruct
	for _, f1 := range filts1 {
		thisOneFound = false
		for _, f2 := range filts2 {
			if f1.Equals(f2) {
				thisOneFound = true
			}
		}
		if !thisOneFound {
			newFilters = append(newFilters, f1)
		}
	}
	return newFilters, nil
}

func DiffFilterSlices(filts1 []config.AptlyFilterStruct, filts2 []config.AptlyFilterStruct) ([]config.AptlyFilterStruct, []config.AptlyFilterStruct, error) {
	var newFilters1 []config.AptlyFilterStruct
	var newFilters2 []config.AptlyFilterStruct
	var err error
	if newFilters1, err = uniqueFiltersOnLeft(filts1, filts2); err != nil {
		return nil, nil, err
	}
	if newFilters2, err = uniqueFiltersOnLeft(filts2, filts1); err != nil {
		return nil, nil, err
	}
	return newFilters1, newFilters2, nil
}
