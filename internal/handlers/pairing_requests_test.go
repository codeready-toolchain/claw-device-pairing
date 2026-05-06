package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
)

func TestHandlePairDevice_ValidRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{"requestId":"test-request-123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingRequestsHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedBody := `{"status":"success","message":"pairing request received"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, actualBody)
	}
}

func TestHandlePairDevice_InvalidJSON(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{invalid json}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingRequestsHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestHandlePairDevice_MissingIDField(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingRequestsHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"requestId cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}

func TestHandlePairDevice_EmptyIDField(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{"requestId":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingRequestsHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"requestId cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}

func TestHandlePairDevice_WhitespaceOnlyID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pairing-requests", strings.NewReader(`{"requestId":"   "}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingRequestsHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"requestId cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}

func TestHandleGetPairingStatus_ValidRequest(t *testing.T) {
	// Setup
	e := echo.New()
	handler := NewPairingRequestsHandler()
	e.GET("/pairing-requests/:id", handler.HandleGetPairingStatus)

	req := httptest.NewRequest(http.MethodGet, "/pairing-requests/test-request-123", nil)
	rec := httptest.NewRecorder()

	// Execute
	e.ServeHTTP(rec, req)

	// Assert
	if rec.Code != http.StatusAccepted {
		t.Errorf("expected status %d, got %d", http.StatusAccepted, rec.Code)
	}

	expectedBody := `{"status":"pending"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, actualBody)
	}
}

func TestHandleGetPairingStatus_MissingID(t *testing.T) {
	// Setup
	e := echo.New()
	handler := NewPairingRequestsHandler()
	e.GET("/pairing-requests/:id", handler.HandleGetPairingStatus)

	// Request with empty path parameter
	req := httptest.NewRequest(http.MethodGet, "/pairing-requests/", nil)
	rec := httptest.NewRecorder()

	// Execute
	e.ServeHTTP(rec, req)

	// Assert - Echo will return 404 for unmatched route
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}
