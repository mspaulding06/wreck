// Copyright 2017 Matt Spaulding.  All rights reserved.

// Package wreck implements the wreck command line tool
package wreck

import "testing"

func TestGetCommand(t *testing.T) {
	content, err := Get("https://httpbin.org/get")
	if err != nil {
		t.Fail()
	}
	if content == "" {
		t.Fail()
	}
}
