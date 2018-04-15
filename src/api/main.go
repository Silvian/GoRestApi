package main

import (
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
  ID        string `json:"id"`
  Isbn      string `json:"isbn"`
  Title     string `json:"title"`
  Author    *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
  Firstname   string `json:"firstname"`
  Lastname    string `json:"lastname"`
}

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

func main() {
  // Init router
  router := mux.NewRouter()

  // Mock Data - @TODO implement DB
  books = append(books, Book{
    ID: "1",
    Isbn: "44874",
    Title: "To Kill a Mocking Bird",
    Author: &Author{
      Firstname: "Harper",
      Lastname: "Lee",
    },
  })

  books = append(books, Book{
    ID: "2",
    Isbn: "45723",
    Title: "Pride and Prejudice",
    Author: &Author{
      Firstname: "Jane",
      Lastname: "Austin",
    },
  })

  // Route handlers / Endpoints
  router.HandleFunc("/api/books", getBooks).Methods("GET")
  router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/api/books", createBook).Methods("POST")
  router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  server := http.ListenAndServe(":8080", router)
  log.Fatal(server)
}
