// Copyright 2017 Matt Spaulding.  All rights reserved.

// Package wreck implements the wreck command line tool
package wreck

import "testing"

func TestPostCommand(t *testing.T) {
	content, err := Post("https://httpbin.org/post")
	if err != nil {
		t.Fail()
	}
	if content == "" {
		t.Fail()
	}
}
