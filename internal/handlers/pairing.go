package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/xcoulon/claw-device-pairing/internal/models"
)

// PairingHandler handles device pairing requests
type PairingHandler struct {
	// Future: add dependencies like database client
}

// NewPairingHandler creates a new pairing handler
func NewPairingHandler() *PairingHandler {
	return &PairingHandler{}
}

// HandlePairDevice processes POST requests to the /pair-device endpoint
func (h *PairingHandler) HandlePairDevice(c *echo.Context) error {
	// Parse JSON request body
	var req models.PairDeviceRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("error parsing request", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "Invalid request format",
			Status: "error",
		})
	}

	// Validate ID field (trim whitespace and check non-empty)
	req.ID = strings.TrimSpace(req.ID)
	if req.ID == "" {
		slog.Warn("empty or whitespace-only ID received")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "id cannot be empty",
			Status: "error",
		})
	}

	// Log successful request
	slog.Info("pairing request received", "device_id", req.ID)

	// Return success response
	return c.JSON(http.StatusOK, models.PairDeviceResponse{
		Status:  "success",
		Message: "pairing request received",
	})
}
