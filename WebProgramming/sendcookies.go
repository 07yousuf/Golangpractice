package main

import (
	"net/http"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "book publication",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}
func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.HandleFunc("/cookies", Cookies)
	server.ListenAndServe()
}
