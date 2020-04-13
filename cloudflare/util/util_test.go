package util

import (
	"testing"
	"time"
)

func TestGetTimestamp(t *testing.T) {
	expectedFormat := "2020-04-11T16:37:06.498+03:00"
	testTime := time.Date(2020, 04, 11, 16, 37, 06, 498*1000000, time.FixedZone("+3", 3*60*60))
	testFormat := getTimestamp(testTime)
	if testFormat != expectedFormat {
		t.Error("Invalid timestamp")
	}
}
