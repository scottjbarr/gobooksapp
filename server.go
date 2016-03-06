package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var db gorm.DB
var config *Config

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// It is nicer when this is in init() but that breaks testing because
	// the "config" flag is required, so we need to handle that in main().
	configFile := flag.String("config", "", "Config file")

	flag.Parse()

	if *configFile == "" {
		fmt.Println("A --config must be given. See config/example.conf.sample")
		os.Exit(1)
	}

	config = ParseConfig(*configFile)

	var err error
	db, err = gorm.Open("postgres", config.DB.URL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	// main router
	router := mux.NewRouter()

	// sub router for /api/v1
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/books", BooksIndex).Methods("GET")

	log.Printf("Listening on %v\n", config.GetHTTPBindAddress())

	http.ListenAndServe(config.GetHTTPBindAddress(), logger(router))
}
