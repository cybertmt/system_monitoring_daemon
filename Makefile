BIN := "./bin/daemon"
BIN_CLIENT := "./bin/client"
DOCKER_IMG="system_monitoring_daemon:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build_daemon:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/daemon

build_client:
	go build -v -o $(BIN_CLIENT) -ldflags "$(LDFLAGS)" ./cmd/client

run_daemon: build_daemon
	$(BIN)

run_client: build_client
	$(BIN_CLIENT)

test:
	go test -race -count 100 ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

generate:
	go generate ./...
