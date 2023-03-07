package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Status struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func getStatus(c echo.Context) error {
	s := Status{
		Name: "Jukan",
		Text: "command test hahaha",
	}
	return c.JSON(http.StatusOK, s)
}
