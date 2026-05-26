# auto-handshake Specification

## Purpose
TBD - created by archiving change integrate-handshake-with-ui. Update Purpose after archive.
## Requirements
### Requirement: Automatic Handshake on Page Load
The application SHALL automatically initiate the WebSocket handshake when the App component mounts.

#### Scenario: Handshake triggered on mount
- **WHEN** the App component mounts
- **THEN** performHandshake is called automatically

#### Scenario: Handshake uses configured gateway URL
- **WHEN** the handshake is initiated
- **THEN** the gateway URL is read from VITE_GATEWAY_URL environment variable or defaults to ws://localhost:8080/ws

### Requirement: Handshake State Tracking
The application SHALL track the handshake status using React state. When the handshake completes successfully (device already paired), the application SHALL also mark the pairing and approval as complete.

#### Scenario: Initial state is idle
- **WHEN** the component first renders
- **THEN** handshakeStatus is 'idle'

#### Scenario: Loading state during handshake
- **WHEN** the handshake is in progress
- **THEN** handshakeStatus is 'loading'

#### Scenario: Success state after handshake with already-paired device
- **WHEN** the handshake completes successfully (ok: true, device already paired)
- **THEN** handshakeStatus is 'success', pairingStatus is 'success', and approvalStatus is 'approved'

#### Scenario: Error state on failure
- **WHEN** the handshake fails
- **THEN** handshakeStatus is 'error' and handshakeError contains the error message

### Requirement: Progress Step Completion
The ProgressSteps SHALL reflect the combined handshake and pairing state. When the device is already paired, both steps SHALL show success immediately.

#### Scenario: Both steps succeed for already-paired device
- **WHEN** the handshake completes successfully (device already paired)
- **THEN** the first ProgressStep has variant="success" and the second ProgressStep has variant="success"

#### Scenario: Step 1 success variant
- **WHEN** handshakeStatus is 'success'
- **THEN** the first ProgressStep has variant="success"

#### Scenario: Step 1 pending before success
- **WHEN** handshakeStatus is not 'success'
- **THEN** the first ProgressStep does not have variant="success"

### Requirement: Navigation Button for Already-Paired Device
The "Go to OpenClaw" button SHALL be enabled immediately when the device is already paired.

#### Scenario: Button enabled for already-paired device
- **WHEN** the handshake completes successfully (device already paired)
- **THEN** the "Go to OpenClaw" button is enabled (approvalStatus is 'approved')

#### Scenario: Button disabled when pairing not complete
- **WHEN** approvalStatus is not 'approved'
- **THEN** the "Go to OpenClaw" button is disabled

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
The handshake SHALL use appropriate connection parameters for device pairing, including optional authentication token from URL fragment.

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

#### Scenario: Token extracted from URL fragment without prefix
- **WHEN** the URL fragment is "#token=abc123"
- **THEN** the token value "abc123" is passed to performHandshake (without the "token=" prefix)

#### Scenario: Token extracted from URL fragment with hash only
- **WHEN** the URL fragment is "#abc123"
- **THEN** the token value "abc123" is passed to performHandshake

#### Scenario: Token extracted from URL fragment with slash
- **WHEN** the URL fragment is "#/token=abc123"
- **THEN** the token value "abc123" is passed to performHandshake (without the "token=" prefix)

#### Scenario: No token when fragment is empty
- **WHEN** the URL has no fragment
- **THEN** performHandshake is called with undefined token value

