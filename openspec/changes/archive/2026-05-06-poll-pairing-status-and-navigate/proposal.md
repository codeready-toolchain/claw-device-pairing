## Why

After submitting a pairing request, the UI has no way to know when the backend has completed the pairing process. Users cannot proceed to OpenClaw until pairing is approved, creating a poor user experience with no clear next step or feedback mechanism.

## What Changes

- Add "Go to OpenClaw" button below the progress stepper (initially disabled)
- Poll pairing status endpoint after submission to detect when pairing is approved
- Enable navigation button when backend confirms pairing is complete
- Navigate to OpenClaw when user clicks the enabled button
- Add backend GET endpoint to check pairing request status
- Handle polling timeout (30s) if pairing doesn't complete

## Capabilities

### New Capabilities
- `pairing-status-polling`: Poll the pairing request status endpoint to detect when pairing is approved, with automatic retry logic and timeout handling
- `navigation-to-openclaw`: Display navigation button and handle browser navigation to OpenClaw after successful pairing

### Modified Capabilities
- `pairing-ui`: Add "Go to OpenClaw" button component below the progress stepper
- `pairing-requests-endpoint`: Add GET endpoint to retrieve pairing request status
- `pairing-submission`: Integrate status polling after pairing request submission

## Impact

- `ui/src/App.jsx` - add button component and polling logic
- `internal/handlers/pairing.go` - add GET handler for status retrieval
- `cmd/main.go` - register new GET route for status endpoint
- `internal/models/pairing.go` - add status response model
- User experience - clear feedback on pairing approval and seamless navigation to OpenClaw
- Backend API - new GET `/pairing-requests/:id` endpoint
