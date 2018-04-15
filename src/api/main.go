package api

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  // Init router
  router := mux.NewRouter()

  // Mock data
  loadMockData()

  // Route handlers / Endpoints
  router.HandleFunc("/api/books", getBooks).Methods("GET")
  router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  router.HandleFunc("/api/books", createBook).Methods("POST")
  router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  server := http.ListenAndServe(":8080", router)
  log.Fatal(server)
}
