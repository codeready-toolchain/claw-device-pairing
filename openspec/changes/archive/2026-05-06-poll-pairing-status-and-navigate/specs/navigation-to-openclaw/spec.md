## ADDED Requirements

### Requirement: Display Navigation Button
The system SHALL display a "Go to OpenClaw" button below the progress stepper.

#### Scenario: Button is initially disabled
- **WHEN** the page loads
- **THEN** the "Go to OpenClaw" button is displayed but disabled

#### Scenario: Button placement
- **WHEN** the UI renders
- **THEN** the button appears below the ProgressStepper component

#### Scenario: Button label
- **WHEN** the button is rendered
- **THEN** the button text reads "Go to OpenClaw"

### Requirement: Enable Button on Approval
The system SHALL enable the navigation button when pairing is approved.

#### Scenario: Enable button on approval
- **WHEN** polling receives HTTP 200 OK response indicating pairing approval
- **THEN** the "Go to OpenClaw" button becomes enabled

#### Scenario: Button remains disabled while pending
- **WHEN** pairing status is pending (HTTP 202)
- **THEN** the button remains disabled

#### Scenario: Button remains disabled on timeout
- **WHEN** polling times out after 30 seconds
- **THEN** the button remains disabled

### Requirement: Navigate to OpenClaw
The system SHALL navigate to OpenClaw when the enabled button is clicked.

#### Scenario: Navigate on button click
- **WHEN** user clicks the enabled "Go to OpenClaw" button
- **THEN** browser navigates to the OpenClaw root URL

#### Scenario: Preserve authentication token
- **WHEN** navigating to OpenClaw
- **THEN** the URL fragment with authentication token is preserved

#### Scenario: Remove pairing path
- **WHEN** constructing the navigation URL
- **THEN** the `/integration/device-pairing` path is removed

#### Scenario: Construct root URL with token
- **WHEN** building the navigation URL
- **THEN** the URL is constructed as `${protocol}//${host}#token=${tokenValue}`

### Requirement: Handle Missing Token
The system SHALL handle navigation when authentication token is missing.

#### Scenario: Navigate without token
- **WHEN** no authentication token is present in the URL fragment
- **THEN** navigation proceeds to root URL without fragment

#### Scenario: Log warning for missing token
- **WHEN** token is missing during navigation
- **THEN** a warning is logged to console for debugging
