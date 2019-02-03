package jst_go

import (
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	actual := JstNow()
	if expected, _ := time.LoadLocation("Asia/Tokyo"); actual.Location() != expected {
		t.Fatal("failed test")
	}
}