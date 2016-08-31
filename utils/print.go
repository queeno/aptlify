package utils

import (
	"fmt"
)

func PrintSlice(s []string) {

	var f string
	for _, m := range s {
		f = fmt.Sprintf("  - %s", m)
		fmt.Println(f)
	}

}
