package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	u := User{
		Name:"Joe",
		Email:"joe@labstack",
	}

	e.GET("/json", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/json/pretty", func(c echo.Context) (err error) {
		return c.JSONPretty(http.StatusOK, u, "    ")
	})

	e.GET("/json/stream", func(c echo.Context) (err error) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/json/blob", func(c echo.Context) (err error) {
		bytes, _ := json.Marshal(u)
		return c.JSONBlob(http.StatusOK, bytes)
	})

	e.Logger.Fatal(e.Start(":1323"))
}