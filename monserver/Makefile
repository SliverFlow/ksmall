BINARY=app
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

build-linux:
    GOOS=linux GOARCH=amd64 go build -o $(BINARY)-$(GOOS)-$(GOARCH) ./cmd

build-osx:
    GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-$(GOOS)-$(GOARCH) ./cmd

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY)-$(GOOS)-$(GOARCH).exe ./cmd

wire:
	cd cmd && wire

run:
	go run ./cmd