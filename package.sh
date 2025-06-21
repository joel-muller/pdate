#!/bin/bash

echo "Format the code"
go fmt ./...

echo "Run the tests"
go test ./...

echo "Build the binary"
go build -o pdate ./cmd/pdate

echo "Done"