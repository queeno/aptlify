package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
)



// Prepare container for next test
//-------------------------------------------------------------

func prepareContainer(containerID string, c *docker.Client) error {

	runCommand(containerID, c, "echo", "Hello", "World")
	return nil
}
