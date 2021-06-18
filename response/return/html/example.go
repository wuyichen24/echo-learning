package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/html", func(c echo.Context) (err error) {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})
	e.GET("/html/blob", func(c echo.Context) (err error) {
		bytes := []byte(`<strong>Hello, World!</strong>`)
		return c.HTMLBlob(http.StatusOK, bytes)    // bytes is byte array
	})

	e.Logger.Fatal(e.Start(":1323"))
}