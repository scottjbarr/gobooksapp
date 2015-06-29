package main

import (
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"log"
)

type Config struct {
	DB struct {
		URL string
	}
	HTTP struct {
		Address string
		Port    int
	}
}

func ParseConfig(filename string) *Config {
	config := Config{}

	log.Printf("Loading config file %v", filename)
	sconf.Must(&config).Read(ini.File(filename))
	log.Printf("Found config : %v\n", config)
	return &config
}
