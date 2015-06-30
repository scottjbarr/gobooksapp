package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
var config *Config

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	configFile := flag.String("config", "", "Config file")

	flag.Parse()

	if *configFile == "" {
		fmt.Println("A --config must be given. See config/example.conf.sample")
		os.Exit(1)
	}

	config = ParseConfig(*configFile)

	var err error
	db, err = sql.Open("mysql", config.DB.URL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/books", BooksIndex)
	log.Printf("Listening on %v\n", config.getHTTPBindAddress())
	http.ListenAndServe(config.getHTTPBindAddress(), Log(http.DefaultServeMux))
}
