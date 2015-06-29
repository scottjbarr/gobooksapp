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
	configFile := flag.String("config", "", "Config file")

	flag.Parse()

	if *configFile == "" {
		fmt.Println("A --config must be given. See config/example.conf")
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

func main() {
	http.HandleFunc("/books", BooksIndex)

	bindAddress := fmt.Sprintf("%v:%v", config.HTTP.Address, config.HTTP.Port)
	log.Printf("Listening on %v\n", bindAddress)

	http.ListenAndServe(bindAddress, nil)
}
