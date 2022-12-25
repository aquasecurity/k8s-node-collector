SHELL := /bin/bash

GOCMD=go
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test


all:
	$(info  "completed running make file for go-opa-validate")
fmt:
	@go fmt ./...
lint:
	./lint.sh
tidy:
	$(GOMOD) tidy -v
test:
	$(GOTEST) ./... 

.PHONY: install-req fmt lint tidy test imports .