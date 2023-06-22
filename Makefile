.PHONY: build

VERSION := $(shell git describe --tags --always)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -R)
REPO_NAME := "github.com/pokeyaro/gocli"
BUILD_NAME := "gocli"

define LDFLAGS
"-X '${REPO_NAME}/config.Version=${VERSION}' \
-X '${REPO_NAME}/config.GitCommit=${GIT_COMMIT}' \
-X '${REPO_NAME}/config.BuildTime=${BUILD_TIME}'"
endef

build:
	go build -ldflags ${LDFLAGS} -o ${BUILD_NAME} .
