package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func listBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode(Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {

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

func editBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	variables := mux.Vars(r)

	id, err := strconv.Atoi(variables["bookId"])

	if err != nil {
		fmt.Println("Error: Problem with price value")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newBook Book

	body, err_request := ioutil.ReadAll(r.Body)

	if err_request != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err_json := json.Unmarshal(body, &newBook)

	if err_json != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ind := -1

	for i := range Books {
		if Books[i].Id == id {
			ind = i
			break
		}
	}

	if ind < 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	newBook.Id = Books[ind].Id
	Books[ind] = newBook

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newBook)

}
