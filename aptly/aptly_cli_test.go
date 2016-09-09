package aptly

import (
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
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
