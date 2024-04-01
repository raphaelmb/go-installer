test:
	@go test -v ./...

run:
	@go run cmd/main.go

build:
	@go build -o go-install ./cmd/main.go

.PHONY: run