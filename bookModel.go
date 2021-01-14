package main

type Book struct {
	Id     int     `json: "id"`
	Title  string  `json: "title"`
	Author string  `json: "author"`
	Price  float64 `json: "price"`
}
