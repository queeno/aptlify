package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
	colour "github.com/fatih/color"
)

// Prepare container for next test
//-------------------------------------------------------------

func PrepareContainer(c *docker.Client, containerID string) {

	// Clean up aptly
	colour.Green("PREP: removing aptly DB")
	var removeAptlyDB = []string{"rm", "-rf", "/root/.aptly"}
	runCommand(containerID, c, removeAptlyDB...)

	colour.Green("PREP: removing aptlify.state")
	var removeAptlifyState = []string{"rm", "-f", "/aptlify/aptlify.state"}
	runCommand(containerID, c, removeAptlifyState...)

}
