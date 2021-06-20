package main

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	basicAuthMiddleware := middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("abcd1234")) == 1 {
			return true, nil
		}
		return false, nil
	})

	e.GET("/basicauth", basicAuthHandler, basicAuthMiddleware)
	e.Logger.Fatal(e.Start(":1323"))
}

func basicAuthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}
