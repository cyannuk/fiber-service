wire:
	wire gen ./composition

generate:
	go generate ./domain/model

build:
	go build -o ./bin/ -v -ldflags="-w -s" ./main/

dependencies:
	go mod download

.PHONY: build dependencies wire generate
.DEFAULT_GOAL := build
