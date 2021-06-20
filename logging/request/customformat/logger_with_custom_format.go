package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_custom} ${method} ${uri} ${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.000",
	}))

	e.GET("/logging", func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "Success")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
