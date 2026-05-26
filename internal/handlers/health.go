package handlers

import (
	"net/http"

	"github.com/codeready-toolchain/claw-device-pairing/internal/models"
	"github.com/codeready-toolchain/claw-device-pairing/internal/version"
	"github.com/labstack/echo/v5"
)

func HandleHealth(c *echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{
		Status:     "ok",
		CommitHash: version.CommitHash,
		BuildTime:  version.BuildTime,
	})
}
