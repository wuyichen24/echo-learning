package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// login implements the authentication for users and provide a token if pass the authentication
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Verify username and password
	if username != "jone" || password != "abcd1234" {
		return echo.ErrUnauthorized
	}

	tokenMap, err := generateTokenPair ()                            // Generate access token and refresh token
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, tokenMap)
}

// generateTokenPair implements the functionality to generate access token and refresh token
func generateTokenPair () (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	expire := time.Now().Add(time.Hour * 1)                          // The token will expire in 1 hour

	// Set claims
	claims := token.Claims.(jwt.MapClaims)                           // Use JWT map claim
	claims["name"] = "jone"
	claims["admin"] = true
	claims["exp"] = expire.Unix()

	// Generate access token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": t,
		"refresh_token": rt,
		"expire": expire.Format("2006-01-02T15:04:05.999999-07:00"),
	}, nil
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

func refresh(c echo.Context) error {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	err := c.Bind(&tokenReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(tokenReq.RefreshToken) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "The refresh_token can not be missing or empty.")
	}

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["sub"].(float64)) == 1 {

			newTokenPair, err := generateTokenPair()
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, newTokenPair)
		}
		return echo.ErrUnauthorized
	}
	return err
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)                                      // Get token
	e.POST("/refresh", refresh)                                  // Refresh token

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/verify", verify)                                     // Verify token

	e.Logger.Fatal(e.Start(":1323"))
}
