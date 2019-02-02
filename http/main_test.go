package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"testing"
)

func TestHttpServer(t *testing.T) {

	// This function considers the HTTP server already running and queries it
	// with a special value for variable 'name'
	// Following is similar to code in router.go

	// Quering the HTTP server
	response, err := http.Get("http://localhost:9090/?name=test")

	// Prepare a variable to host the final result
	var responseDataFormatted string = ""

	if err != nil {
		//fmt.Print(err.Error())
		//os.Exit(1)
		t.Fatal("Couldn't connect to the HTTP server, make sure it's running!")
	} else {

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		// Formatting the HTTP response string
		responseData = responseData[8:len(responseData)]
	
		responseDataFormatted = string(responseData)
		if responseDataFormatted != "000" {
			t.Fatal("Wrong value returned.")
		}
	}
}
