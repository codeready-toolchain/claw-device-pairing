## ADDED Requirements

### Requirement: Initiate Status Polling After Submission
The system SHALL start polling for pairing approval status after successful pairing request submission.

#### Scenario: Start polling on successful submission
- **WHEN** pairing request submission receives any response (2xx or 4xx)
- **THEN** status polling begins immediately

#### Scenario: Pass request ID to polling
- **WHEN** initiating status polling
- **THEN** the pairing request ID is provided to the polling mechanism

#### Scenario: Polling state reflects submission
- **WHEN** polling is initiated
- **THEN** pairing status state is set to indicate approval is pending
