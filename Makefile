SHELL := /bin/bash -o pipefail
SERVICE := $(shell git remote -v | head -n1 | awk '{print $$2}' | sed 's/.*\///' | sed 's/\.git//')
VERSION := $(shell cat internal/pkg/config/config.go | grep 'env:"VERSION"' | grep -o 'envDefault:".*"' | sed s/[^0-9.]*//g)
COMMIT_ID := $(shell git rev-parse HEAD | awk '{print substr($$0,0,8)}')
GO_PATH := $(shell go env GOPATH)
FILE_PATH := $(shell pwd)
SOURCE_PATH := $(FILE_PATH:$(GO_PATH)%=%)
PKG_TO_TEST := $(shell go list ./... | grep -v docs | grep -v test | tr '\n' ',' | sed 's/,$$//')
IMAGE := xxxxxxxxxxxx.dkr.ecr.us-west-2.amazonaws.com/account/$(SERVICE):$(VERSION)_$(COMMIT_ID)

echo:
	@echo $(SERVICE)
	@echo $(VERSION)
	@echo $(COMMIT_ID)
	@echo $(GO_PATH)
	@echo $(FILE_PATH)
	@echo $(SOURCE_PATH)
	@echo $(PKG_TO_TEST)
	@echo $(IMAGE)

install:
	@go mod tidy

fmt:
	@gofmt -w .

build:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@go build -v -o $(SERVICE)

run:
	@make build
	@./$(SERVICE)

clean:
	@rm -rf $(SERVICE)
	@rm -rf coverage
	@rm -rf pkg
	@rm -f Dockerfile
	@go clean -i ./...

testReport:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@rm -rf coverage
	@mkdir coverage
	@go test -v -coverpkg=$(PKG_TO_TEST) -coverprofile=coverage/coverage.out -covermode=count 2>&1 | tee coverage/test.log
	@cat coverage/test.log | go-junit-report > coverage/test_result.xml
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@gocover-cobertura < coverage/coverage.out > coverage/coverage.xml

docker:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO_PATH)/bin/swag init -g cmd/main.go
	@export SOURCE_PATH=$(SOURCE_PATH); envsubst '$$SOURCE_PATH' < Dockerfile.temp > Dockerfile;
	@docker build -t $(IMAGE) .

help:
	@echo "make install: install all dependency"
	@echo "make fmt: format the coding style of this project"
	@echo "make build: build this project to binary file"
	@echo "make clean: remove binary file and dir when building/testing"
	@echo "make testReport: run test and generate coverage.html in coverage/"
