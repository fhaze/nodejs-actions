package main

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"net/http"
	"nodejs-actions/message"
)

func main() {
	e := echo.New()

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, message.HelloPerson("1.0.3"))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
