.PHONY: ;

INFRASTRUCTURE_PREFIX = infrastructure
BINARY_DIR = bin
NPM = npm

ci: test

full: clean test build synth

build:
	go build -o ./$(BINARY_DIR)/lambda/handler ./goapp/lambda

synth:
	npx cdk synth

clean:
	rm -rf $(BINARY_DIR)
test:
	go test --count=1 ./goapp/...


