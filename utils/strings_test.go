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
