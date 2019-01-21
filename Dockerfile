FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache git
# gRPC files have been copied to the project 'vendor' folder
# RUN go get -u google.golang.org/grpc
COPY . /go/src/skeleton-grpc
COPY ./docker-entrypoint.sh /
# Here, I'm building the two executables into $GOPATH/bin
# This is equivalent to go install, plus, renaming binaries
RUN go build -o /go/bin/grpc /go/src/skeleton-grpc/main.go
RUN go build -o /go/bin/http /go/src/skeleton-grpc/http/main.go
# The following script will launch created commands above: http and grpc
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 50051