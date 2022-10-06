.DEFAULT_GOAL := default
.PHONY: default test

GO_OS := "linux"
GO_ARCH := "amd64"

default: build test

lint:
	@golangci-lint run ./...

build: build-game

build-game:
	@GOOS=${GO_OS} GOARCH=${GO_ARCH} CGO_ENABLED=0 go build -o bin/game \
		./cmd/game

test:
	@go test -short -race -count=1 -v ./...
