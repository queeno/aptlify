package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
	"path/filepath"
	"fmt"
	"os"
)

// Start docker container
// ------------------------------------------------------------------

func createHostOptions() *docker.HostConfig {

  currentPath, _ := os.Getwd()
	basePath := filepath.Dir(currentPath)
	aptlifyBind := fmt.Sprintf("%s:/root/gowork/src/github.com/queeno/aptlify:rw", basePath)

	hostConfig := &docker.HostConfig{
		Binds:			[]string{aptlifyBind},
		AutoRemove:	true,
		Privileged: false,
	}

	return hostConfig

}

func createImageOptions() *docker.Config {

  imageConfig := &docker.Config{
          Image:        "aptlify",
					WorkingDir:		"/aptlify",
  }

	return imageConfig
}

func StartAptlifyDocker(dockClient **docker.Client, dockId *string) {

	var err error

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to Docker daemon: %s", err))
	}

	createContOps := docker.CreateContainerOptions{
			 Name:       "aptlify",
			 Config:     createImageOptions(),
			 HostConfig: createHostOptions(),
	 }

	container, err := client.CreateContainer(createContOps)
	if err != nil {
		panic(fmt.Sprintf("Cannot create Docker container: %s", err))
	}

	err = client.StartContainer(container.ID, createHostOptions())
	if err != nil {
		panic(fmt.Sprintf("Cannot start Docker container: %s", err))
	}

	*dockClient = client
	*dockId = container.ID

}

func StopAptlifyDocker(dockClient *docker.Client, dockId string) {

	if err := dockClient.RemoveContainer(docker.RemoveContainerOptions{
		ID:    dockId,
		Force: true,
	}); err != nil {
		panic(fmt.Sprintf("cannot remove container: %s", err))
	}

}
