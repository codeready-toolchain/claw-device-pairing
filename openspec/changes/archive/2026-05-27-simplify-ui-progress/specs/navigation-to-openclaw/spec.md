## REMOVED Requirements

### Requirement: Display Navigation Button
**Reason**: Replaced by automatic redirect on pairing approval. The "Go to OpenClaw" button is no longer needed.
**Migration**: Navigation is now triggered automatically via a useEffect when approvalStatus becomes 'approved'.

### Requirement: Enable Button on Approval
**Reason**: No button exists to enable. Navigation is automatic.
**Migration**: The approval state still triggers navigation, but via auto-redirect instead of button enablement.

## MODIFIED Requirements

### Requirement: Navigate to OpenClaw
The system SHALL automatically navigate to OpenClaw when pairing is approved, without user interaction.

#### Scenario: Auto-redirect on approval
- **WHEN** polling receives HTTP 200 OK response indicating pairing approval
- **THEN** the browser automatically navigates to the OpenClaw root URL

#### Scenario: Preserve authentication token
- **WHEN** navigating to OpenClaw
- **THEN** the URL fragment with authentication token is preserved

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
