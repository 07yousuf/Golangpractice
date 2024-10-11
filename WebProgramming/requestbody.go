package main

import (
	"fmt"
  "io"
	"net/http"
)

func Header(w http.ResponseWriter, r *http.Request) {
	//len := r.ContentLength
	//body := make([]byte, len)
	//_, err := r.Body.Read(body)
  body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Error while resquest reading ", err)
	}
	fmt.Fprintln(w, string(body))
}
func main() {
	server := &http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/header/", Header)
	server.ListenAndServe()
}
