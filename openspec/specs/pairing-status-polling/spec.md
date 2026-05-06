## Purpose

This specification defines the polling mechanism for checking pairing request approval status. It establishes polling intervals, response handling, timeout behavior, and cleanup requirements for monitoring when the backend approves a pairing request.

## Requirements

### Requirement: Poll Pairing Status
The system SHALL poll the pairing request status endpoint at regular intervals after pairing submission.

#### Scenario: Start polling after submission
- **WHEN** pairing request is successfully submitted
- **THEN** polling begins with 1-second intervals

#### Scenario: Poll with request ID
- **WHEN** each polling request is made
- **THEN** the request includes the pairing request ID in the URL path

#### Scenario: Use relative URL for polling
- **WHEN** constructing the polling endpoint URL
- **THEN** the URL uses a relative path `pairing-requests/:id`

### Requirement: Handle Polling Responses
The system SHALL interpret HTTP status codes to determine pairing status.

#### Scenario: Continue polling on 202 response
- **WHEN** polling request returns HTTP 202 No Content
- **THEN** polling continues after the interval delay

#### Scenario: Stop polling on 200 response
- **WHEN** polling request returns HTTP 200 OK
- **THEN** polling stops and pairing is marked as approved

#### Scenario: Handle polling errors
- **WHEN** polling request fails with network error or HTTP error status
- **THEN** polling stops and error is displayed to user

### Requirement: Polling Timeout
The system SHALL enforce a maximum polling duration.

#### Scenario: Stop polling after timeout
- **WHEN** 30 seconds have elapsed since polling started
- **THEN** polling stops automatically

#### Scenario: Display timeout message
- **WHEN** polling timeout occurs
- **THEN** user is shown a message indicating pairing approval is pending

### Requirement: Cleanup Polling on Unmount
The system SHALL stop polling when the component unmounts.

#### Scenario: Clear interval on unmount
- **WHEN** the React component unmounts
- **THEN** the polling interval is cleared

#### Scenario: Prevent memory leaks
- **WHEN** polling is active and component unmounts
- **THEN** no further polling requests are made
