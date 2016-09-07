package utils

import (
	"os/exec"
)

func Exec(command string, args... string) ([]string, error) {

	out, err := exec.Command(command, args...).CombinedOutput()
	arrayOut := SplitStringToSlice(string(out))
	if err != nil {
		return arrayOut, err
	}

	return arrayOut, nil

}
