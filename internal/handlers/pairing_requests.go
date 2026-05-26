package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/codeready-toolchain/claw-device-pairing/internal/models"
	"github.com/labstack/echo/v5"
)

// K8sManager interface for Kubernetes operations
type K8sManager interface {
	IsEnabled() bool
	CreatePairingRequest(ctx context.Context, requestID string) error
	GetPairingRequestStatus(ctx context.Context, requestID string) (ready bool, err error)
}

// PairingRequestsHandler handles device pairing requests
type PairingRequestsHandler struct {
	k8sManager K8sManager
}

// NewPairingRequestsHandler creates a new pairing handler
func NewPairingRequestsHandler(k8sManager K8sManager) *PairingRequestsHandler {
	return &PairingRequestsHandler{
		k8sManager: k8sManager,
	}
}

// HandlePairDevice processes POST requests to the /pairing-requests endpoint
func (h *PairingRequestsHandler) HandlePairDevice(c *echo.Context) error {
	// Parse JSON request body
	var req models.PairingDeviceRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("error parsing request", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "Invalid request format",
			Status: models.PairingStatusError,
		})
	}

	// Validate RequestID field (trim whitespace and check non-empty)
	req.RequestID = strings.TrimSpace(req.RequestID)
	if req.RequestID == "" {
		slog.Warn("empty or whitespace-only requestId received")
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:  "requestId cannot be empty",
			Status: models.PairingStatusError,
		})
	}

	// Log successful request
	slog.Info("pairing request received", "request_id", req.RequestID)

	// Create ClawDevicePairingRequest CR
	if h.k8sManager != nil && h.k8sManager.IsEnabled() {
		if err := h.k8sManager.CreatePairingRequest(c.Request().Context(), req.RequestID); err != nil {
			slog.Error("failed to create pairing request CR", "error", err, "request_id", req.RequestID)
			return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:  "Something wrong happened, could not pair the device",
				Status: models.PairingStatusError,
			})
		}
		slog.Info("pairing request CR created", "request_id", req.RequestID)
	} else {
		slog.Warn("Kubernetes client not available, skipping CR creation", "request_id", req.RequestID)
	}

	// Return success response
	return c.JSON(http.StatusOK, models.PairingDeviceResponse{
		Status:  models.PairingStatusSuccess,
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
			Status: models.PairingStatusError,
		})
	}

	requestID = strings.TrimSpace(requestID)
	if requestID == "" {
		slog.Warn("empty request ID in path parameter")
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:  "request ID cannot be empty",
			Status: models.PairingStatusError,
		})
	}

	slog.Info("pairing status check", "request_id", requestID)

	if h.k8sManager != nil && h.k8sManager.IsEnabled() {
		ready, err := h.k8sManager.GetPairingRequestStatus(c.Request().Context(), requestID)
		if err != nil {
			slog.Error("failed to get pairing request status", "error", err, "request_id", requestID)
			return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error:  "Something wrong happened, could not get pairing status",
				Status: models.PairingStatusError,
			})
		}
		if ready {
			return c.JSON(http.StatusOK, models.PairingStatusResponse{
				Status: models.PairingStatusReady,
			})
		}
	}

	return c.JSON(http.StatusAccepted, models.PairingStatusResponse{
		Status: models.PairingStatusPending,
	})
}
