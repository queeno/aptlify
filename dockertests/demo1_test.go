package dockertests

import (
  "testing"
  "fmt"
	"github.com/queeno/aptlify/dockertests/dockerutils"
)


func Test (t *testing.T) {
	dockerutils.PrepareContainer(client, id)
  fmt.Println("simon")
}
