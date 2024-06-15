build-linux:
    GOOS=linux GOARCH=amd64 go build -o $(BINARY)-$(GOOS)-$(GOARCH) ./cmd

build-osx:
    GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-$(GOOS)-$(GOARCH) ./cmd

build-windows:
	go build -o app.exe ./cmd

wire:
	cd cmd && wire

run:
	go run ./cmd