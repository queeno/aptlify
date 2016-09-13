package dockerutils

import (
	"bytes"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
)

// Run commands in docker container
// --------------------------------------------------------------

func RunCommand(c *docker.Client, containerID string, cmd ...string) (string, string) {

	var err error

	de := docker.CreateExecOptions{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Tty:          false,
		Cmd:          cmd,
		Container:    containerID,
	}
	fmt.Sprintf("derp")
	var exec *docker.Exec

	if exec, err = c.CreateExec(de); err != nil {
		panic(fmt.Sprintf("Error creating docker exec command: %s", err))
	}

	var stdout, stderr bytes.Buffer

	opts := docker.StartExecOptions{
		OutputStream: &stdout,
		ErrorStream:  &stderr,
		RawTerminal:  true,
	}

	if err = c.StartExec(exec.ID, opts); err != nil {
		panic(fmt.Sprintf("Error running command in docker container: %s", err))
	}

	return stdout.String(), stderr.String()
}
