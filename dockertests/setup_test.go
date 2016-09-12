package dockertests

import (
	"testing"
	"os"
	"github.com/queeno/aptlify/dockertests/dockerutils"
	docker "github.com/fsouza/go-dockerclient"
)

var container *docker.Container

func TestMain(m *testing.M) {

	dockerutils.StartAptlifyDocker()
	exitCode := m.Run()
	os.Exit(exitCode)

}
