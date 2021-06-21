package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// login implements the authentication for users and provide a token if pass
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Verify username and password
	if username != "jone" || password != "abcd1234" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	expire := time.Now().Add(time.Hour * 1)                          // The token will expire in 1 hour

	// Set claims
	claims := token.Claims.(jwt.MapClaims)                           // Use JWT map claim
	claims["name"] = "jone"
	claims["admin"] = true
	claims["exp"] = expire.Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
		"expire": expire.Format("2006-01-02T15:04:05.999999-07:00"),
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func verify(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)                                           // Get token

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/verify", verify)                                     // Verify token

	e.Logger.Fatal(e.Start(":1323"))
}
