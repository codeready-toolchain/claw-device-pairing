## ADDED Requirements

### Requirement: Automatic Handshake on Page Load
The application SHALL automatically initiate the WebSocket handshake when the App component mounts.

#### Scenario: Handshake triggered on mount
- **WHEN** the App component mounts
- **THEN** performHandshake is called automatically

#### Scenario: Handshake uses configured gateway URL
- **WHEN** the handshake is initiated
- **THEN** the gateway URL is read from VITE_GATEWAY_URL environment variable or defaults to ws://localhost:8080/ws

### Requirement: Handshake State Tracking
The application SHALL track the handshake status using React state.

#### Scenario: Initial state is idle
- **WHEN** the component first renders
- **THEN** handshakeStatus is 'idle'

#### Scenario: Loading state during handshake
- **WHEN** the handshake is in progress
- **THEN** handshakeStatus is 'loading'

#### Scenario: Success state after handshake
- **WHEN** the handshake completes successfully
- **THEN** handshakeStatus is 'success' and deviceToken is stored in state

#### Scenario: Error state on failure
- **WHEN** the handshake fails
- **THEN** handshakeStatus is 'error' and handshakeError contains the error message

### Requirement: Progress Step Completion
The first ProgressStep SHALL be marked as complete when the handshake succeeds.

#### Scenario: Step 1 success variant
- **WHEN** handshakeStatus is 'success'
- **THEN** the first ProgressStep has variant="success"

#### Scenario: Step 1 pending before success
- **WHEN** handshakeStatus is not 'success'
- **THEN** the first ProgressStep does not have variant="success"

### Requirement: Error Display
The application SHALL display error messages when the handshake fails.

#### Scenario: Error message shown
- **WHEN** handshakeStatus is 'error'
- **THEN** the error message is displayed in the CardBody

#### Scenario: Error replaces stepper
- **WHEN** handshakeStatus is 'error'
- **THEN** the ProgressStepper is not rendered

### Requirement: Loading State Display
The application SHALL indicate loading state while the handshake is in progress.

#### Scenario: Loading indicator shown
- **WHEN** handshakeStatus is 'loading'
- **THEN** a loading indicator or message is displayed

#### Scenario: Stepper hidden during loading
- **WHEN** handshakeStatus is 'loading'
- **THEN** the ProgressStepper may be hidden or show pending state

### Requirement: Connection Parameters
The handshake SHALL use appropriate connection parameters for device pairing.

#### Scenario: Client ID is set
- **WHEN** performHandshake is called
- **THEN** clientId is set to "device-pairing-ui"

#### Scenario: Role is operator
- **WHEN** performHandshake is called
- **THEN** role is set to "operator"

#### Scenario: Scopes are provided
- **WHEN** performHandshake is called
- **THEN** scopes array is set appropriately (may be empty for pairing flow)

#### Scenario: Client version is set
- **WHEN** performHandshake is called
- **THEN** clientVersion is set to a valid version string
