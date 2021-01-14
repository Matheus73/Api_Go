package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Books []Book = []Book{
	{
		Id:     1,
		Title:  "The Lord of the rings",
		Author: "Tolkien",
		Price:  39.90,
	},
	{
		Id:     2,
		Title:  "Harry Potter",
		Author: "J. K. Rolling",
		Price:  29.90,
	},
}

func list_books(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode(Books)
}

func create_book(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)

	var newBook Book
	json.Unmarshal(body, &newBook)
	newBook.Id = len(Books) + 1

	Books = append(Books, newBook)

	encoder := json.NewEncoder(w)
	encoder.Encode(newBook)

}
