test:
	@go test -v ./...

run:
	@go run cmd/main.go

build:
	@go build -o bin/go-installer ./cmd/main.go

.PHONY: run test build
