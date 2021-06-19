package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	g1 := e.Group("/v1")                  // Group v1
	g1.GET("/login",  loginHandler1)      // This will match /v1/login
	g1.GET("/submit", submitHandler1)     // This will match /v1/submit
	g1.GET("/read",   readHandler1)       // This will match /v1/read

	g2 := e.Group("/v2")                  // Group v2
	g2.GET("/login",  loginHandler2)      // This will match /v2/login
	g2.GET("/submit", submitHandler2)     // This will match /v2/submit
	g2.GET("/read",   readHandler2)       // This will match /v2/read

	e.Logger.Fatal(e.Start(":1323"))
}

func loginHandler1(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling loginHandler1")
}

func submitHandler1(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling submitHandler1")
}

func readHandler1(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling readHandler1")
}

func loginHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling loginHandler2")
}

func submitHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling submitHandler2")
}

func readHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "You are calling readHandler2")
}
