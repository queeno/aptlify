package aptly

import (
	"github.com/queeno/aptlify/utils"
)

type AptlyCli struct {}

var string aptlyCmd := "/usr/local/bin/aptly"



func (a AptlyCli) Mirror_list() ([]string, error) {

	cmd := fmt.Sprintf("%s mirror list", aptlyCmd)

	out, err := utils.Exec(cmd)
	if err != nil {
		return nil, err
	}

	mirrors := utils.SplitStringToSlice(out)

	return mirrors
}
