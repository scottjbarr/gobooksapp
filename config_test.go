package main

import (
	"reflect"
	"testing"
)

// Test helper. Thanks again, @keighl
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf(
			"Expected %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a))
	}
}

// Creates a Game that has a Glider in it.
func buildConfig() *Config {
	d := DB{
		URL: "user:pass@host/db",
	}

	h := HTTP{
		Address: "127.0.0.1",
		Port:    65535,
	}

	config := &Config{DB: d, HTTP: h}

	return config
}

func TestGetHTTPBindAddress(t *testing.T) {
	c := buildConfig()

	expect(t, "127.0.0.1:65535", c.GetHTTPBindAddress())
}
