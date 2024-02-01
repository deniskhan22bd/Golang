package main

import (
	"log"
	"github.com/deniskhan22bd/Golang/pkg/handlers"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/id/{id}", handlers.GetBookById).Methods("GET")
	r.HandleFunc("/books/title/{title}", handlers.GetBookByTitle).Methods("GET")
	r.HandleFunc("/books/year/{year}", handlers.GetBookByPublishedYear).Methods("GET")
	http.Handle("/", r)

	log.Print("starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}