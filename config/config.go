package config

import (
	"fmt"
	"os"

	"github.com/exploit-rs/gothic/config/database"

	toml "github.com/pelletier/go-toml/v2"
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
	version  uint
}

func (c *Config) New() *Config {
	return &Config{
		server:   NewServer(),
		database: database.New(),
	}
}

func (c *Config) Load(path string) {
	read, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	toml.Unmarshal(read, &c)
}
