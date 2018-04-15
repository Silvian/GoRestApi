package api

import (
  "encoding/json"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

// Init books var as a slice Book struct
 var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) //Get params
  for _, item := range books {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
  books = append(books, book)
  json.NewEncoder(w).Encode(book)
}

// Update Single Book
func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.ID == params["id"] {
      books = append(books[:index], books[index+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.ID = params["id"]
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

// Delete Single Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.ID == params["id"] {
      books = append(books[:index], books[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(books)
}