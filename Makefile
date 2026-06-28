APP_NAME=ripper

build:
	go build -o bin/$(APP_NAME) ./cmd/ripper

install:
	go install ./cmd/ripper/

run:
	go run ./cmd/ripper

test:
	go test ./...

clean:
	rm -rf bin