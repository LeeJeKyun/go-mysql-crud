package main

import (
	"go-mysql-crud/handler"
	"net/http"
)

func main() {
	c := new(handler.Client)
	http.Handle("/client", c)
	http.ListenAndServe(":8080", nil)
}
