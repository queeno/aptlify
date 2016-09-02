package aptly

import (
	"github.com/queeno/aptlify/config"
	. "gopkg.in/check.v1"
	"testing"
)

func TestAptlyCli(t *testing.T) { TestingT(t) }

type AptlyCliSuite struct{}

var _ = Suite(&AptlyCliSuite{})

func (s *AptlyCliSuite) TestCreateAptlyMirrorFilterCommand(c *C) {

	testFilter := config.AptlyFilterStruct{}
	testFilter.Name = "package"
	testFilter.Version = "1.0.0"

	testCommand := createAptlyMirrorFilterCommand(testFilter)
	c.Check(testCommand, Equals, "( Name (= package) , $Version (= 1.0.0) )")
}
