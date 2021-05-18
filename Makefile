# ---------------------------------------------------------------------------
# Makefile for CLI utilities
# ---------------------------------------------------------------------------

BUILD_TAG=$(shell git describe --tags 2>/dev/null || echo undefined)
LDFLAGS=-ldflags=all="-X main.version=${BUILD_TAG} -s -w"

all: build

build:
	go build ${LDFLAGS}

test:
	go test -v -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

install:
	go install ${LDFLAGS} ./...

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

release: 
	goreleaser release --rm-dist

dist: clean build
	upx -9 epoch.exe

clean:
	go clean
