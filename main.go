package main

import (
	// "flag"
	// "fmt"
	// "net/http"

	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/exploit-rs/nerd/frontend"
	"github.com/exploit-rs/nerd/server"
)

func main() {
	devMode := false
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.Parse()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Exploit.RS!")
	})

	e.GET("/api/status", func(c echo.Context) error {
		s := server.Status{
			Name: "Jukan",
			Text: "command test hahaha",
		}
		return c.JSON(http.StatusOK, s)
	})

	// e.Static("/", "kkkkkkkkkkk")
	frontend.PrintFsys()

	if devMode {
		e.Use(middleware.CORSWithConfig(server.Cors()))
		fmt.Println("Server running in dev mode")
	}

	e.Logger.Fatal(e.Start(":1337"))
}
