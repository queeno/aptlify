package utils

import (
	"bytes"
	. "gopkg.in/check.v1"
	"testing"
)

var testLogger = NewLogging()

func captureOutput(f func()) string {
	var logBuffer bytes.Buffer
	testLogger.Trace.SetOutput(&logBuffer)
	testLogger.Info.SetOutput(&logBuffer)
	testLogger.Warning.SetOutput(&logBuffer)
	testLogger.Error.SetOutput(&logBuffer)
	testLogger.Fatal.SetOutput(&logBuffer)
	f()
	return logBuffer.String()
}

func TestLogging(t *testing.T) { TestingT(t) }

type LoggingSuite struct{}

var _ = Suite(&LoggingSuite{})

func (l *LoggingSuite) TestLoggingFunctions(c *C) {

	out := captureOutput(func() {
		testLogger.Trace.Println("teststring")
	})
	c.Check(out, Matches, "TRACE:.*..../../.. ..:..:.. .*:.*: teststring\n")

	out = captureOutput(func() {
		testLogger.Info.Println("teststring")
	})
	c.Check(out, Matches, "INFO:.*..../../.. ..:..:.. .*:.*: teststring\n")

	out = captureOutput(func() {
		testLogger.Warning.Println("teststring")
	})
	c.Check(out, Matches, "WARNING:.*..../../.. ..:..:.. .*:.*: teststring\n")

	out = captureOutput(func() {
		testLogger.Error.Println("teststring")
	})
	c.Check(out, Matches, "ERROR:.*..../../.. ..:..:.. .*:.*: teststring\n")

	out = captureOutput(func() {
		testLogger.Fatal.Println("teststring")
	})
	c.Check(out, Matches, "FATAL:.*..../../.. ..:..:.. .*:.*: teststring\n")

}
