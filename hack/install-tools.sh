#! /usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# sudo apt install -y protobuf-compiler
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
