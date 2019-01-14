FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache git
RUN go get -u google.golang.org/grpc
COPY . /go/src/skeleton-grpc
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 50051