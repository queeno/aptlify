package dockertests

import (
  "testing"
	"github.com/queeno/aptlify/dockertests/dockerutils"
)


func Test (t *testing.T) {
	dockerutils.PrepareContainer(client, id)


}
