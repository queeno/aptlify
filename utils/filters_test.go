package utils

import (
	"github.com/queeno/aptlify/config"
	. "gopkg.in/check.v1"
	"testing"
)

func TestFilters(t *testing.T) { TestingT(t) }

type FilterSuite struct{}

var _ = Suite(&FilterSuite{})

func (s *FilterSuite) TestUniqueFiltersOnLeftWithNoDiff(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out, err := uniqueFiltersOnLeft(a, b)
	c.Assert(len(out), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestUniqueFiltersOnLeftWithSingleDiff(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"z", "1.0.2"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out, err := uniqueFiltersOnLeft(a, b)
	c.Assert(len(out), Equals, 1)
	c.Assert(out[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.2"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestUniqueFiltersOnLeftWithMultipleDiff(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"z", "1.0.2"})
	a = append(a, config.AptlyFilterStruct{"x", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"a", "1.0.2"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out, err := uniqueFiltersOnLeft(a, b)
	c.Assert(len(out), Equals, 3)
	c.Assert(out[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.2"})
	c.Assert(out[1], Equals, config.AptlyFilterStruct{Name: "x", Version: "1.0.1"})
	c.Assert(out[2], Equals, config.AptlyFilterStruct{Name: "a", Version: "1.0.2"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestUniqueFiltersOnLeftWithDiffOnRight(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"z", "1.0.2"})
	out, err := uniqueFiltersOnLeft(a, b)
	c.Assert(len(out), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithNoDiff(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 0)
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithOneDiffLeft(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"z", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 1)
	c.Assert(len(out2), Equals, 0)
	c.Assert(out1[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.1"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithOneDiffRight(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"z", "1.0.1"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 1)
	c.Assert(out2[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.1"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithMultipleDiffLeft(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"z", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"k", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"p", "1.0.7"})
	a = append(a, config.AptlyFilterStruct{"a", "1.1.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 4)
	c.Assert(len(out2), Equals, 0)
	c.Assert(out1[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.1"})
	c.Assert(out1[1], Equals, config.AptlyFilterStruct{Name: "k", Version: "1.0.0"})
	c.Assert(out1[2], Equals, config.AptlyFilterStruct{Name: "p", Version: "1.0.7"})
	c.Assert(out1[3], Equals, config.AptlyFilterStruct{Name: "a", Version: "1.1.1"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithMultipleDiffRight(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"z", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"k", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"p", "1.0.7"})
	b = append(b, config.AptlyFilterStruct{"a", "1.1.1"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 0)
	c.Assert(len(out2), Equals, 4)
	c.Assert(out2[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.1"})
	c.Assert(out2[1], Equals, config.AptlyFilterStruct{Name: "k", Version: "1.0.0"})
	c.Assert(out2[2], Equals, config.AptlyFilterStruct{Name: "p", Version: "1.0.7"})
	c.Assert(out2[3], Equals, config.AptlyFilterStruct{Name: "a", Version: "1.1.1"})
	c.Assert(err, Equals, nil)
}

func (s *FilterSuite) TestDiffFilterSlicesWithMultipleDiffs(c *C) {
	a := []config.AptlyFilterStruct{}
	b := []config.AptlyFilterStruct{}
	a = append(a, config.AptlyFilterStruct{"x", "1.0.0"})
	a = append(a, config.AptlyFilterStruct{"y", "1.0.1"})
	a = append(a, config.AptlyFilterStruct{"p", "1.1.7"})
	a = append(a, config.AptlyFilterStruct{"q", "2.0.0"})
	a = append(a, config.AptlyFilterStruct{"r", "1.1.7"})
	a = append(a, config.AptlyFilterStruct{"s", "1.0.7"})
	b = append(b, config.AptlyFilterStruct{"x", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"y", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"z", "1.0.1"})
	b = append(b, config.AptlyFilterStruct{"k", "1.0.0"})
	b = append(b, config.AptlyFilterStruct{"p", "1.1.7"})
	out1, out2, err := DiffFilterSlices(a, b)
	c.Assert(len(out1), Equals, 3)
	c.Assert(len(out2), Equals, 2)
	c.Assert(out1[0], Equals, config.AptlyFilterStruct{Name: "q", Version: "2.0.0"})
	c.Assert(out1[1], Equals, config.AptlyFilterStruct{Name: "r", Version: "1.1.7"})
	c.Assert(out1[2], Equals, config.AptlyFilterStruct{Name: "s", Version: "1.0.7"})
	c.Assert(out2[0], Equals, config.AptlyFilterStruct{Name: "z", Version: "1.0.1"})
	c.Assert(out2[1], Equals, config.AptlyFilterStruct{Name: "k", Version: "1.0.0"})
	c.Assert(err, Equals, nil)
}
