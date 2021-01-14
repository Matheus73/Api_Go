package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main_route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá mundo")
}

func configRoutes(router *mux.Router) {
	router.HandleFunc("/", main_route)
	router.HandleFunc("/books", list_books).Methods("GET")
	router.HandleFunc("/books", create_book).Methods("POST")
}
