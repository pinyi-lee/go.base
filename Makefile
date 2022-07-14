SHELL := /bin/bash -o pipefail
SERVICE := $(shell git remote -v | head -n1 | awk '{print $$2}' | sed 's/.*\///' | sed 's/\.git//')
VERSION := $(shell cat internal/pkg/config/config.go | grep 'env:"VERSION"' | grep -o 'envDefault:".*"' | sed s/[^0-9.]*//g)
COMMIT_ID := $(shell git rev-parse HEAD | awk '{print substr($$0,0,8)}')
GO_PATH := $(shell go env GOPATH)
FILE_PATH := $(shell pwd)
SOURCE_PATH := $(FILE_PATH:$(GO_PATH)%=%)
IMAGE := xxxxxxxxxxxx.dkr.ecr.us-west-2.amazonaws.com/account/$(SERVICE):$(VERSION)_$(COMMIT_ID)

echo:
	@echo "SERVICE: $(SERVICE)"
	@echo "VERSION: $(VERSION)"
	@echo "COMMIT_ID: $(COMMIT_ID)"
	@echo "GO_PATH: $(GO_PATH)"
	@echo "FILE_PATH: $(FILE_PATH)"
	@echo "SOURCE_PATH: $(SOURCE_PATH)"
	@echo "IMAGE: $(IMAGE)"

install:
	@go mod tidy

fmt:
	@gofmt -w .

clean:
	@rm -rf $(SERVICE)
	@rm -rf coverage
	@rm -rf pkg
	@rm -f Dockerfile
	@go clean -i ./...

build:
	@make clean
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@go build -v -o $(SERVICE) $(SOURCE_PATH)/cmd

run:
	@make build
	@./$(SERVICE)

check:
	@make clean
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@rm -rf coverage
	@mkdir coverage
	@go test ./... | tee coverage/test.log

docker:
	@make clean
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@export SOURCE_PATH=$(SOURCE_PATH); envsubst '$$SOURCE_PATH' < Dockerfile.temp > Dockerfile;
	@docker build -t $(IMAGE) .

help:
	@echo "make echo: echo parameter"
	@echo "make install: install all dependency"
	@echo "make fmt: format the coding style"
	@echo "make clean: remove binary file and dir"
	@echo "make build: build binary file"
	@echo "make run: run binary file"
	@echo "make check: run test"
	@echo "make docker: build docker"
