package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nfode/dummy-app/internal/apis"
)

func main() {
	e := echo.New()
	apis.API{Echo: e}.Setup()
	e.Logger.Fatal(e.Start(":1323"))
}
