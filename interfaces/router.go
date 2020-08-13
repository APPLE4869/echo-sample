package interfaces

import (
	"github.com/apple/bff-web/interfaces/handler"
	"github.com/labstack/echo/v4"
)

// Router .
func Router(e *echo.Echo) {
	e.GET("/health_check", handler.HealthCheck)
}
