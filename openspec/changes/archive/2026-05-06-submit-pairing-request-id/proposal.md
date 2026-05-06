## Why

When the WebSocket handshake completes with a `NOT_PAIRED` error, the server includes a pairing request ID in `error.details.requestId` that must be submitted to the backend to complete device pairing. Currently, the UI does not extract or submit this request ID, preventing successful device pairing when devices are not yet paired.

## What Changes

- Extract `requestId` from `NOT_PAIRED` error response in the handshake error handler
- Rename the `/pair-device` endpoint to `/pairing-requests` on the backend
- Rename the `PairingHandler` and related tests to match the new endpoint name
- Send the `requestId` to the `/pairing-requests` POST endpoint on the backend
- Handle the pairing submission response and update UI state accordingly
- Display pairing status to the user (success, pending, or error)

## Capabilities

### New Capabilities
- `pairing-submission`: Extract pairing request ID from handshake error and submit it to the backend pairing endpoint

### Modified Capabilities
- `auto-handshake`: Handle `NOT_PAIRED` error and extract the `requestId` from error details
- `pairing-ui`: Display pairing submission status and handle POST request to `/pair-device`

## Impact

- `ui/src/App.jsx` - handshake error handling and pairing submission logic
- `ui/src/handshake.ts` - error response parsing (if needed)
- `cmd/main.go` - rename endpoint route from `/pair-device` to `/pairing-requests`
- Backend handler - rename `PairingHandler` and related functions
- Backend tests - update test names and endpoint references
- User experience - enables actual device pairing workflow completion
