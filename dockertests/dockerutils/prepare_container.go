package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
)



// Prepare container for next test
//-------------------------------------------------------------

func PrepareContainer(c *docker.Client, containerID string) error {

	runCommand(containerID, c, "echo", "Hello", "World")
	return nil
}
