package gpg

import "fmt"

func Plan() error {
	if keyLoaded("F76221572C52609D") {
		fmt.Println("found F76221572C52609D")
	}
	return nil
}
