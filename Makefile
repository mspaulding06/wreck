# Copyright 2017 Matt Spaulding.  All rights reserved.

.PHONY: all build install update

ifdef ($(GOPATH),)
$(error "Must have GOPATH set!")
endif

ifdef ($(GOROOT),)
$(error "Must have GOROOT set!")
endif

all: install

build:
	@go build ./cmd/...

install:
	@go install ./cmd/...

test:
	@go test

update:
	@glide update