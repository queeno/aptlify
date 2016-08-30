package utils

import (
	"os/exec"
	"strings"
)

func Exec(commandString string) ([]string, error) {

	splitString := strings.Fields(commandString)
	command := splitString[0]
	args := splitString[1:len(splitString)]

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return nil, err
	}

	arrayOut := strings.Split(string(out), "\n")

	return arrayOut, nil

}
