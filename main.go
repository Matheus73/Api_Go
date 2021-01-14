package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server running in the port: http://localhost:1337/")
	router := mux.NewRouter().StrictSlash(true)
	configRoutes(router)

	log.Fatal(http.ListenAndServe(":1337", router))
}
