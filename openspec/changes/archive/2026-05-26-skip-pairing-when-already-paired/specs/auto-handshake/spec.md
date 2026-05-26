## MODIFIED Requirements

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
