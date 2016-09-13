package dockertests

import (
	"github.com/queeno/aptlify/dockertests/dockerutils"
	. "gopkg.in/check.v1"
	"testing"
)

func TestDockerTests(t *testing.T) { TestingT(t) }

type DockerTestsSuite struct{}

var _ = Suite(&DockerTestsSuite{})

func Test(t *testing.T) {
	dockerutils.PrepareContainer(client, id)
}

func (s *DockerTestsSuite) TestEmptyConfigFile(c *C) {
	dockerutils.PrepareContainer(client, id)
	testExecuteCommand := []string{"bash", "-c", "echo {} | ./aptlify apply -config /dev/stdin"}
	dockerutils.RunCommand(client, id, testExecuteCommand...)
	testFetchOutput := []string{"bash", "-c", "cat aptlify.state"}
	stdout, _ := dockerutils.RunCommand(client, id, testFetchOutput...)
	c.Check(stdout, Equals, "\x01\x00\x00\x00\x00\x00\x00h{\n  \"mirrors\": null,\n  \"repos\": null,\n  \"gpg_keys\": {\n    \"fingerprint\": null\n  },\n  \"snapshots\": null\n}")
}

func (s *DockerTestsSuite) TestBaseConfigFile(c *C) {
	dockerutils.PrepareContainer(client, id)
	testExecuteCommand := []string{"bash", "-c", "./aptlify apply -config dockertests/data/TestBaseConfigFile.json; cat aptlify.state"}
	dockerutils.RunCommand(client, id, testExecuteCommand...)
	testFetchOutput := []string{"bash", "-c", "cat aptlify.state"}
	stdout, _ := dockerutils.RunCommand(client, id, testFetchOutput...)
	c.Check(stdout, Equals, "\x01\x00\x00\x00\x00\x00\x03W{\n"+
		"  \"mirrors\": [\n"+
		"    {\n"+
		"      \"name\": \"devenvironment-aptly\",\n"+
		"      \"url\": \"http://repo.aptly.info\",\n"+
		"      \"dist\": \"squeeze\",\n"+
		"      \"component\": \"main\",\n"+
		"      \"filter\": [\n"+
		"        {\n"+
		"          \"name\": \"aptly\",\n"+
		"          \"version\": \"0.9.7\"\n"+
		"        }\n"+
		"      ],\n"+
		"      \"filter-with-deps\": false\n"+
		"    },\n"+
		"    {\n"+
		"      \"name\": \"haproxy\",\n"+
		"      \"url\": \"http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu\",\n"+
		"      \"dist\": \"trusty\",\n"+
		"      \"component\": \"\",\n"+
		"      \"filter\": null,\n"+
		"      \"filter-with-deps\": false\n"+
		"    }\n"+
		"  ],\n"+
		"  \"repos\": [\n"+
		"    {\n"+
		"      \"name\": \"devenvironment-internal\"\n"+
		"    },\n"+
		"    {\n"+
		"      \"name\": \"test\"\n"+
		"    }\n"+
		"  ],\n"+
		"  \"gpg_keys\": {\n"+
		"    \"fingerprint\": [\n"+
		"      \"9E3E53F19C7DE460\",\n"+
		"      \"353525F9\",\n"+
		"      \"505D97A41C61B9CD\",\n"+
		"      \"1C61B9CD\"\n"+
		"    ]\n"+
		"  },\n"+
		"  \"snapshots\": [\n"+
		"    {\n"+
		"      \"name\": \"devenvironment\",\n"+
		"      \"resources\": null,\n"+
		"      \"revision\": 1\n"+
		"    }\n"+
		"  ]\n"+
		"}")
}
