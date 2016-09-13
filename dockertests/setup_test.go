package dockertests

import (
	"testing"
	"os"
	"github.com/queeno/aptlify/dockertests/dockerutils"
	docker "github.com/fsouza/go-dockerclient"
)

var client *docker.Client
var id string

func TestMain(m *testing.M) {

	dockerutils.StartAptlifyDocker(&client, &id)
	dockerutils.CompileAptlify(client, id)

	exitCode := m.Run()

	dockerutils.StopAptlifyDocker(client, id)
	os.Exit(exitCode)

}
