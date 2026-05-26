package models

// PairingDeviceRequest represents the JSON request body for device pairing
type PairingDeviceRequest struct {
	RequestID string `json:"requestId"`
}

// PairingDeviceResponse represents the success response for device pairing
type PairingDeviceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PairingStatusResponse represents the response for pairing status queries
type PairingStatusResponse struct {
	Status string `json:"status"`
}

const (
	PairingStatusSuccess = "success"
	PairingStatusError   = "error"
	PairingStatusPending = "pending"
	PairingStatusReady   = "ready"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

// HealthResponse represents the response for health check queries
type HealthResponse struct {
	Status     string `json:"status"`
	CommitHash string `json:"commit_hash"`
	BuildTime  string `json:"build_time"`
}
