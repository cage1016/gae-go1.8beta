package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
