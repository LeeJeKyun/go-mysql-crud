package main

import (
	"go-mysql-crud/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	c := new(handler.Client)

	http.Handle("/client", c)

	http.ListenAndServe(":8080", nil)
}
