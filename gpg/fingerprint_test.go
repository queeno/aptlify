package gpg

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestGpg(t *testing.T) { TestingT(t) }

type GpgSuite struct{}

var _ = Suite(&GpgSuite{})

func (s *GpgSuite) TestExtractFingerprintsWithGoodText(c *C) {
	keysText := []string{"tru:t:1:1472835012:0:3:1:5",
		"pub:-:1024:17:40976EAF437D05B5:2004-09-12:::-:Ubuntu Archive Automatic Signing Key <ftpmaster@ubuntu.com>::scESC:",
		"fpr:::::::::630239CC130E1A7FD81A27B140976EAF437D05B5:",
		"sub:-:2048:16:251BEFF479164387:2004-09-12::::::e:",
		"pub:-:4096:1:3B4FE6ACC0B21F32:2012-05-11:::-:Ubuntu Archive Automatic Signing Key (2012) <ftpmaster@ubuntu.com>::scSC:",
		"fpr:::::::::790BC7277767219C42C86F933B4FE6ACC0B21F32:",
		"pub:-:4096:1:D94AA3F0EFE21092:2012-05-11:::-:Ubuntu CD Image Automatic Signing Key (2012) <cdimage@ubuntu.com>::scSC:",
		"fpr:::::::::843938DF228D22F7B3742BC0D94AA3F0EFE21092:",
		"pub:-:1024:17:46181433FBB75451:2004-12-30:::-:Ubuntu CD Image Automatic Signing Key <cdimage@ubuntu.com>::scSC:",
		"fpr:::::::::C5986B4F1257FFA86632CBA746181433FBB75451:",
		"pub:-:4096:1:F76221572C52609D:2015-07-14:::-:Docker Release Tool (releasedocker) <docker@docker.com>::escaESCA:",
		"fpr:::::::::58118E89F3A912897C070ADBF76221572C52609D:",
		"pub:-:2048:1:9E3E53F19C7DE460:2016-03-15:2018-03-15::-:Andrey Smirnov <me@smira.ru>::scESC:",
		"fpr:::::::::DF32BC15E2145B3FA151AED19E3E53F19C7DE460:",
		"sub:-:2048:1:57A48F2F1793CB0C:2016-03-15:2018-03-15:::::e:",
		""}

	fingerprints, err := extractFingerprints(keysText)
	c.Assert(len(fingerprints), Equals, 6)
	c.Assert(fingerprints[0], Equals, "40976EAF437D05B5")
	c.Assert(fingerprints[1], Equals, "3B4FE6ACC0B21F32")
	c.Assert(fingerprints[2], Equals, "D94AA3F0EFE21092")
	c.Assert(fingerprints[3], Equals, "46181433FBB75451")
	c.Assert(fingerprints[4], Equals, "F76221572C52609D")
	c.Assert(fingerprints[5], Equals, "9E3E53F19C7DE460")
	c.Assert(err, Equals, nil)
}

func (s *GpgSuite) TestExtractFingerprintsWithBadText(c *C) {
	keysText := []string{"pr:::::::::630239CC130E1A7FD81A27B140976EAF437D05B5:"}
	fingerprints, err := extractFingerprints(keysText)
	c.Assert(len(fingerprints), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *GpgSuite) TestExtractFingerprintsWithTruncatedText(c *C) {
	keysText := []string{"fpr:::::::::9CC130E1A7FD81A27B140976EAF437D05B5:"}
	fingerprints, err := extractFingerprints(keysText)
	c.Assert(len(fingerprints), Equals, 0)
	c.Check(err, ErrorMatches, "malformed gpg-fingerprint returned by apt-key")
}
