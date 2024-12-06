test:
	@go test -v ./...

run:
	@go run cmd/main.go

build:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/go-installer ./cmd/main.go

.PHONY: run test build
