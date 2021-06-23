package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/error", func(c echo.Context) (err error) {
		return echo.NewHTTPError(http.StatusBadRequest, "This is the error message.")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
