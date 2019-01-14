package router

import (
	"context" // defines request-scoped values across API boundaries and between processes.
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	//"os"
	"skeleton-grpc/api"
)

// We are defining the Server struct that will be used to start our gRPC server/service
// I used "Server" to make the variable 'exportable' and readable from service.go file

type Server struct{}

// The following methods are related to the "interface": GreeterServer in pb.go (api)
// Here, we define the "real" work of the server regarding the configured Calls/Methods in the .proto file

func (s *Server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	
	// Quering the HTTP server
	response, err := http.Get("http://localhost:9090/?name=" + in.Name)

	// Prepare a variable to host the final result
	var responseDataFormatted string = ""

	if err != nil {
		//fmt.Print(err.Error())
		//os.Exit(1)
		fmt.Println("Couldn't connect to the HTTP server, make sure it's running!")
	} else {

		responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	// Formatting the HTTP response string
	// in future more complex examples, parsing or scrapping will happen here
	responseData = responseData[8:len(responseData)]
	
	responseDataFormatted = string(responseData)
	fmt.Println("Hello", responseDataFormatted)

	}

	return &api.HelloReply{Message: "Hello " + responseDataFormatted}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "Hello again " + in.Name}, nil
}
