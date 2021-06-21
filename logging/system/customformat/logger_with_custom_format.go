package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${short_file}:${line}   ${message}")
	}

	e.GET("/logging", func(c echo.Context) (err error) {
		e.Logger.Info("This is info message")
		return c.String(http.StatusOK, "Success")
	})

	e.Logger.Fatal(e.Start(":1323"))
}