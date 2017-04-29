# meta
NAME := mark2-server
VERSION := v1.0.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## setup
setup:
	go get github.com/golang/lint
	go get github.com/Masterminds/glide
	go get github.com/Songmu/make2help/cmd/make2help
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/{proto,protoc-gen-go}

## install dependencies
deps: setup
	glide install

## update dependencies
update: deps
	glide update

## run tests
test:
	go test $$(glide novendor -x | grep -v proto)

## lint
lint:
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x | grep -v proto); do\
		golint --set_exit_status $$pkg || exit $$?; \
	done

## build
bin/$(NAME):
	go build \
		-a \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		-o bin/$(NAME)

## install
install:
	go install $(LDFLAGS)

## show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint install
