package utils

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestPrint(t *testing.T) { TestingT(t) }

type PrintSuite struct{}

var _ = Suite(&PrintSuite{})

func ExamplePrintEmptySlice() {
	test_string_array := []string{}
	PrintSlice(test_string_array)
	// Output:
}

func ExamplePrintSliceLengthOne() {
	test_string_array := []string{"hello"}
	PrintSlice(test_string_array)
	// Output: - hello
}

func ExamplePrintSliceLengthTwo() {
	test_string_array := []string{"hello", "world"}
	PrintSlice(test_string_array)
	// Output:   - hello
	//   - world
}
