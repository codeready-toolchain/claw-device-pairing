## ADDED Requirements

### Requirement: Create ClawDevicePairingRequest CR on pairing request
The system SHALL create a ClawDevicePairingRequest custom resource when HandlePairDevice receives a PairingRequest.

#### Scenario: Valid pairing request received
- **WHEN** HandlePairDevice receives a valid PairingRequest with requestId "abc-123"
- **THEN** system creates a ClawDevicePairingRequest CR named "abc-123"

#### Scenario: CR creation fails
- **WHEN** HandlePairDevice receives a PairingRequest and CR creation fails
- **THEN** system logs the detailed error and returns an error response to the UI

### Requirement: CR naming using requestId
The system SHALL name the ClawDevicePairingRequest CR using the incoming requestId field, sanitized for DNS-1123 compliance.

#### Scenario: RequestId is DNS-1123 compliant
- **WHEN** requestId is "valid-request-123"
- **THEN** CR is named "valid-request-123"

#### Scenario: RequestId requires sanitization
- **WHEN** requestId contains uppercase or invalid characters
- **THEN** CR name is sanitized to lowercase alphanumeric with hyphens, max 63 characters

### Requirement: Populate CR spec with RequestID
The system SHALL populate the ClawDevicePairingRequest spec.RequestID field with the incoming requestId value.

#### Scenario: RequestID field set
- **WHEN** creating CR for requestId "req-456"
- **THEN** spec.RequestID is set to "req-456"

### Requirement: Populate CR spec with Selector
The system SHALL populate the ClawDevicePairingRequest spec.Selector field with a label selector matching the current Claw instance.

#### Scenario: Instance label available
- **WHEN** pod has claw.sandbox.redhat.com/instance label with value "instance-1"
- **THEN** spec.Selector uses a selector matching "claw.sandbox.redhat.com/instance=instance-1"

#### Scenario: Instance label missing
- **WHEN** pod does not have claw.sandbox.redhat.com/instance label
- **THEN** system logs warning and sets spec.Selector to empty value or uses fallback selector

### Requirement: Retrieve namespace from environment variable
The system SHALL read the target namespace for CR creation from the NAMESPACE environment variable.

#### Scenario: Namespace environment variable set
- **WHEN** NAMESPACE environment variable is "claw-prod"
- **THEN** CR is created in "claw-prod" namespace

#### Scenario: Namespace environment variable missing
- **WHEN** NAMESPACE environment variable is not set
- **THEN** system defaults to "default" namespace and logs warning

### Requirement: Retrieve instance label from pod metadata
The system SHALL retrieve the claw.sandbox.redhat.com/instance label value from pod metadata via environment variable.

#### Scenario: Instance label environment variable set
- **WHEN** CLAW_INSTANCE environment variable is "prod-instance-1"
- **THEN** system uses "prod-instance-1" for selector construction

#### Scenario: Instance label environment variable missing
- **WHEN** CLAW_INSTANCE environment variable is not set
- **THEN** system logs warning and handles gracefully per spec.Selector requirements

### Requirement: Initialize Kubernetes client at startup
The system SHALL initialize the Kubernetes in-cluster client at server startup, not per-request.

#### Scenario: Server starts in cluster
- **WHEN** server starts with in-cluster configuration available
- **THEN** Kubernetes client is initialized and ready for CR creation

#### Scenario: Server starts outside cluster
- **WHEN** server starts without in-cluster configuration
- **THEN** client initialization fails gracefully and CR creation is skipped with appropriate logging

### Requirement: Return succinct error message to UI on CR creation failure
The system SHALL return a succinct, user-friendly error message to the UI when CR creation fails, while logging the detailed error server-side.

#### Scenario: CR creation fails with Kubernetes API error
- **WHEN** CR creation fails with error "etcdserver: request timed out"
- **THEN** system logs "Failed to create ClawDevicePairingRequest: etcdserver: request timed out" and returns "Something wrong happened, could not pair the device" to UI

#### Scenario: CR creation fails due to RBAC
- **WHEN** CR creation fails with "forbidden: User cannot create resource"
- **THEN** system logs the full RBAC error and returns "Something wrong happened, could not pair the device" to UI
