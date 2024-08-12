#!/bin/bash

echo "Generating gRPC Codes"

# Ensure Go binaries are in PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go Codes
protoc --go_out=./go --go-grpc_out=./go auth.proto
protoc --go_out=./go --go-grpc_out=./go api.proto
echo "Go Codes Generated"

# Generate JavaScript code
protoc --js_out=import_style=commonjs,binary:./js/auth-service auth.proto
protoc --js_out=import_style=commonjs,binary:./js/api-service api.proto
echo "JS Code Generated"

# Generate gRPC-Web code
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/auth-service auth.proto
protoc --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web-js/api-service api.proto
echo "JS Web Generated"