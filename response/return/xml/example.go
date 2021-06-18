package main

import (
	"encoding/xml"
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

	e.GET("/xml", func(c echo.Context) (err error) {
		return c.XML(http.StatusOK, u)
	})

	e.GET("/xml/pretty", func(c echo.Context) (err error) {
		return c.XMLPretty(http.StatusOK, u, "    ")
	})

	e.GET("/xml/stream", func(c echo.Context) (err error) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return xml.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/xml/blob", func(c echo.Context) (err error) {
		bytes, _ := xml.Marshal(u)
		return c.XMLBlob(http.StatusOK, bytes)
	})

	e.Logger.Fatal(e.Start(":1323"))
}