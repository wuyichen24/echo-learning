package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.WARN)     // You can try different levels: DEBUG, INFO, WARN, ERROR, OFF

	e.GET("/info", func(c echo.Context) (err error) {
		e.Logger.Info("This is info log")
		return c.String(http.StatusOK, "info log")
	})

	e.GET("/debug", func(c echo.Context) (err error) {
		e.Logger.Debug("This is debug log")
		return c.String(http.StatusOK, "debug log")
	})

	e.GET("/warn", func(c echo.Context) (err error) {
		e.Logger.Warn("This is warn log")
		return c.String(http.StatusOK, "warn log")
	})

	e.GET("/error", func(c echo.Context) (err error) {
		e.Logger.Error("This is error log")
		return c.String(http.StatusOK, "error log")
	})

	e.Logger.Fatal(e.Start(":1323"))
}