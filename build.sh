#!/bin/bash

pushd proto || exit 1
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  service.proto
popd || exit 1
