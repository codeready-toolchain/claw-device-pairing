## ADDED Requirements

### Requirement: Spinner Display
The card body SHALL display a PatternFly Spinner component centered within the card during all loading and in-progress states.

#### Scenario: Spinner is visible during handshake
- **WHEN** the handshake is in progress (loading state)
- **THEN** a PatternFly `Spinner` component is displayed in the card body

#### Scenario: Spinner is visible during pairing
- **WHEN** the pairing status is pending or progressing
- **THEN** the Spinner component remains displayed

#### Scenario: Spinner is hidden on success
- **WHEN** the pairing approval is received (approvalStatus is 'approved')
- **THEN** the Spinner is still displayed (briefly, before auto-redirect)

#### Scenario: Spinner is hidden on error
- **WHEN** an error occurs during handshake or pairing
- **THEN** the Spinner is not displayed

### Requirement: Status Label
The card body SHALL display a text label below the Spinner that describes the current step.

#### Scenario: Label during handshake
- **WHEN** the handshake is in progress
- **THEN** the label text reads "Generating device id..."

#### Scenario: Label during pairing submission
- **WHEN** the pairing request is being submitted (pending status)
- **THEN** the label text reads "Pairing device with OpenClaw..."

#### Scenario: Label during pairing polling
- **WHEN** the pairing status is being polled (progressing status)
- **THEN** the label text reads "Pairing device with OpenClaw..."

#### Scenario: Label on success
- **WHEN** the pairing is approved
- **THEN** the label text reads "Redirecting to OpenClaw..."

#### Scenario: Label on error
- **WHEN** an error occurs
- **THEN** the label displays the error message text
