CLI_NAME := go-expect-cli-example

build:
	go build \
		-o "./$(CLI_NAME)" \
		./main.go

test:
	go test

.DEFAULT_GOAL := build
