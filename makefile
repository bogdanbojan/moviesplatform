SHELL := /bin/bash

run:
	go run ./cmd
build:
	go build ./cmd

# ==============================================================================
# Building the container and running it

VERSION := 1.0
DISTRO := alpine

HOST_PORT := 4000
CONTAINER_PORT := 4000

all: build-docker run-docker

build-docker:
	docker build -t moviesplatform-$(DISTRO):$(VERSION) .

run-docker:
	docker run -p $(HOST_PORT):$(CONTAINER_PORT) moviesplatform-$(DISTRO):$(VERSION)
