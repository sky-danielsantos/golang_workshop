GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=csv2json
VERSION?=0.0.0

## Build:
build:
	@echo "== build"
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -o out/bin/$(BINARY_NAME) .


clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

test:
	$(GOTEST) -v -race ./...

test.unit:
	$(GOTEST) -v -race -run 'Unit' ./...

check: vet lint

.PHONY: vet
vet:
	@echo "== vet"
	@go vet ./...

.PHONY: lint
lint:
	@echo "== go lint"
	@GOGC=45 golangci-lint --timeout 5m run

run:
	go run main.go

all: check build
