# Skeleton gRPC

This is a very basic implementation of a gRPC server that relies on a HTTP server.

This simple project has been done through the course of learning Golang and gRPC.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need to install Go for your particular system:

* [Go](https://golang.org/dl/)

If you want to edit the gRPC proto file, you will need to install Google Protocol Buffer:

* [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1)

To run as dockerized service, you will need to install Docker:

* [Docker](https://www.docker.com/get-started)

### Installing and running locally

Clone this repository to your local machine:

```
git clone https://github.com/Zed-codes/skeleton-grpc
```

Point to the repository root folder and use the command `Make build`

```
make build
```

Use the command `Make run` to run the gRPC and HTTP servers

```
make run
```

### Installing and running in Docker

Make sure Docker is installed and running:

```
docker version
docker ps
```

Point to the repository root folder and use the command `Make dockerize`
This will dockerize the service in a minimal image.

```
make dockerize
```

Point to the repository root folder and use the command `Make dockerun`
This will execute above task then run the created image.

```
make dockerun
```

## Testing

To test if the gRPC server is working as expected, we make use of gRPC reflection employing the tool Evans.

Make sure [Evans](https://github.com/ktr0731/evans) is installed on your system, you can install it with Go:

```
go get github.com/ktr0731/evans
```

To start Evans in reflection mode, use the following command which will check for a listening gRPC server on the default port 50051:

```
evans -r
```

## Usage

With Evans successfully started in reflection mode with previous command, we can select the service "Greeter" with following command:

```
service Greeter
```

We can now view available RPCs on the selected server:

```
show rpc
```

The important RPC in this project is SayHello which takes a name as input, makes a query to the already running HTTP server with the entered name, if it is known, the related ID is returned, otherwise 'Stranger' is returned, the gRPC server returns finally `Hello <ID>`:

```
call SayHello
```

