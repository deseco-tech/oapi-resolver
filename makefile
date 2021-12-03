## install: install oapi-resolver
.PHONY: install
install:
	@go install
	@echo "opai-resolver installed"

## precommit/install: install pre-commit config file from {root}/.pre-commit-config.yaml
.PHONY: precommit/install
precommit/install:
	@pre-commit install

## tests/run: run all tests
.PHONY: tests/run
tests/run:
	@go test -v -failfast -vet=off ./...

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
