GO_VERSION := 1.21.0
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse HEAD)
DATE := $(shell date +%Y-%m-%d.%H:%M:%S)
LDFLAGS := -w -X github.com/wolftsao/hello-api/handlers.hash=$(HASH) -X github.com/wolftsao/hello-api/handlers.tag=$(TAG) -X github.com/wolftsao/hello-api/handlers.date=$(DATE)

.PHONY: install-go init-go

setup: install-go init-go install-lint copy-hooks install-godog

# TODO add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

install-lint:
	sudo curl -sSfL \
	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
	| sh -s -- -b $$(go env GOPATH)/bin v1.54.1

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.

install-godog:
	go install github.com/cucumber/godog/cmd/godog@latest

static-check:
	golangci-lint run

build:
	go build -ldflags "$(LDFLAGS)" -o api cmd/main.go
	
test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
	awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)