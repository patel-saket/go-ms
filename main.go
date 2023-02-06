package main

import (
	"go-ms/handler"
	"net/http"
)

func main() {
	println("Welcome to NewTech's watch me code session")
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	serveMux.Handle("/shoppingcart", &handler.ShoppingCartHandler{})
	http.ListenAndServe(":8080", serveMux)
}
