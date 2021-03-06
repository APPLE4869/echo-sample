package main

import (
	"github.com/apple/bff-web/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	interfaces.Router(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
