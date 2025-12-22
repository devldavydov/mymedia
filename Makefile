BUILD_DATE := $(shell date +'%d.%m.%Y %H:%M:%S')
BUILD_COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: all
all: clean generate build test

.PHONY: generate
generate:
	@echo "\n### $@"
	go generate ./...

.PHONY: build
build: build_bot build_cli

.PHONY: build_bot
build_bot:
	@echo "\n### $@"
	@mkdir -p ./bin
	@cd cmd/mymediabot && \
	go build \
	-ldflags "-X 'main.buildDate=$(BUILD_DATE)' -X main.buildCommit=$(BUILD_COMMIT)" \
	-o ../../bin/mymediabot .	 

.PHONY: build_cli
build_bot:
	@echo "\n### $@"
	@mkdir -p ./bin
	@cd cmd/mymediacli && \
	go build \
	-ldflags "-X 'main.buildDate=$(BUILD_DATE)' -X main.buildCommit=$(BUILD_COMMIT)" \
	-o ../../bin/mymediacli .	

.PHONY: test
test:
	@echo "\n### $@"
	go test ./... -v --count 1

.PHONY: clean
clean:
	@echo "\n### $@"
	@rm -rf ./bin		 