package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
)

var db gorm.DB
var config *Config

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// nicer if this is in init() but that breaks testing because init()
	// doesn't know/care that a test is running.
	configFile := flag.String("config", "", "Config file")

	flag.Parse()

	if *configFile == "" {
		fmt.Println("A --config must be given. See config/example.conf.sample")
		os.Exit(1)
	}

	config = ParseConfig(*configFile)

	var err error
	db, err = gorm.Open("mysql", config.DB.URL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", BooksIndex)
	log.Printf("Listening on %v\n", config.GetHTTPBindAddress())
	http.ListenAndServe(config.GetHTTPBindAddress(), Log(http.DefaultServeMux))
}
