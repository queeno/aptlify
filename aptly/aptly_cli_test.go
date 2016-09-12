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

func (s *AptlyCliSuite) TestSnapShotCreateSuccess(c *C) {

	a := AptlyCli{}
	testResource := snapshot.ResourceStruct{}
	testResource.Name = "test_mirror"
	testResource.Type = "mirror"
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    //Fake time
    timestamp = fakeTimestamp
    defer func(){ timestamp = realTimestamp }()
	outstring, err, snapname := a.SnapshotCreate(testResource)
	c.Check(outstring[1], Equals, "Snapshot test_mirror_1970-01-01_00:00:00 successfully created.")
	c.Check(snapname, Equals, "test_mirror_1970-01-01_00:00:00")
	c.Assert(err, Equals, nil)
}

func (s *AptlyCliSuite) TestSnapShotMergeSuccess(c *C) {
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
func (s *AptlyCliSuite) TestSnapShotMergeFailure(c *C) {
    a := AptlyCli {}
    combinedName := "testCombinedSnapshot"
    inputSnapshotNames := []string{"input1", "input_no_exist"}
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.SnapshotMerge(combinedName, inputSnapshotNames)
    c.Check(outstring[1], Equals, "ERROR: unable to load snapshot: snapshot with name input1 not found")
	c.Assert(err, ErrorMatches, "exit status 1")
}

func fakeExecExec(command string, args...string) ([]string, error) {
    cs := []string{"-test.run=TestHelperProcess", "--", command}
    cs = append(cs, args...)
    cmd, err := exec.Exec(os.Args[0], cs...)
    return cmd, err
}

func fakeTimestamp() string {
    return "1970-01-01_00:00:00"
}

func TestHelperProcess(t *testing.T){
    testOutput := map[string]string{
        "aptly snapshot merge testCombinedSnapshot input1 input2": `
Snapshot testCombinedSnapshot successfully created.
You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.`,
        "aptly snapshot merge testCombinedSnapshot input1 input_no_exist": `
ERROR: unable to load snapshot: snapshot with name input1 not found`,
        "aptly snapshot create from mirror test_mirror": `
Snapshot testCombinedSnapshot successfully created.
You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.`,
        "aptly snapshot create test_mirror_1970-01-01_00:00:00 from mirror test_mirror": `
Snapshot test_mirror_1970-01-01_00:00:00 successfully created.
You can run 'aptly publish snapshot test_mirror_1970-01-01_00:00:00' to publish snapshot as Debian repository.`,
    }
    testError:= map[string]int{
        "aptly snapshot merge testCombinedSnapshot input1 input_no_exist": 1,
    }
    commandString := strings.Join(os.Args[3:], " ")
    if testOutput[commandString] == "" {
        fmt.Fprintf(os.Stdout, commandString)
    }
    fmt.Fprintf(os.Stdout, testOutput[commandString])
    if testError[commandString] != 0 {
        os.Exit(testError[commandString])
    }
    os.Exit(0)
}
