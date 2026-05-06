package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/xcoulon/claw-device-pairing/internal/models"
)

// PairingRequestsHandler handles device pairing requests
type PairingRequestsHandler struct {
	// Future: add dependencies like database client
}

// NewPairingRequestsHandler creates a new pairing handler
func NewPairingRequestsHandler() *PairingRequestsHandler {
	return &PairingRequestsHandler{}
}

// HandlePairDevice processes POST requests to the /pairing-requests endpoint
func (h *PairingRequestsHandler) HandlePairDevice(c *echo.Context) error {
	// Parse JSON request body
	var req models.PairingDeviceRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("error parsing request", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "Invalid request format",
			Status: "error",
		})
	}

	// Validate RequestID field (trim whitespace and check non-empty)
	req.RequestID = strings.TrimSpace(req.RequestID)
	if req.RequestID == "" {
		slog.Warn("empty or whitespace-only requestId received")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "requestId cannot be empty",
			Status: "error",
		})
	}

	// Log successful request
	slog.Info("pairing request received", "request_id", req.RequestID)

	// Return success response
	return c.JSON(http.StatusOK, models.PairingDeviceResponse{
		Status:  "success",
		Message: "pairing request received",
	})
}
