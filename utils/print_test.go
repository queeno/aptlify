package utils

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestPrint(t *testing.T) { TestingT(t) }

type PrintSuite struct{}

var _ = Suite(&PrintSuite{})

func ExamplePrintEmptySlice() {
	testStringArray := []string{}
	PrintSlice(testStringArray)
	// Output:
}

func ExamplePrintSliceLengthOne() {
	testStringArray := []string{"hello"}
	PrintSlice(testStringArray)
	// Output: - hello
}

func ExamplePrintSliceLengthTwo() {
	testStringArray := []string{"hello", "world"}
	PrintSlice(testStringArray)
	// Output:   - hello
	//   - world
}
