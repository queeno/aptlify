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
	"time"
)

func TestAptlyCli(t *testing.T) { TestingT(t) }

type AptlyCliSuite struct{}

var _ = Suite(&AptlyCliSuite{})

func (s *AptlyCliSuite) TestGpgAddSuccess(c *C) {
	a := AptlyCli{}
	testGpgKey := "9E3E53F19C7DE460"
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.GpgAdd(testGpgKey)
	c.Check(outstring[0], Equals, "gpg: requesting key 9C7DE460 from hkp server keys.gnupg.net")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestGpgAddFailure(c *C) {
	a := AptlyCli{}
	testGpgKey := "FAKE"
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.GpgAdd(testGpgKey)
	c.Check(outstring[0], Equals, "gpg: requesting key FAKE from hkp server keys.gnupg.net")
	c.Check(err, ErrorMatches, "exit status 2")
}
func (s *AptlyCliSuite) TestMirrorDropSuccess(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	testData := "test_mirror"
	outstring, err := a.MirrorDrop(testData)
	c.Check(outstring[0], Equals, "Mirror `test_mirror` has been removed.")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestRepoCreateSuccess(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	testData := "test_repo"
	outstring, err := a.RepoCreate(testData)
	c.Check(outstring[2], Equals, "Local repo [test_repo] successfully added.")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestRepoListSuccess(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.RepoList()
	c.Check(outstring[0], Equals, "some_mirror")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorListSuccess(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.MirrorList()
	c.Check(outstring[0], Equals, "some_mirror")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorUpdateSuccess(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	testMirror := "working_mirror"
	outstring, err := a.MirrorUpdate(testMirror)
	c.Check(outstring[0], Equals, "Mirror `working_mirror` has been successfully updated.")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestMirrorUpdateFailure(c *C) {
	a := AptlyCli{}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	testMirror := "test_mirror_no_exist"
	outstring, err := a.MirrorUpdate(testMirror)
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
	defer func() { execExec = exec.Exec }()
	outstring, err := a.MirrorCreate(testData)
	c.Check(outstring[4], Equals, "Mirror [test_mirror]: http://example.com test_dist successfully added.")
	c.Check(err, Equals, nil)
}
func (s *AptlyCliSuite) TestMirrorCreateSinglefilter(c *C) {
	a := AptlyCli{}
	testData := mirror.AptlyMirrorStruct{}
	var testFilters = []mirror.AptlyFilterStruct{
		mirror.AptlyFilterStruct{
			Name:    "software1",
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
	defer func() { execExec = exec.Exec }()
	outstring, err := a.MirrorCreate(testData)
	c.Check(outstring[4], Equals, "Mirror [test_mirror]: http://example.com test_dist successfully added.")
	c.Check(err, Equals, nil)
}
func (s *AptlyCliSuite) TestMirrorCreateMultifilter(c *C) {
	a := AptlyCli{}
	testData := mirror.AptlyMirrorStruct{}
	var testFilters = []mirror.AptlyFilterStruct{
		mirror.AptlyFilterStruct{
			Name:    "software1",
			Version: "1.2.3",
		},
		{
			Name:    "software2",
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
	defer func() { execExec = exec.Exec }()
	outstring, err := a.MirrorCreate(testData)
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
	defer func() { execExec = exec.Exec }()
	_, err := a.MirrorCreate(testData)
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
	defer func() { execExec = exec.Exec }()
	_, err := a.MirrorCreate(testData)
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
	defer func() { execExec = exec.Exec }()
	_, err := a.MirrorCreate(testData)
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
	defer func() { execExec = exec.Exec }()
	//Fake time
	timestamp = fakeTimestamp
	defer func() { timestamp = realTimestamp }()
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
	defer func() { execExec = exec.Exec }()
	//Fake time
	timestamp = fakeTimestamp
	defer func() { timestamp = realTimestamp }()
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
	defer func() { execExec = exec.Exec }()
	//Fake time
	timestamp = fakeTimestamp
	defer func() { timestamp = realTimestamp }()
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
	strA := []string{"a", "b", "c", "d"}
	strB := []string{"a", " ", "c", ""}
	strC := []string{"", "b", "c", "d"}
	strD := []string{" ", "b", "c", "d"}
	strE := []string{" ", "", "", ""}
	strF := []string{"", "", "", ""}
	testA := cleanSlice(strA)
	testB := cleanSlice(strB)
	testC := cleanSlice(strC)
	testD := cleanSlice(strD)
	testE := cleanSlice(strE)
	testF := cleanSlice(strF)
	c.Assert(testA[0], Equals, strA[0])
	c.Assert(testA[1], Equals, strA[1])
	c.Assert(testA[2], Equals, strA[2])
	c.Assert(testA[3], Equals, strA[3])
	c.Assert(testB[0], Equals, "a")
	c.Assert(testB[1], Equals, " ")
	c.Assert(testB[2], Equals, "c")
	c.Assert(testC[0], Equals, "b")
	c.Assert(testC[1], Equals, "c")
	c.Assert(testC[2], Equals, "d")
	c.Assert(testD[0], Equals, " ")
	c.Assert(testD[1], Equals, "b")
	c.Assert(testD[2], Equals, "c")
	c.Assert(testD[3], Equals, "d")
	c.Assert(testE[0], Equals, " ")
	c.Assert(len(testF), Equals, 0)
}
func (s *AptlyCliSuite) TestSnapShotMergeFailure(c *C) {
	a := AptlyCli{}
	combinedName := "testCombinedSnapshot"
	inputSnapshotNames := []string{"input1", "input_no_exist"}
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.SnapshotMerge(combinedName, inputSnapshotNames)
	c.Check(outstring[1], Equals, "ERROR: unable to load snapshot: snapshot with name input1 not found")
	c.Assert(err, ErrorMatches, "exit status 1")
}

func (s *AptlyCliSuite) TestSnapshotFilter(c *C) {
	a := AptlyCli{}
	r := snapshot.ResourceStruct{}
	r.Name = "test"
	r.Type = "mirror"
	absentTestBaseSnapName := "absenttestbase"
	presentTestBaseSnapName := "presenttestbase"

	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()

	out, err, outName := a.SnapshotFilter(r, absentTestBaseSnapName)
	c.Assert(out[0], Equals, "ERROR: unable to filter: snapshot with name absenttestbase not found")
	c.Assert(err, ErrorMatches, "exit status 1")
	c.Assert(outName, Equals, "absenttestbase_filtered")

	out, err, outName = a.SnapshotFilter(r, presentTestBaseSnapName)
	c.Assert(out[0], Equals, "Loading packages (19)...")
	c.Assert(err, Equals, nil)
	c.Assert(outName, Equals, "presenttestbase_filtered")

	filt1 := mirror.AptlyFilterStruct{}
	filt1.Name = "package"
	filt1.Version = "1.0.0"

	r.Filter = []mirror.AptlyFilterStruct{}
	r.Filter = append(r.Filter, filt1)

	out, err, outName = a.SnapshotFilter(r, absentTestBaseSnapName)
	c.Assert(out[0], Equals, "ERROR: unable to filter: snapshot with name absenttestbase not found")
	c.Assert(err, ErrorMatches, "exit status 1")
	c.Assert(outName, Equals, "absenttestbase_filtered")

	out, err, outName = a.SnapshotFilter(r, presentTestBaseSnapName)
	c.Assert(out[0], Equals, "Loading packages (19)...")
	c.Assert(err, Equals, nil)
	c.Assert(outName, Equals, "presenttestbase_filtered")

	filt2 := mirror.AptlyFilterStruct{}
	filt2.Name = "another_package"
	filt2.Version = "2.3.4"

	r.Filter = append(r.Filter, filt2)

	out, err, outName = a.SnapshotFilter(r, absentTestBaseSnapName)
	c.Assert(out[0], Equals, "ERROR: unable to filter: snapshot with name absenttestbase not found")
	c.Assert(err, ErrorMatches, "exit status 1")
	c.Assert(outName, Equals, "absenttestbase_filtered")

	out, err, outName = a.SnapshotFilter(r, presentTestBaseSnapName)
	c.Assert(out[0], Equals, "Loading packages (19)...")
	c.Assert(err, Equals, nil)
	c.Assert(outName, Equals, "presenttestbase_filtered")
}

func (s *AptlyCliSuite) TestSnapshotDrop(c *C) {
	a := AptlyCli{}
	presentSnapshotToDrop := "presentTestSnapshotToDrop"
	absentSnapshotToDrop := "absentTestSnapshotToDrop"
	//Fake exec
	execExec = fakeExecExec
	defer func() { execExec = exec.Exec }()
	outstring, err := a.SnapshotDrop(presentSnapshotToDrop, false)
	c.Check(outstring[0], Equals, "Snapshot `presentTestSnapshotToDrop` has been dropped")
	c.Check(err, Equals, nil)
	outstring, err = a.SnapshotDrop(absentSnapshotToDrop, false)
	c.Check(outstring[0], Equals, "ERROR: unable to drop: snapshot with name absentTestSnapshotToDrop not found")
	c.Check(err, Equals, nil)
	outstring, err = a.SnapshotDrop(presentSnapshotToDrop, true)
	c.Check(outstring[0], Equals, "Snapshot `presentTestSnapshotToDrop` has been dropped")
	c.Check(err, Equals, nil)
	outstring, err = a.SnapshotDrop(absentSnapshotToDrop, true)
	c.Check(outstring[0], Equals, "ERROR: unable to drop: snapshot with name absentTestSnapshotToDrop not found")
	c.Check(err, Equals, nil)
}

func (s *AptlyCliSuite) TestRealTimestamp(c *C) {
	testTime := time.Now()
	funcTimeStr := realTimestamp()
	funcTime, err := time.Parse("2006-01-02_15:04:05", funcTimeStr)
	equalEnough := false
	if funcTime == testTime || funcTime.Sub(testTime) < 3*time.Second || testTime.Sub(funcTime) < 3*time.Second {
		equalEnough = true
	}
	c.Assert(equalEnough, Equals, true)
	c.Assert(err, Equals, nil)
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

func TestHelperProcess(t *testing.T) {
	testOutput := map[string]string{
		fmt.Sprintf("%s snapshot merge testCombinedSnapshot input1 input2", aptlyCmd): `
Snapshot testCombinedSnapshot successfully created.
You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot merge testCombinedSnapshot input1 input_no_exist", aptlyCmd): `
ERROR: unable to load snapshot: snapshot with name input1 not found`,
		fmt.Sprintf("%s snapshot create from mirror test_mirror", aptlyCmd): `
Snapshot testCombinedSnapshot successfully created.
You can run 'aptly publish snapshot testCombinedSnapshot' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot create test_mirror_1970-01-01_00:00:00 from mirror test_mirror", aptlyCmd): `
Snapshot test_mirror_1970-01-01_00:00:00 successfully created.
You can run 'aptly publish snapshot test_mirror_1970-01-01_00:00:00' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot create test_mirror_fail_1970-01-01_00:00:00 from mirror test_mirror_fail", aptlyCmd): `ERROR: unable to create snapshot: mirror with name test_mirror_fail not found`,
		fmt.Sprintf("%s mirror update test_mirror_no_exist", aptlyCmd):                                                `ERROR: unable to update: mirror with name test_mirror_no_exist not found`,

		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ", aptlyCmd): "ERROR: unable to filter: snapshot with name absenttestbase not found",
		fmt.Sprintf("%s snapshot filter presenttestbase presenttestbase_filtered ", aptlyCmd): `Loading packages (19)...
Building indexes...

Snapshot presenttestbase successfully filtered.
You can run 'aptly publish snapshot presenttestbase' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) )", aptlyCmd): "ERROR: unable to filter: snapshot with name absenttestbase not found",
		fmt.Sprintf("%s snapshot filter presenttestbase presenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) )", aptlyCmd): `Loading packages (19)...
Building indexes...

Snapshot presenttestbase successfully filtered.
You can run 'aptly publish snapshot presenttestbase' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot filter presenttestbase presenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) ) | ( Name (= another_package ) , $Version (= 2.3.4 ) )", aptlyCmd): `Loading packages (19)...
Building indexes...

Snapshot presenttestbase successfully filtered.
You can run 'aptly publish snapshot presenttestbase' to publish snapshot as Debian repository.`,
		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) ) | ( Name (= another_package ) , $Version (= 2.3.4 ) )", aptlyCmd): "ERROR: unable to filter: snapshot with name absenttestbase not found",

		fmt.Sprintf("%s snapshot drop -force=false presentTestSnapshotToDrop", aptlyCmd):                             "Snapshot `presentTestSnapshotToDrop` has been dropped",
		fmt.Sprintf("%s snapshot drop -force=true presentTestSnapshotToDrop", aptlyCmd):                              "Snapshot `presentTestSnapshotToDrop` has been dropped",
		fmt.Sprintf("%s snapshot drop -force=false absentTestSnapshotToDrop", aptlyCmd):                              "ERROR: unable to drop: snapshot with name absentTestSnapshotToDrop not found",
		fmt.Sprintf("%s snapshot drop -force=true absentTestSnapshotToDrop", aptlyCmd):                               "ERROR: unable to drop: snapshot with name absentTestSnapshotToDrop not found",
		"gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys 9E3E53F19C7DE460": `gpg: requesting key 9C7DE460 from hkp server keys.gnupg.net`,
		"gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys FAKE":             `gpg: requesting key FAKE from hkp server keys.gnupg.net`,

		fmt.Sprintf("%s mirror list -raw", aptlyCmd):             `some_mirror`,
		fmt.Sprintf("%s mirror update broken_mirror", aptlyCmd):  `ERROR: unable to update: mirror with name broken_mirror not found`,
		fmt.Sprintf("%s mirror update working_mirror", aptlyCmd): "Mirror `working_mirror` has been successfully updated.",
		fmt.Sprintf("%s mirror create -filter=( Name (= software1 ) , $Version (= 1.2.3 ) ) | ( Name (= software2 ) , $Version (= 1.2.3 ) ) -filter-with-deps test_mirror http://example.com test_dist test_component", aptlyCmd): `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
		fmt.Sprintf("%s mirror create -filter=( Name (= software1 ) , $Version (= 1.2.3 ) ) -filter-with-deps test_mirror http://example.com test_dist", aptlyCmd): `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
		fmt.Sprintf("%s mirror create -filter-with-deps test_mirror http://example.com test_dist test_component", aptlyCmd): `Downloading http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu/dists/trusty/InRelease...
gpgv: Signature made Fri May 13 12:29:45 2016 UTC using RSA key ID 1C61B9CD
gpgv: Good signature from "Launchpad PPA for Vincent Bernat"

Mirror [test_mirror]: http://example.com test_dist successfully added.
You can run 'aptly mirror update test_mirror' to download repository contents.`,
		fmt.Sprintf("%s repo list -raw", aptlyCmd): "some_mirror",
		fmt.Sprintf("%s repo create test_repo", aptlyCmd): `

Local repo [test_repo] successfully added.
You can run 'aptly repo add test_repo ...' to add packages to repository.`,
		fmt.Sprintf("%s mirror drop test_mirror", aptlyCmd): "Mirror `test_mirror` has been removed.",
	}

	testError := map[string]int{
		"gpg --no-default-keyring --keyring trustedkeys.gpg --keyserver keys.gnupg.net --recv-keys FAKE":                                                                                     2,
		fmt.Sprintf("%s mirror update test_mirror_no_exist", aptlyCmd):                                                                                                                       1,
		fmt.Sprintf("%s snapshot merge testCombinedSnapshot input1 input_no_exist", aptlyCmd):                                                                                                1,
		fmt.Sprintf("%s snapshot create test_mirror_fail_1970-01-01_00:00:00 from mirror test_mirror_fail", aptlyCmd):                                                                        1,
		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ", aptlyCmd):                                                                                                  1,
		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) )", aptlyCmd):                                                       1,
		fmt.Sprintf("%s snapshot filter absenttestbase absenttestbase_filtered ( Name (= package ) , $Version (= 1.0.0 ) ) | ( Name (= another_package ) , $Version (= 2.3.4 ) )", aptlyCmd): 1,
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
}
