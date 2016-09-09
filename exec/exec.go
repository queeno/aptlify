package exec

import (
	"github.com/queeno/aptlify/utils"
	ex "os/exec"
)

func Exec(command string, args ...string) ([]string, error) {

	out, err := ex.Command(command, args...).CombinedOutput()
	arrayOut := utils.SplitStringToSlice(string(out))
	if err != nil {
		return arrayOut, err
	}

	return arrayOut, nil

}
