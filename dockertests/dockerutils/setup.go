package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
	"fmt"
	"os"
)

// Start docker container
// ------------------------------------------------------------------

func createHostOptions() *docker.HostConfig {

	basePath, _ := os.Getwd()

	volumes := []string{
			basePath, "/root/gowork/src/github.com/queeno/aptlify",
	}

	hostConfig := &docker.HostConfig{
		Binds:			volumes,
		AutoRemove:	true,
	}

	return hostConfig

}

func createImageOptions() docker.CreateContainerOptions {

	image := "aptlify"

  opts := docker.CreateContainerOptions{
      Config: &docker.Config{
          Image:        image,
      },
  }

    return opts
}

func StartAptlifyDocker() (*docker.Client, string){

	var err error

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to Docker daemon: %s", err))
	}

	container, err := client.CreateContainer(createImageOptions())
	if err != nil {
		panic(fmt.Sprintf("Cannot create Docker container: %s", err))
	}

	err = client.StartContainer(container.ID, createHostOptions())
	if err != nil {
		panic(fmt.Sprintf("Cannot start Docker container: %s", err))
	}

	return client, container.ID
}
