package config

import (
	"github.com/exploit-rs/gothic/config/database"
)

type Mode string

const (
	Prod Mode = "prod"
	Dev       = "dev"
	Demo      = "demo"
)

type Config struct {
	server   *Server
	database *database.Database
	mode     Mode
}

func (c *Config) New() *Config {
	return &Config{
		server:   NewServer(),
		database: database.New(),
	}
}

func (c *Config) Load(path string) {}
