package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/xcoulon/claw-device-pairing/internal/models"
	"github.com/xcoulon/claw-device-pairing/internal/version"
)

func HandleHealth(c *echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{
		Status:     "ok",
		CommitHash: version.CommitHash,
		BuildTime:  version.BuildTime,
	})
}
