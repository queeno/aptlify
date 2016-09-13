package dockerutils

import (
	colour "github.com/fatih/color"
	docker "github.com/fsouza/go-dockerclient"
)

// Prepare container for next test
//-------------------------------------------------------------

func PrepareContainer(c *docker.Client, containerID string) {

	// Clean up aptly
	colour.Green("PREP: removing aptly DB")
	var removeAptlyDB = []string{"rm", "-rf", "/root/.aptly"}
	RunCommand(c, containerID, removeAptlyDB...)

	colour.Green("PREP: removing aptlify.state")
	var removeAptlifyState = []string{"rm", "-f", "/aptlify/aptlify.state"}
	RunCommand(c, containerID, removeAptlifyState...)

}
