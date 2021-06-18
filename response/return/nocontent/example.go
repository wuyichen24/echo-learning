package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/nocontent", func(c echo.Context) (err error) {
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":1323"))
}