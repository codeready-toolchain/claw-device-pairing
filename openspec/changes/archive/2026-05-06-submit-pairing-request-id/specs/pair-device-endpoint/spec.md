## RENAMED pair-device-endpoint → pairing-requests-endpoint

## MODIFIED Requirements

### Requirement: Endpoint Registration
The system SHALL register a POST endpoint at `/pairing-requests` that accepts JSON requests.

#### Scenario: Endpoint is registered
- **WHEN** server starts
- **THEN** POST /pairing-requests endpoint is available

#### Scenario: GET method not allowed
- **WHEN** GET request is made to /pairing-requests
- **THEN** response status is 405 Method Not Allowed

### Requirement: Handler Structure
The system SHALL implement the endpoint handler using a PairingRequestsHandler struct with constructor.

#### Scenario: Handler is instantiated
- **WHEN** server initializes
- **THEN** PairingRequestsHandler is created via NewPairingRequestsHandler constructor

#### Scenario: Handler method signature
- **WHEN** handler is called
- **THEN** handler method receives echo.Context and returns error
