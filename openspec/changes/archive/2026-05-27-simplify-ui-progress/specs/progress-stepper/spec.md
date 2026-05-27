## REMOVED Requirements

### Requirement: Stepper Component Display
**Reason**: The ProgressStepper is replaced by a simpler Spinner + status label. See the new `spinner-status` capability.
**Migration**: Use Spinner component from `@patternfly/react-core` with a text label below it.

### Requirement: Step Definitions
**Reason**: Steps are no longer displayed as discrete stepper items. The current step is communicated via the status label text.
**Migration**: Step information is conveyed through the status label (e.g., "Generating device id...", "Pairing device with OpenClaw...").

### Requirement: Step State Tracking
**Reason**: Step index tracking is no longer needed. The existing `handshakeStatus`, `pairingStatus`, and `approvalStatus` states drive the label text directly.
**Migration**: No additional state variable needed; derive the label from existing state.

### Requirement: Patternfly Step Components
**Reason**: ProgressStep components are removed along with the ProgressStepper.
**Migration**: Remove ProgressStep and ProgressStepper imports from `@patternfly/react-core`.
