package main

import (
	"log"
	"net/http"
)

func main() {
	//create a server
	myServer := &http.Server{
		//set server address
		Addr: "127.0.0.1:8080",
		//making handler function
		Handler: &myHandler{},
	}
	//starting the server
	log.Fatal(myServer.ListenAndServe())
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ResponseBody := []byte("<html><head></head><body>Hello</body></html>")
	_, err := w.Write(ResponseBody)
	if err != nil {
		log.Printf("error while writing on the body %s", err)
	}
}
