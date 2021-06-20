package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

func main() {
	e := echo.New()

	f, _ := os.Create("echo.log")
	e.Logger.SetOutput(io.MultiWriter(f, os.Stdout))   // To both log file and console at the same time.

	e.Logger.Fatal(e.Start(":1323"))
}