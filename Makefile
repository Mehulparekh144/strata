# Project variables
BINARY_NAME = strata-server
PROTO_SRC   = api/strata.proto
PROTO_OUT   = api

# Default target
all: build

## -------- Proto --------
proto:
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_SRC)

## -------- Build --------
build:
	go build -o bin/$(BINARY_NAME) ./cmd/server

run: build
	./bin/$(BINARY_NAME)

clean:
	rm -rf bin/ $(PROTO_OUT)/*.pb.go data/*

## -------- Dev Helpers --------
tidy:
	go mod tidy

deps:
	go get google.golang.org/grpc
	go get google.golang.org/protobuf
	go get github.com/cockroachdb/pebble