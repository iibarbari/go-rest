package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

// Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

var books []Book

func main() {
	// Init
	r := mux.NewRouter()

	// Mock Data TODO:Implement DB
	books = append(books, Book{ID: "1", Isbn: "123456", Title: "Book One", Author: &Author{Firstname: "Me me", Lastname: "Big Boi"}})
	books = append(books, Book{ID: "2", Isbn: "678910", Title: "Book Two", Author: &Author{Firstname: "Me you", Lastname: "Big Mamma"}})
	books = append(books, Book{ID: "3", Isbn: "111213", Title: "Book Three", Author: &Author{Firstname: "You you", Lastname: "Big You"}})
	books = append(books, Book{ID: "4", Isbn: "141516", Title: "Book Four", Author: &Author{Firstname: "Me us", Lastname: "Big Usher"}})

	// Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
