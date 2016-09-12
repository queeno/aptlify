package aptly

import (
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
	"github.com/queeno/aptlify/exec"
    "os"
    "fmt"
    "strings"
	. "gopkg.in/check.v1"
	"testing"
)

func TestAptlyCli(t *testing.T) { TestingT(t) }

type AptlyCliSuite struct{}

var _ = Suite(&AptlyCliSuite{})

func (s *AptlyCliSuite) TestCreateAptlyMirrorFilterCommand(c *C) {

	testFilter := mirror.AptlyFilterStruct{}
	testFilter.Name = "package"
	testFilter.Version = "1.0.0"

	testCommand := createAptlyMirrorFilterCommand(testFilter)
	c.Check(testCommand, Equals, "( Name (= package ) , $Version (= 1.0.0 ) )")
}

func (s *AptlyCliSuite) TestSnapShotCreate(c *C) {

	a := AptlyCli{}
	testResource := snapshot.ResourceStruct{}
	testResource.Name = "NNNNNNNotAResource1092340987213"
	testResource.Type = "mirror"
	outstring, err, snapname := a.SnapshotCreate(testResource)
	c.Check(outstring[0], Equals, "ERROR: unable to update: mirror with name NNNNNNNotAResource1092340987213 not found")
	c.Check(snapname, Matches, "NNNNNNNotAResource1092340987213_....-..-.._..:..:..")
	c.Assert(err, ErrorMatches, "exit status 1")
}

func (s *AptlyCliSuite) TestSnapShotMerge(c *C) {
    a := AptlyCli {}
    combinedName := "testCombinedSnapshot"
    inputSnapshotNames := []string{"input1", "input2"}
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.SnapshotMerge(combinedName, inputSnapshotNames)
    c.Check(outstring[1], Equals, "Snapshot testCombinedSnapshot successfully created.")
    c.Check(outstring[2], Equals, "You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.")
    c.Assert(err, Equals, nil)
}

func fakeExecExec(command string, args...string) ([]string, error) {
    cs := []string{"-test.run=TestHelperProcess", "--", command}
    cs = append(cs, args...)
    cmd, err := exec.Exec(os.Args[0], cs...)
    return cmd, err
}

func TestHelperProcess(t *testing.T){
    testOutput := map[string]string{
        "aptly snapshot merge testCombinedSnapshot input1 input2": `
Snapshot testCombinedSnapshot successfully created.
You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.`,
    }
    testError:= map[string]int{
        "placeholder failing command - replace when": 2,
    }
    commandString := strings.Join(os.Args[3:], " ")
    fmt.Fprintf(os.Stdout, testOutput[commandString])
    if testError[commandString] != 0 {
        os.Exit(testError[commandString])
    }
    os.Exit(0)
}
