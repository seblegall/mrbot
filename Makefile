# set default shell
SHELL := $(shell which bash)
OSARCH := "linux/amd64 darwin/amd64"
ENV = /usr/bin/env

.SHELLFLAGS = -c

.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: all
.DEFAULT: build

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


dep: ## Get build dependencies
	  go get -v -u github.com/golang/dep/cmd/dep && \
      go get github.com/mitchellh/gox && \
      go get github.com/mattn/goveralls

build: ## Build blackbeard
	dep ensure && go build

cross-build: ## Build blackbeard for multiple os/arch
	gox -osarch=$(OSARCH) -output "bin/mrbot_{{.OS}}_{{.Arch}}"

test: ## Launch tests
	go test -v ./...