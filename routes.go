package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá mundo")
}

func configRoutes(router *mux.Router) {
	router.HandleFunc("/", mainRoute)
	router.HandleFunc("/books", listBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", editBook).Methods("PUT")
}
