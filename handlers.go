package main

import (
	"encoding/json"
	"net/http"
)

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

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
