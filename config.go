package main

import (
	"fmt"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"log"
)

// DB configuration struct
type DB struct {
	URL string
}

// HTTP configuration struct
type HTTP struct {
	Address string
	Port    int
}

// Config top level configuration struct
type Config struct {
	DB
	HTTP
}

// GetHTTPBindAddress returns a complete address:port
func (c *Config) GetHTTPBindAddress() string {
	return fmt.Sprintf("%v:%v", c.HTTP.Address, c.HTTP.Port)
}

// ParseConfig parses a config file
func ParseConfig(filename string) *Config {
	config := Config{}

	log.Printf("Loading config file %v", filename)
	sconf.Must(&config).Read(ini.File(filename))
	log.Printf("Loaded config : %v\n", config)
	return &config
}
