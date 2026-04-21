.PHONY: build test run lint

build:
	go build -o bin/claw-device-pairing cmd/main.go

test:
	go test ./...

run:
	go run cmd/main.go serve

lint:
	go vet ./...
