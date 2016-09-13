package dockerutils

import (
	docker "github.com/fsouza/go-dockerclient"
	"fmt"
	"bytes"
)


// Run commands in docker container
// --------------------------------------------------------------

func runCommand(containerID string, c *docker.Client, cmd ...string) (string, string) {

	var err error

	de := docker.CreateExecOptions{
	        AttachStderr: true,
	        AttachStdin:  true,
	        AttachStdout: true,
	        Tty:          false,
	        Cmd:          cmd,
	        Container:    containerID,
	}

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
	fmt.Println("STDIN:")
	fmt.Println(stdout.String())
	fmt.Println("STDERR:")
	fmt.Println(stderr.String())
	return stdout.String(), stderr.String()
}
