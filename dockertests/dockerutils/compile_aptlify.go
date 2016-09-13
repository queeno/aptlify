package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
	colour "github.com/fatih/color"
)

// Compile aptlify in container
//-------------------------------------------------------------

func CompileAptlify(c *docker.Client, containerID string) {

	// Install deps
	colour.Green("installing deps")
	var installDeps = []string{"gom", "install"}
	runCommand(containerID, c, installDeps...)

	// Build Aptlify
  colour.Green("compiling aptlify")
	var compileAptlifyCommand = []string{"bash", "-c", "gom build -o ${GOPATH}/bin/aptlify"}
	runCommand(containerID, c, compileAptlifyCommand...)

}
