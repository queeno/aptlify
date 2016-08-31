package utils

import (
	"os/exec"
	"strings"
)

func Exec(commandString string) ([]string, error) {

	splitString := strings.Fields(commandString)
	command := splitString[0]
	args := splitString[1:len(splitString)]

	out, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	arrayOut := SplitStringToSlice(string(out))

	return arrayOut, nil

}
