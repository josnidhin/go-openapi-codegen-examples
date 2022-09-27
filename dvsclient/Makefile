#
# @author Jose Nidhin
#
PROJECT_NAME := $(shell basename "$(PWD)")
GO_SRC_FILES := $(shell find . -type f -name '*.go')
GO_SRC_MAIN := $(shell find . -type f -name '*.go' ! -name '*_test.go')
GO_ENVFLAGS=CGO_ENABLED=0

include .env
export

all: tidy vet fmt simplify test clean run

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: simplify
simplify:
	gofmt -s -l -w $(GO_SRC_FILES)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	go clean -x
	rm -f cover.out coverage.html

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -v -o $(PROJECT_NAME) $(GO_SRC_MAIN)

.PHONY: generate-dvsapi
generate-dvsapi:
	docker run -it --rm \
	--volume "$(PWD):/local" \
	--volume "/usr/local/go:/usr/local/go" \
	--env GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -w" \
	openapitools/openapi-generator-cli:v6.2.0 generate \
	--enable-post-process-file \
	--input-spec /local/dvs-api-spec.json \
	--generator-name go \
	--package-name dvsapi \
	--additional-properties structPrefix=true,enumClassPrefix=true \
	--output /local/dvsapi

	sudo chown -R $(USER):$(USER) dvsapi
	chmod -R g+w dvsapi
	rm dvsapi/go.mod dvsapi/go.sum
