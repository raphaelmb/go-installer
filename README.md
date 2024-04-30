# Go Installer

This project installs [Go](https://go.dev/) on Linux. If it is already installed, it upgrades to the latest version.

## Steps

Download the binary on the [releases](https://github.com/raphaelmb/go-installer/releases) page.

Grant execute permission to the binary `chmod +x go-installer`

Run `./go-installer`

## Build the project yourself

Run `make build` (or `go build -o bin/go-installer ./cmd/main.go`) on the root directory. The binary will be located at `bin/` in the same directory.
