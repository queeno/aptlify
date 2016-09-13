package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
)

// Prepare container for next test
//-------------------------------------------------------------

func PrepareContainer(c *docker.Client, containerID string) {

	// Clean up aptly
	var removeAptlyContainer = []string{"rm", "-rf", "/root/.aptly"}
	runCommand(containerID, c, removeAptlyContainer...)

}
