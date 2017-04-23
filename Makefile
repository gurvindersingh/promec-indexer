GO_EXECUTABLE ?= go
PACKAGE_DIRS := $(shell glide nv)
VERSION := $(shell git describe --tags --dirty --always)
DIST_DIRS := find * -type d -exec

all: test build

build:
	mkdir -p dist/
	GOOS=linux GOARCH=amd64 ${GO_EXECUTABLE} build -o dist/promec-indexer-linux-amd64 -ldflags "-X main.version=${VERSION}"
	GOOS=darwin GOARCH=amd64 ${GO_EXECUTABLE} build -o dist/promec-indexer-darwin-amd64 -ldflags "-X main.version=${VERSION}"
	GOOS=windows GOARCH=amd64 ${GO_EXECUTABLE} build -o dist/promec-indexer-windows-amd64.exe -ldflags "-X main.version=${VERSION}"

test:
	${GO_EXECUTABLE} test --short $(PACKAGE_DIRS)

clean:
	rm -f ./promec-indexer.test
	rm -f ./promec-indexer
	rm -rf ./dist
