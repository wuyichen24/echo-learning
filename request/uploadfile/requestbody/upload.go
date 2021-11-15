package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}

func upload(c echo.Context) error {
	content, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return c.HTML(http.StatusOK, "success")
}


