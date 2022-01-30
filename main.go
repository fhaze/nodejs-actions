package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"nodejs-actions/message"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, message.HelloPerson("1.0.1"))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
