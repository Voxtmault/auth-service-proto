#!/bin/bash

echo "Generating gRPC Codes"

# Ensure Go binaries are in PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go Codes
protoc --go_out=./go --go-grpc_out=./go auth.proto
echo "Go Codes Generated"

# Generate JavaScript code
protoc --js_out=import_style=commonjs,binary:./js auth.proto
echo "JS Code Generated"

# Generate gRPC-Web code
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./js auth.proto
echo "JS Web Generated"