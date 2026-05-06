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

// HandleGetPairingStatus processes GET requests to retrieve pairing request status
func (h *PairingRequestsHandler) HandleGetPairingStatus(c *echo.Context) error {
	// Extract request ID from URL path parameter
	requestID := c.Param("id")
	if requestID == "" {
		slog.Warn("missing request ID in path parameter")
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:  "request ID is required",
			Status: "error",
		})
	}

	requestID = strings.TrimSpace(requestID)
	if requestID == "" {
		slog.Warn("empty request ID in path parameter")
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:  "request ID cannot be empty",
			Status: "error",
		})
	}

	// TODO: Check actual pairing status from database
	// For now, always return pending (202) as this is MVP without persistence
	slog.Info("pairing status check", "request_id", requestID)

	// Return 202 Accepted with pending status
	return c.JSON(http.StatusAccepted, models.PairingStatusResponse{
		Status: "pending",
	})
}
