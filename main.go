package main

import (
	"fmt"
	"skeleton-grpc/service"
)

func main() {
	fmt.Println("gRPC server starting")
	service.Start()
}
