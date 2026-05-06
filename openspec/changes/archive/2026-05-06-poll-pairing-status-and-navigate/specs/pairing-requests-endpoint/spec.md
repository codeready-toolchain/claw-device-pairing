## ADDED Requirements

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
