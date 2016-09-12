package aptly

import (
	"fmt"
	"github.com/queeno/aptlify/exec"
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
	. "gopkg.in/check.v1"
	"os"
	"strings"
	"testing"
)

func TestAptlyCli(t *testing.T) { TestingT(t) }

type AptlyCliSuite struct{}

var _ = Suite(&AptlyCliSuite{})

func (s *AptlyCliSuite) TestCreateAptlyMirrorFilterCommand(c *C) {

	testFilter1 := mirror.AptlyFilterStruct{}
	testFilter1.Name = "package"
	testFilter1.Version = "1.0.0"

	testFilter2 := mirror.AptlyFilterStruct{}
	testFilter2.Version = "1.0.0"

	testFilter3 := mirror.AptlyFilterStruct{}
	testFilter3.Name = "package"

	testCommand1 := createAptlyMirrorFilterCommand(testFilter1)
	testCommand2 := createAptlyMirrorFilterCommand(testFilter2)
	testCommand3 := createAptlyMirrorFilterCommand(testFilter3)

	c.Check(testCommand1, Equals, "( Name (= package ) , $Version (= 1.0.0 ) )")
	c.Check(testCommand2, Equals, "( $Version (= 1.0.0 ) )")
	c.Check(testCommand3, Equals, "( Name (= package ) )")
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

func (s *AptlyCliSuite) TestSnapShotMerge(c *C) {
	a := AptlyCli{}
	combinedName := "testCombinedSnapshot"
	inputSnapshotNames := []string{"input1", "input2"}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.SnapshotMerge(combinedName, inputSnapshotNames)
	c.Check(outstring[1], Equals, "Snapshot testCombinedSnapshot successfully created.")
	c.Check(outstring[2], Equals, "You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.")
	c.Assert(err, Equals, nil)
}

func (s *AptlyCliSuite) TestCleanSlice(c *C) {
	str_a := []string{"a", "b", "c", "d"}
	str_b := []string{"a", " ", "c", ""}
	str_c := []string{"", "b", "c", "d"}
	str_d := []string{" ", "b", "c", "d"}
	str_e := []string{" ", "", "", ""}
	str_f := []string{"", "", "", ""}
	test_a := cleanSlice(str_a)
	test_b := cleanSlice(str_b)
	test_c := cleanSlice(str_c)
	test_d := cleanSlice(str_d)
	test_e := cleanSlice(str_e)
	test_f := cleanSlice(str_f)
	c.Assert(test_a[0], Equals, str_a[0])
	c.Assert(test_a[1], Equals, str_a[1])
	c.Assert(test_a[2], Equals, str_a[2])
	c.Assert(test_a[3], Equals, str_a[3])
	c.Assert(test_b[0], Equals, "a")
	c.Assert(test_b[1], Equals, " ")
	c.Assert(test_b[2], Equals, "c")
	c.Assert(test_c[0], Equals, "b")
	c.Assert(test_c[1], Equals, "c")
	c.Assert(test_c[2], Equals, "d")
	c.Assert(test_d[0], Equals, " ")
	c.Assert(test_d[1], Equals, "b")
	c.Assert(test_d[2], Equals, "c")
	c.Assert(test_d[3], Equals, "d")
	c.Assert(test_e[0], Equals, " ")
	c.Assert(len(test_f), Equals, 0)
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

// Helper functions

func fakeExecExec(command string, args ...string) ([]string, error) {
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
	if len(os.Args) > 2 {
		if os.Args[1] == "-test.run=TestHelperProcess" {
			commandString := strings.Join(os.Args[3:], " ")
			if testOutput[commandString] == "" {
				fmt.Fprintf(os.Stdout, commandString)
			} else {
				fmt.Fprintf(os.Stdout, testOutput[commandString])
			}
			if testError[commandString] != 0 {
				os.Exit(testError[commandString])
			}
			os.Exit(0)
		}
	}
	return
    os.Exit(0)
}
