package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

//handler for hello
func Hello() func(echo.Context) error {
	
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "heelo world!")
	}
}

func Hello2() func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "heelo2 world!")
	}
}


