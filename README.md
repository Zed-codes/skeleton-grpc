# Skeleton gRPC

This is a very basic implementation of a gRPC server that relies on a HTTP server.
This project has been done in the course of learning Golang and gRPC.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need to install Go for your particullar system:

* [Go](https://golang.org/dl/)

If you want to edit the gRPC proto file, you will need to install Google Protocol Buffer:

* [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1)

To run the service dockerized, you will need to install Docker:

* [Docker](https://www.docker.com/get-started)

### Installing

Clone this repository to your local machine:

```
git clone https://github.com/Zed-codes/skeleton-grpc
```

Point to the repository root folder and use the command 'Make build'

```
make build
```

Use the command 'Make run' to run the gRPC and HTTP servers

```
make run
```
