package utils

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestStrings(t *testing.T) { TestingT(t) }

type StringSuite struct{}

var _ = Suite(&StringSuite{})

func (s *StringSuite) TestSplitStringToSliceNoCR(c *C) {
	outSlice := SplitStringToSlice("hello world")
	c.Check(len(outSlice), Equals, 1)
	c.Check(outSlice[0], Equals, "hello world")
}

func (s *StringSuite) TestSplitStringToSliceNoTrailingCR(c *C) {
	outSlice := SplitStringToSlice("hello world\nhow are you?")
	c.Check(len(outSlice), Equals, 2)
	c.Check(outSlice[0], Equals, "hello world")
	c.Check(outSlice[1], Equals, "how are you?")
}

func (s *StringSuite) TestSplitStringToSliceWithTrailingCR(c *C) {
	outSlice := SplitStringToSlice("hello world\nhow are you?\n")
	c.Check(len(outSlice), Equals, 2)
	c.Check(outSlice[0], Equals, "hello world")
	c.Check(outSlice[1], Equals, "how are you?")
}

func (s *StringSuite) TestSplitStringToSliceWithMultipleCRs(c *C) {
	outSlice := SplitStringToSlice("hello world\nhow are you?\nyou are fine\ngood to hear\n")
	c.Check(len(outSlice), Equals, 4)
	c.Check(outSlice[0], Equals, "hello world")
	c.Check(outSlice[1], Equals, "how are you?")
	c.Check(outSlice[2], Equals, "you are fine")
	c.Check(outSlice[3], Equals, "good to hear")
}

func (s *StringSuite) TestIsStringEmptyWithEmptyString(c *C) {
	b := IsStringEmpty("")
	c.Assert(b, Equals, true)
}

func (s *StringSuite) TestIsStringEmptyWithNonEmptyString(c *C) {
	b := IsStringEmpty("I'm not empty")
	c.Assert(b, Equals, false)
}

func (s *StringSuite) TestUniqueStringsOnLeftWithNoDiff(c *C) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	out, err := uniqueStringsOnLeft(a, b)
	c.Assert(len(out), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestUniqueStringsOnLeftWithSingleDiff(c *C) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"a", "b", "c"}
	out, err := uniqueStringsOnLeft(a, b)
	c.Assert(len(out), Equals, 1)
	c.Assert(out[0], Equals, "d")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestUniqueStringsOnLeftWithMultipleDiff(c *C) {
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	b := []string{"a", "c", "k"}
	out, err := uniqueStringsOnLeft(a, b)
	c.Assert(len(out), Equals, 5)
	c.Assert(out[0], Equals, "b")
	c.Assert(out[1], Equals, "d")
	c.Assert(out[2], Equals, "e")
	c.Assert(out[3], Equals, "f")
	c.Assert(out[4], Equals, "g")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestUniqueStringsOnLeftWithDiffOnRight(c *C) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c", "d"}
	out, err := uniqueStringsOnLeft(a, b)
	c.Assert(len(out), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithNoDiff(c *C) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithOneDiffLeft(c *C) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"a", "b", "c"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 1)
	c.Assert(len(out2), Equals, 0)
	c.Assert(out1[0], Equals, "d")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithOneDiffRight(c *C) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c", "d"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 1)
	c.Assert(out2[0], Equals, "d")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithMultipleDiffLeft(c *C) {
	a := []string{"a", "b", "c", "d", "e", "f"}
	b := []string{"a", "b", "c"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 3)
	c.Assert(len(out2), Equals, 0)
	c.Assert(out1[0], Equals, "d")
	c.Assert(out1[1], Equals, "e")
	c.Assert(out1[2], Equals, "f")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithMultipleDiffRight(c *C) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c", "d", "e", "f"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 3)
	c.Assert(out2[0], Equals, "d")
	c.Assert(out2[1], Equals, "e")
	c.Assert(out2[2], Equals, "f")
	c.Assert(err, Equals, nil)
}

func (s *StringSuite) TestDiffStringSlicesWithMultipleDiffs(c *C) {
	a := []string{"b", "c", "d", "e", "f"}
	b := []string{"a", "b", "c", "e", "g", "h", "i"}
	out1, out2, err := DiffStringSlices(a, b)
	c.Assert(len(out1), Equals, 2)
	c.Assert(len(out2), Equals, 4)
	c.Assert(out1[0], Equals, "d")
	c.Assert(out1[1], Equals, "f")
	c.Assert(out2[0], Equals, "a")
	c.Assert(out2[1], Equals, "g")
	c.Assert(out2[2], Equals, "h")
	c.Assert(out2[3], Equals, "i")
	c.Assert(err, Equals, nil)
}
