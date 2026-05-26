package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/xcoulon/claw-device-pairing/internal/version"
)

func TestHandleHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := HandleHealth(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var resp map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if resp["status"] != "ok" {
		t.Errorf("expected status %q, got %q", "ok", resp["status"])
	}
	if resp["commit_hash"] != version.CommitHash {
		t.Errorf("expected commit_hash %q, got %q", version.CommitHash, resp["commit_hash"])
	}
	if resp["build_time"] != version.BuildTime {
		t.Errorf("expected build_time %q, got %q", version.BuildTime, resp["build_time"])
	}
}
