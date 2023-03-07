package server

import (
	"io/ioutil"

	"github.com/labstack/echo/v4/middleware"

	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Version int
	host    string
	port    string
}

var Cfg ServerConfig

func init() {
	f, err := ioutil.ReadFile("./server_config.toml")

	if err != nil {
		panic(err)
	}

	var cfg ServerConfig

	{
		err := toml.Unmarshal(f, &cfg)
		if err != nil {
			panic(err)
		}
	}

	Cfg = cfg
}

func Cors() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{"%v:%v", Cfg.host, Cfg.port},
	}
}
