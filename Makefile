GO11MODULES=on
APP?=go-algo
COMMIT_SHA=$(shell git rev-parse --short HEAD)


.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build cmd/btree/btree.go

.PHONY: run
## run: runs go run main.go
run:
	go run main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@rm -rf ${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=5 -race ./...

.PHONY: test
## test: runs go test with default values
bench:
	 go test -bench=.  -benchmem ./...

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
