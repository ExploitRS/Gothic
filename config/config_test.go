package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	c := Config{}
	c.Load("./server_config.toml")
	if c.server.address != "localhost" {
		t.Errorf("")
	}
}
