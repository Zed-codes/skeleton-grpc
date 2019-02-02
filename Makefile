VERSION = $(shell cat version)

.PHONY: help
.SILENT: help
help:
	echo "Please use one of the following targets:"
	echo " make test	Performs tests on live servers."
	echo " make build	Builds executables."
	echo " make run	Runs the service locally."
	echo " make dockrize	Dockerize the service."
	echo " make dockerun	Runs the dockerized service."
	echo " make clean	Removes created executables if any."

.PHONY: test
test: build
	bin/http &
	bin/grpc &
	go test -v
	killall -9 grpc
	killall -9 http

.PHONY: run
run: build
	bin/http &
	bin/grpc

.PHONY: build
build: clean
	go build -o bin/grpc main.go
	go build -o bin/http http/main.go

.PHONY: dockerun
dockerun: dockerize
	docker run -it -p 50051:50051 "github.com/zed-codes/skeleton-grpc:${VERSION}"
	
.PHONY: dockerize
dockerize: clean
	docker build -t "github.com/zed-codes/skeleton-grpc:${VERSION}" .
	
.PHONY: clean
clean:
	rm -rf ./bin