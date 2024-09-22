package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	myServer := &http.Server{
		Addr:    "localhost:8081",
		Handler: &myHandler{},
	}
	//launching the Server
	fmt.Printf("Starting the server port at- 8081")
	log.Fatal(myServer.ListenAndServe())
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	DesingA := []byte("<html><head>Welcome to the golang Hotel</head><body><p>The Golang Hotel is Relaxing Place!</p><p>We offer 20% discount if you call this number:<strong>134591</strong></p></body>")
	DesignB := []byte("<html><head>Welcome to the golang Hotel</head><body><h1>The Golang Hotel is Relaxing Place!</h1><h2>We offer 20% discount if you call this number:<strong>134592</strong></h2></body>")

	minutes := time.Now().Minute()

	//logical part
	if minutes%2 != 0 {
		_, err := w.Write(DesingA)
		if err != nil {
			fmt.Printf("error while writing the response body")
		}
	} else {
		_, err := w.Write(DesignB)
		if err != nil {
			fmt.Printf("Error, while writing the response body")
		}
	}
}
