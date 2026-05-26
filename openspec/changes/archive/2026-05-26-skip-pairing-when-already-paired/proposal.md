## Why

When the handshake completes successfully (`ok: true`), the device is already paired with the server. Currently the UI only handles the `NOT_PAIRED` error path and leaves the pairing step idle on success, creating a confusing dead-end for already-paired devices.

## What Changes

- When `performHandshake` resolves successfully (no error thrown), skip the pairing submission and polling steps entirely.
- Set the "Pair device with OpenClaw" progress step to `success` immediately.
- Enable the "Go to OpenClaw" button so the user can proceed without waiting.

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `auto-handshake`: Handle the successful handshake response (`ok: true`) by marking pairing as complete and enabling navigation.

## Impact

- `ui/src/App.jsx`: The handshake success path in the first `useEffect` needs to also update `pairingStatus` and `approvalStatus` when the device is already paired.
- No backend changes required — the server already returns `ok: true` for paired devices.
- No new dependencies.
