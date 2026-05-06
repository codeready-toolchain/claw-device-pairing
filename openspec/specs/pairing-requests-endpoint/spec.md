## Purpose

This specification defines the HTTP endpoint for submitting device pairing requests to the backend. It establishes the endpoint path, request/response format, validation rules, and handler structure.

## Requirements

### Requirement: Endpoint Registration
The system SHALL register a POST endpoint at `/pairing-requests` that accepts JSON requests.

#### Scenario: Endpoint is registered
- **WHEN** server starts
- **THEN** POST /pairing-requests endpoint is available

#### Scenario: GET method not allowed
- **WHEN** GET request is made to /pairing-requests
- **THEN** response status is 405 Method Not Allowed

### Requirement: Request Model
The system SHALL accept JSON requests with a `requestId` field of type string.

#### Scenario: Valid JSON with requestId field
- **WHEN** POST request contains valid JSON with `{"requestId":"request-123"}`
- **THEN** request is parsed successfully

#### Scenario: Missing requestId field
- **WHEN** POST request contains JSON without requestId field
- **THEN** response status is 400 Bad Request with error message

### Requirement: Request Validation
The system SHALL validate that the requestId field is non-empty.

#### Scenario: Empty requestId field
- **WHEN** POST request contains `{"requestId":""}`
- **THEN** response status is 400 Bad Request with error message "requestId cannot be empty"

#### Scenario: Whitespace-only requestId field
- **WHEN** POST request contains `{"requestId":"   "}`
- **THEN** response status is 400 Bad Request with error message "requestId cannot be empty"

#### Scenario: Valid non-empty requestId
- **WHEN** POST request contains `{"requestId":"request-xyz"}`
- **THEN** request passes validation

### Requirement: JSON Parsing
The system SHALL return 400 Bad Request for invalid JSON payloads.

#### Scenario: Malformed JSON
- **WHEN** POST request body is not valid JSON
- **THEN** response status is 400 Bad Request with error "Invalid request format"

#### Scenario: Empty request body
- **WHEN** POST request body is empty
- **THEN** response status is 400 Bad Request

### Requirement: Success Response
The system SHALL return 200 OK with JSON response for valid pairing requests.

#### Scenario: Successful pairing request
- **WHEN** POST request is valid with non-empty requestId
- **THEN** response status is 200 OK and body contains `{"status":"success","message":"pairing request received"}`

### Requirement: Error Response Format
The system SHALL return consistent JSON error responses with error and status fields.

#### Scenario: Error response structure
- **WHEN** request fails validation
- **THEN** response body contains `{"error":"<error message>","status":"error"}`

#### Scenario: Error is logged
- **WHEN** request fails validation or parsing
- **THEN** error details are logged with slog

### Requirement: Handler Structure
The system SHALL implement the endpoint handler using a PairingRequestsHandler struct with constructor.

#### Scenario: Handler is instantiated
- **WHEN** server initializes
- **THEN** PairingRequestsHandler is created via NewPairingRequestsHandler constructor

#### Scenario: Handler method signature
- **WHEN** handler is called
- **THEN** handler method receives echo.Context and returns error

### Requirement: Content-Type Validation
The system SHALL require Content-Type: application/json header for POST requests.

#### Scenario: Missing Content-Type header
- **WHEN** POST request has no Content-Type header
- **THEN** request is rejected with 400 Bad Request

#### Scenario: Wrong Content-Type
- **WHEN** POST request has Content-Type other than application/json
- **THEN** request is rejected with 400 Bad Request

### Requirement: Status Retrieval Endpoint
The system SHALL register a GET endpoint at `/pairing-requests/:id` to retrieve pairing request status.

#### Scenario: GET endpoint is registered
- **WHEN** server starts
- **THEN** GET /pairing-requests/:id endpoint is available

#### Scenario: Valid request ID parameter
- **WHEN** GET request is made with a valid request ID
- **THEN** the handler extracts the ID from the URL path parameter

#### Scenario: Return pending status
- **WHEN** pairing request is pending approval
- **THEN** response status is 202 No Content with body `{"status":"pending"}`

#### Scenario: Return approved status
- **WHEN** pairing request has been approved
- **THEN** response status is 200 OK with body `{"status":"approved"}`

#### Scenario: Handle missing request ID
- **WHEN** GET request is made without request ID parameter
- **THEN** response status is 404 Not Found

#### Scenario: Handle invalid request ID
- **WHEN** GET request is made with non-existent request ID
- **THEN** response status is 404 Not Found with error message

### Requirement: Status Response Model
The system SHALL return consistent JSON responses for status queries.

#### Scenario: Pending response structure
- **WHEN** pairing request is pending
- **THEN** response body contains `{"status":"pending"}`

#### Scenario: Approved response structure
- **WHEN** pairing request is approved
- **THEN** response body contains `{"status":"approved"}`

#### Scenario: Error response for invalid ID
- **WHEN** request ID is not found
- **THEN** response body contains `{"error":"request not found","status":"error"}`
