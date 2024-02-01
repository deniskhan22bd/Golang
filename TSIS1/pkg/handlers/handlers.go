package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/deniskhan22bd/Golang/test/testdata"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r* http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is running! \nWeb app about books by Khan Denis")
}

func GetBooks(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(testdata.Books)

	if err != nil{
		return
	}

	w.Write(jsonResponse)
}

func GetBookById(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	vars_id := vars["id"]

	for _, book := range testdata.Books{
		if strconv.Itoa(book.Id) == vars_id{
			jsonResponse, err := json.Marshal(book)

			if err != nil{
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func GetBookByTitle(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	vars_title := vars["title"]

	for _, book := range testdata.Books{
		if strings.ToLower(vars_title) == strings.ToLower(book.Title){
			jsonResponse, err := json.Marshal(book)

			if err != nil{
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
			return
		}
	}
	
	http.Error(w, "Book not found", http.StatusNotFound)
}

func GetBookByPublishedYear(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	vars_year := vars["year"]

	var books []testdata.Book

	for _, book := range testdata.Books{
		if strconv.Itoa(book.PublishedYear) == vars_year{
			books = append(books, book)
		}
	}
	
	if len(books) == 0{
		http.Error(w, "Book not found", http.StatusNotFound)
	} else{
		jsonResponse, err := json.Marshal(books)

		if err != nil{
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}