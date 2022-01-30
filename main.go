package main

import (
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"net/http"
)

var counter = 0

func main() {
	e := echo.New()

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		counter++
		return c.String(http.StatusOK, fmt.Sprintf("Hello! Counter v2.0: %d", counter))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
