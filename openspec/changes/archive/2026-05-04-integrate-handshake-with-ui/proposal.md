## Why

The device pairing UI has a progress stepper showing two steps (Generate device id, Pair device with OpenClaw), but currently no logic connects the handshake implementation to the UI. Users need to see automatic progression through the pairing flow when they load the page.

## What Changes

- Trigger the WebSocket handshake automatically when the App component mounts
- Track handshake status in React state (loading, success, error)
- Update the ProgressStepper to mark the first step as complete when device identity is successfully generated and handshake succeeds
- Display loading state during handshake and error messages if handshake fails
- Configure gateway WebSocket URL (from environment variable or default)

## Capabilities

### New Capabilities
- `auto-handshake`: Automatic handshake triggering on page load with status tracking and UI updates

### Modified Capabilities
- `pairing-ui`: Update to show handshake status and mark progress stepper steps as complete based on handshake state

## Impact

- Modifies `ui/src/App.jsx` to integrate handshake logic with useState and useEffect hooks
- Imports and uses `performHandshake` from `ui/src/handshake.ts`
- Updates ProgressStepper component to accept and display step completion state
- Adds error handling and display for handshake failures
- May require environment variable or configuration for gateway URL
