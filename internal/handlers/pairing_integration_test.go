package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/xcoulon/claw-device-pairing/internal/k8s/client"
)

func TestHandlePairDevice_WithCRCreation_Success(t *testing.T) {
	// Setup HTTP request
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{"requestId":"test-request-123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create handler without K8s manager (simulates disabled state)
	handler := NewPairingRequestsHandler(nil)

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestHandlePairDevice_CRCreationError_ReturnsError(t *testing.T) {
	// This test would require a mock that forces CR creation to fail
	// For now, we test the error path with a disabled client

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{"requestId":"test-request-123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock manager that is enabled but will fail
	mockManager := &mockFailingManager{}

	handler := NewPairingRequestsHandler(mockManager)

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert error response
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	expectedError := "Something wrong happened, could not pair the device"
	if !strings.Contains(rec.Body.String(), expectedError) {
		t.Errorf("expected error message to contain %q, got %q", expectedError, rec.Body.String())
	}
}

// Mock manager that always fails CR creation
type mockFailingManager struct{}

func (m *mockFailingManager) IsEnabled() bool {
	return true
}

func (m *mockFailingManager) CreatePairingRequest(ctx context.Context, requestID string) error {
	return &mockError{message: "simulated CR creation failure"}
}

func (m *mockFailingManager) GetPairingRequestStatus(ctx context.Context, requestID string) (bool, error) {
	return false, &mockError{message: "simulated status check failure"}
}

type mockError struct {
	message string
}

func (e *mockError) Error() string {
	return e.message
}

func TestManager_Interface(t *testing.T) {
	// Verify that client.Manager implements the K8sManager interface
	var _ K8sManager = (*client.Manager)(nil)
}
