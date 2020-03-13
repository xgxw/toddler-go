package tests

import "testing"

var debug = false

func Print(t *testing.T, msg string) {
	if debug {
		t.Logf(msg)
	}
}
