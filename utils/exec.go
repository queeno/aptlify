package utils

import (
	"os/exec"
	"fmt"
	ctx "github.com/queeno/aptlify/context"
)

func Exec (string command) (string, error) {

	ctx.Logging.Info.Println(fmt.Sprintf("Running command: %s", command))

	out, err := exec.Command(command).Output()
	if err != nil {
		return "", err
	}

	return out, nil

}
