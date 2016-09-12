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

func (s *AptlyCliSuite) TestGpgAddSuccess(c *C) {
	a := AptlyCli{}
    testGpgKey := "9E3E53F19C7DE460"
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Gpg_add(testGpgKey)
	c.Check(outstring[0], Equals, "gpg: requesting key 9C7DE460 from hkp server keys.gnupg.net")
    c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestGpgAddFailure(c *C) {
	a := AptlyCli{}
    testGpgKey := "FAKE"
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Gpg_add(testGpgKey)
	c.Check(outstring[0], Equals, "gpg: requesting key FAKE from hkp server keys.gnupg.net")
    c.Check(err, ErrorMatches, "exit status 2")
}

func (s *AptlyCliSuite) TestMirrorListSuccess(c *C) {
	a := AptlyCli{}
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Mirror_list()
	c.Check(outstring[0], Equals, "some_mirror")
    c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorUpdateSuccess(c *C) {
	a := AptlyCli{}
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    testMirror := "working_mirror"
    outstring, err := a.Mirror_update(testMirror)
	c.Check(outstring[0], Equals, "Mirror `working_mirror` has been successfully updated.")
    c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorUpdateFailure(c *C) {
	a := AptlyCli{}
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    testMirror := "test_mirror_no_exist"
    outstring, err := a.Mirror_update(testMirror)
	c.Check(outstring[0], Equals, "ERROR: unable to update: mirror with name test_mirror_no_exist not found")
    c.Check(err, ErrorMatches, "exit status 1")
}

func (s *AptlyCliSuite) TestMirrorCreateNofilter(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    testData.Name = "test_mirror"
    testData.Url = "http://example.com"
    testData.Dist = "test_dist"
    testData.Component = "test_component"
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Mirror_create(testData)
	c.Check(outstring[4], Equals, "Mirror [test_mirror]: http://example.com test_dist successfully added.")
    c.Check(err, Equals, nil)
}
func (s *AptlyCliSuite) TestMirrorCreateSinglefilter(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    var testFilters = []mirror.AptlyFilterStruct {
        mirror.AptlyFilterStruct {
            Name: "software1",
            Version: "1.2.3",
        },
    }
    testData.Name = "test_mirror"
    testData.Url = "http://example.com"
    testData.Dist = "test_dist"
    testData.Filter = testFilters
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Mirror_create(testData)
	c.Check(outstring[4], Equals, "Mirror [test_mirror]: http://example.com test_dist successfully added.")
    c.Check(err, Equals, nil)
}
func (s *AptlyCliSuite) TestMirrorCreateMultifilter(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    var testFilters = []mirror.AptlyFilterStruct {
        mirror.AptlyFilterStruct {
            Name: "software1",
            Version: "1.2.3",
        },
        {
            Name: "software2",
            Version: "1.2.3",
        },
    }
    testData.Name = "test_mirror"
    testData.Url = "http://example.com"
    testData.Dist = "test_dist"
    testData.Component = "test_component"
    testData.Filter = testFilters
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    outstring, err := a.Mirror_create(testData)
	c.Check(outstring[4], Equals, "Mirror [test_mirror]: http://example.com test_dist successfully added.")
    c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorCreateMissingName(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    testData.Url = "http://example.com"
    testData.Dist = "test_dist"
    testData.Component = "test_component"
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    _, err := a.Mirror_create(testData)
    c.Check(err, ErrorMatches, "Missing name from mirror")
}
func (s *AptlyCliSuite) TestMirrorCreateMissingUrl(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    testData.Name = "test_name"
    testData.Dist = "test_dist"
    testData.Component = "test_component"
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    _, err := a.Mirror_create(testData)
    c.Check(err, ErrorMatches, "Missing url from mirror")
}
func (s *AptlyCliSuite) TestMirrorCreateMissingDist(c *C) {
	a := AptlyCli{}
    testData := mirror.AptlyMirrorStruct{}
    testData.Name = "test_name"
    testData.Url = "http://example.com"
    testData.Component = "test_component"
    testData.FilterDeps = true
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    _, err := a.Mirror_create(testData)
    c.Check(err, ErrorMatches, "Missing distribution from mirror")
}
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

func (s *AptlyCliSuite) TestSnapShotCreateFailure(c *C) {

	a := AptlyCli{}
	testResource := snapshot.ResourceStruct{}
	testResource.Name = "test_mirror_fail"
	testResource.Type = "mirror"
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    //Fake time
    timestamp = fakeTimestamp
    defer func(){ timestamp = realTimestamp }()
	outstring, err, snapname := a.SnapshotCreate(testResource)
	c.Check(outstring[0], Equals, "ERROR: unable to create snapshot: mirror with name test_mirror_fail not found")
	c.Check(snapname, Equals, "test_mirror_fail_1970-01-01_00:00:00")
	c.Assert(err, ErrorMatches, "exit status 1")
}

func (s *AptlyCliSuite) TestSnapShotCreateUpdateFail(c *C) {

	a := AptlyCli{}
	testResource := snapshot.ResourceStruct{}
	testResource.Name = "test_mirror_no_exist"
	testResource.Type = "mirror"
    //Fake exec
    execExec = fakeExecExec
    defer func(){ execExec = exec.Exec }()
    //Fake time
    timestamp = fakeTimestamp
    defer func(){ timestamp = realTimestamp }()
	outstring, err, snapname := a.SnapshotCreate(testResource)
	c.Check(outstring[0], Equals, "ERROR: unable to update: mirror with name test_mirror_no_exist not found")
	c.Check(snapname, Equals, "test_mirror_no_exist_1970-01-01_00:00:00")
	c.Assert(err, ErrorMatches, "exit status 1")
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
        "aptly snapshot create test_mirror_fail_1970-01-01_00:00:00 from mirror test_mirror_fail": `ERROR: unable to create snapshot: mirror with name test_mirror_fail not found`,
        "aptly mirror update test_mirror_no_exist": `ERROR: unable to update: mirror with name test_mirror_no_exist not found`,
        "gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys 9E3E53F19C7DE460": `gpg: requesting key 9C7DE460 from hkp server keys.gnupg.net`,
        "gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys FAKE": `gpg: requesting key FAKE from hkp server keys.gnupg.net`,
        "aptly mirror list -raw": `some_mirror`,
        "aptly mirror update broken_mirror": `ERROR: unable to update: mirror with name broken_mirror not found`,
        "aptly mirror update working_mirror": "Mirror `working_mirror` has been successfully updated.",
        "aptly mirror create -filter=( Name (= software1 ) , $Version (= 1.2.3 ) ) | ( Name (= software2 ) , $Version (= 1.2.3 ) ) -filter-with-deps test_mirror http://example.com test_dist test_component": `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
        "aptly mirror create -filter=( Name (= software1 ) , $Version (= 1.2.3 ) ) -filter-with-deps test_mirror http://example.com test_dist":  `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
        "aptly mirror create -filter-with-deps test_mirror http://example.com test_dist test_component":  `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
    }
    testError:= map[string]int{
        "aptly mirror update test_mirror_no_exist": 1,
        "aptly snapshot merge testCombinedSnapshot input1 input_no_exist": 1,
        "aptly snapshot create test_mirror_fail_1970-01-01_00:00:00 from mirror test_mirror_fail": 1,
        "gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys FAKE": 2,
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
