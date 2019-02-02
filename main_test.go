// This is basically a gRPC client written in Golang

package main

import (
	"testing"
	"context"
	"time"
	"google.golang.org/grpc"
	"skeleton-grpc/api"
)

const (
	address     = "localhost:50051"
	testValue = "test"
)

func TestGrpcServer(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Couldn't connect: %v", err)
	}
	defer conn.Close()
	c := api.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SayHello(ctx, &api.HelloRequest{Name: testValue})
	if err != nil {
		t.Fatalf("Could not get response: %v", err)
	}
}

func TestFullExecution(t * testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Couldn't connect: %v", err)
	}
	defer conn.Close()
	c := api.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: testValue})
	if err != nil {
		t.Fatalf("Could not get response: %v", err)
	}
	var responseData string = string(r.Message)
	responseData = responseData[6:len(responseData)]
	if responseData != "000" {
		t.Fatal("Wrong value returned by server.")
	}
}