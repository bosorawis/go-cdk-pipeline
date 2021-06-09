.PHONY: ;

INFRASTRUCTURE_PREFIX = infrastructure
BINARY_DIR = bin
NPM = npm
APP_DIR = ./goapp

ci: test

full: clean test build synth

app: clean test build

install: install-go install-cdk

install-go:
	wget -q -O - https://git.io/vQhTU | bash

install-cdk:
	npm ci

build:
	GOOS=linux GO111MODULE=on go build -o ./$(BINARY_DIR)/lambda/handler $(APP_DIR)/lambda

synth:
	npx cdk synth

clean:
	rm -rf $(BINARY_DIR)

test:
	go test --count=1 $(APP_DIR)/...


