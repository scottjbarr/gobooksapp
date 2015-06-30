package main

import (
	"fmt"
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

func (c *Config) getHTTPBindAddress() string {
	return fmt.Sprintf("%v:%v", c.HTTP.Address, c.HTTP.Port)
}

func ParseConfig(filename string) *Config {
	config := Config{}

	log.Printf("Loading config file %v", filename)
	sconf.Must(&config).Read(ini.File(filename))
	log.Printf("Loaded config : %v\n", config)
	return &config
}
