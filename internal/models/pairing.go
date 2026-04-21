package models

// PairDeviceRequest represents the JSON request body for device pairing
type PairDeviceRequest struct {
	ID string `json:"id"`
}

// PairDeviceResponse represents the success response for device pairing
type PairDeviceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}
