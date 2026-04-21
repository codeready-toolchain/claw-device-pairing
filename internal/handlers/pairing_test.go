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
	req := httptest.NewRequest(http.MethodPost, "/pair-device", strings.NewReader(`{"id":"test-device"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingHandler()

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
	req := httptest.NewRequest(http.MethodPost, "/pair-device", strings.NewReader(`{invalid json}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingHandler()

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
	req := httptest.NewRequest(http.MethodPost, "/pair-device", strings.NewReader(`{}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"id cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}

func TestHandlePairDevice_EmptyIDField(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pair-device", strings.NewReader(`{"id":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"id cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}

func TestHandlePairDevice_WhitespaceOnlyID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/pair-device", strings.NewReader(`{"id":"   "}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewPairingHandler()

	// Execute
	if err := handler.HandlePairDevice(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedError := `{"error":"id cannot be empty","status":"error"}`
	actualBody := strings.TrimSpace(rec.Body.String())
	if actualBody != expectedError {
		t.Errorf("expected body %q, got %q", expectedError, actualBody)
	}
}
