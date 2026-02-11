SHELL := /bin/bash

GOCMD=go
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test

all:
	$(info  "completed running make file for k8s node collector")
fmt:
	@go fmt ./...
tidy:
	$(GOMOD) tidy -v
test:
	$(GOTEST) ./... 

build:
	go build -o node-collector ./cmd/node-collector/

build-docker:
	docker build -t ghcr.io/aquasecurity/node-collector:dev -f ./build/node-collector/Dockerfile .

.PHONY: install-req fmt lint tidy test imports build build-docker
