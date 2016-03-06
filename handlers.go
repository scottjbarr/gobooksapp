package main

import (
	"encoding/json"
	"net/http"
)

// BooksIndex responds to index requests for the Book resource
func BooksIndex(w http.ResponseWriter, r *http.Request) {
	var bks []*Book
	err := db.Find(&bks)

	if err.Error != nil {
		http.Error(w, err.Error.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(bks)
	w.Write(js)
}
