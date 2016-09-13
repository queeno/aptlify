package dockerutils

import (
	colour "github.com/fatih/color"
	docker "github.com/fsouza/go-dockerclient"
)

// Compile aptlify in container
//-------------------------------------------------------------

func CompileAptlify(c *docker.Client, containerID string) {

	// Install deps
	colour.Green("installing deps")
	var installDeps = []string{"gom", "install"}
	RunCommand(c, containerID, installDeps...)

	// Build Aptlify
	colour.Green("compiling aptlify")
	var compileAptlifyCommand = []string{"bash", "-c", "gom build -o ${GOPATH}/bin/aptlify"}
	RunCommand(c, containerID, compileAptlifyCommand...)

}
