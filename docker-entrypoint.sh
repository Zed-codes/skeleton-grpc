#!/bin/sh
echo "Starting service"
go run /go/src/skeleton-grpc/http/main.go &
go run /go/src/skeleton-grpc/main.go