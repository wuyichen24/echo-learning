package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/pathparam/:value", func(c echo.Context) (err error) {
		value := c.Param("value")
		return c.String(http.StatusOK, value)
	})

	e.GET("/queryparam", func(c echo.Context) (err error) {
		value := c.QueryParam("key")
		return c.String(http.StatusOK, value)
	})

	e.GET("/header", func(c echo.Context) (err error) {
		value := c.Request().Header.Get("key")
		return c.String(http.StatusOK, value)
	})

	e.POST("/form", func(c echo.Context) (err error) {
		value := c.FormValue("key")
		return c.String(http.StatusOK, value)
	})

	e.Logger.Fatal(e.Start(":1323"))
}