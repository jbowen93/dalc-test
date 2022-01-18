#!/usr/bin/make -f

PACKAGES=$(shell go list ./...)
BUILDDIR ?= $(CURDIR)/build

## build: Build dalc-test binary.
build:
	@echo "--> Building DALC-test"
	@go build .
.PHONY: build
