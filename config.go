package main

import (
	"fmt"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"log"
)

type DB struct {
	URL string
}

type HTTP struct {
	Address string
	Port    int
}

type Config struct {
	DB
	HTTP
}

func (c *Config) GetHTTPBindAddress() string {
	return fmt.Sprintf("%v:%v", c.HTTP.Address, c.HTTP.Port)
}

func ParseConfig(filename string) *Config {
	config := Config{}

	log.Printf("Loading config file %v", filename)
	sconf.Must(&config).Read(ini.File(filename))
	log.Printf("Loaded config : %v\n", config)
	return &config
}
