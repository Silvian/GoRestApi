package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Silvian/GoRestApi/api"
)

func main() {
	// Init router
	router := mux.NewRouter()

	// Mock data
	api.LoadMockData()

	// Route handlers / Endpoints
	router.HandleFunc("/api/books", api.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", api.GetBook).Methods("GET")
	router.HandleFunc("/api/books", api.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", api.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", api.DeleteBook).Methods("DELETE")

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
