package gotime

import (
	"testing"
	"time"
)

var (
	testDate    = time.Date(2018, 6, 12, 13, 15, 17, 0, time.UTC)
	testDateStr = "June, Jun, 06, Tuesday, Tue, 12, 13, 15, 17, 2018, 18, pm, pm"
	testDateFmt = "%B, %b, %m, %A, %a, %d, %H, %M, %S, %Y, %y, %p, %p"
)

// TestFormatDateTime is a test for FormatDateTime
func TestFormatDateTime(t *testing.T) {
	if testDateStr != FormatDateTime(testDate, testDateFmt) {
		t.Fatalf("Format did not match output string.\n Expected %s received %s", testDateStr, FormatDateTime(testDate, testDateFmt))
	}
}

// BenchmarkSetValues is a benchmark for setValues
func BenchmarkSetValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setValues()
	}
}

// BenchmarkFormatDateTime is a benchmark for FormatDateTime
func BenchmarkFormatDateTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FormatDateTime(testDate, testDateFmt)
	}
}
