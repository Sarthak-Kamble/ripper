APP_NAME=ripper
VERSION ?= dev

build:
	go build -ldflags "-X github.com/Sarthak-Kamble/ripper/internal/cli.Version=$(VERSION)" -o bin/$(APP_NAME).exe ./cmd/ripper

install:
	go install ./cmd/ripper

run:
	go run ./cmd/ripper

test:
	go test ./...

clean:
	rmdir /s /q bin