package handler

import (
	"net/http"

	"github.com/apple/bff-web/cache"
	"github.com/labstack/echo/v4"
)

// HealthCheckResult .
type HealthCheckResult struct {
	Message string `json:"message"`
}

// HealthCheck .
func HealthCheck(c echo.Context) error {
	res := cache.Get("ok")
	cache.Set("ok", "XXXXXXXXXXX")
	r := &HealthCheckResult{
		Message: res,
	}
	return c.JSON(http.StatusOK, r)
}
