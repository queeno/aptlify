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

func (s *DockerTestsSuite) TestConfigMigration(c *C) {
	dockerutils.PrepareContainer(client, id)
	//Apply base configuration
	testExecuteCommand := []string{"bash", "-c", "./aptlify apply -config dockertests/data/TestBaseConfigFile.json; cat aptlify.state"}
	dockerutils.RunCommand(client, id, testExecuteCommand...)
	testExecuteCommand = []string{"bash", "-c", "./aptlify apply -config dockertests/data/TestConfigMigration.json; cat aptlify.state"}
	dockerutils.RunCommand(client, id, testExecuteCommand...)
	testFetchOutput := []string{"bash", "-c", "cat aptlify.state"}
	stdout, _ := dockerutils.RunCommand(client, id, testFetchOutput...)
	c.Check(stdout, Equals, ""+
		"\x01\x00\x00\x00\x00\x00\x03V{\n"+
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
		"      \"url\": \"http://ppa.launchpad.net/vbernat/haproxy-1.6/ubuntu\",\n"+
		"      \"dist\": \"vivid\",\n"+
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
		"      \"revision\": 2\n"+
		"    }\n"+
		"  ]\n"+
		"}")
}
func (s *DockerTestsSuite) TestAptlifyPlan(c *C) {
	dockerutils.PrepareContainer(client, id)
	//Apply base configuration
	testExecuteCommand := []string{"bash", "-c", "./aptlify plan -config dockertests/data/TestBaseConfigFile.json"}
	stdout, _ := dockerutils.RunCommand(client, id, testExecuteCommand...)
	c.Check(stdout, Equals, ""+
		"\x01\x00\x00\x00\x00\x00\x022+gpg key 9E3E53F19C7DE460 will be added. Reason: GPG key not found\n"+
		"+gpg key 353525F9 will be added. Reason: GPG key not found\n"+
		"+gpg key 505D97A41C61B9CD will be added. Reason: GPG key not found\n"+
		"+gpg key 1C61B9CD will be added. Reason: GPG key not found\n"+
		"+mirror devenvironment-aptly will be created. Reason: new_mirror\n"+
		"+mirror haproxy will be created. Reason: new_mirror\n"+
		"+repo devenvironment-internal will be created. Reason: new repo\n"+
		"+repo test will be created. Reason: new repo\n"+
		"+snapshot devenvironment will be updated at revision 00001. Reason: update_snapshot\n")
	testExecuteCommand = []string{"bash", "-c", "./aptlify apply -config dockertests/data/TestBaseConfigFile.json"}
	stdout, _ = dockerutils.RunCommand(client, id, testExecuteCommand...)
	c.Check(stdout, Equals, ""+
		"\x01\x00\x00\x00\x00\x00\x00(gpg 9E3E53F19C7DE460 creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00 gpg 353525F9 creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00(gpg 505D97A41C61B9CD creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00 gpg 1C61B9CD creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00-mirror devenvironment-aptly create succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00 mirror haproxy create succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x000repo devenvironment-internal creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x00\x1drepo test creation succeeded\n"+
		"\x01\x00\x00\x00\x00\x00\x007combined snapshot created devenvironment at revision 1\n")
}
