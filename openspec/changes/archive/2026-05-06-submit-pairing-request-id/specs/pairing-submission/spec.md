## ADDED Requirements

### Requirement: Extract Pairing Request ID from Error
The system SHALL extract the pairing request ID from NOT_PAIRED error responses.

#### Scenario: Extract requestId from error details
- **WHEN** handshake fails with a NOT_PAIRED error
- **THEN** the requestId is extracted from error.details.requestId

#### Scenario: Handle missing requestId
- **WHEN** handshake fails with a NOT_PAIRED error but requestId is missing from error.details
- **THEN** the system displays an error message indicating incomplete error information

### Requirement: Submit Pairing Request to Backend
The system SHALL send the pairing request ID to the /pairing-requests endpoint via HTTP POST.

#### Scenario: POST to /pairing-requests endpoint
- **WHEN** a pairing request ID is extracted from the error
- **THEN** an HTTP POST request is made to /pairing-requests with the requestId in the request body

#### Scenario: Include requestId in request body
- **WHEN** making the POST request to /pairing-requests
- **THEN** the request body contains JSON with requestId field

#### Scenario: Set appropriate content type
- **WHEN** making the POST request to /pairing-requests
- **THEN** the Content-Type header is set to application/json

### Requirement: Handle Pairing Submission Response
The system SHALL handle success and failure responses from the pairing submission.

#### Scenario: Handle successful pairing submission
- **WHEN** /pairing-requests responds with 200 OK
- **THEN** the system displays a success state indicating pairing is complete

#### Scenario: Handle pairing submission failure
- **WHEN** /pairing-requests responds with an error status
- **THEN** the system displays an error message with details

#### Scenario: Handle network errors
- **WHEN** the POST request to /pairing-requests fails due to network issues
- **THEN** the system displays an error message indicating connection problems

### Requirement: Display Pairing Status
The system SHALL provide visual feedback on the pairing submission status.

#### Scenario: Show pending state during submission
- **WHEN** the pairing request is being submitted
- **THEN** a loading indicator is displayed

#### Scenario: Show success state after completion
- **WHEN** the pairing submission succeeds
- **THEN** a success indicator is displayed

#### Scenario: Show error state on failure
- **WHEN** the pairing submission fails
- **THEN** an error message is displayed with details
