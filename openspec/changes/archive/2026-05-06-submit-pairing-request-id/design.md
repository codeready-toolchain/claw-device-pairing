## Context

The device pairing workflow requires a two-step process:
1. WebSocket handshake attempt (which fails with `NOT_PAIRED` for unpaired devices)
2. HTTP POST to `/pairing-requests` with the pairing request ID from the error

Currently, the UI performs step 1 but doesn't handle the `NOT_PAIRED` error or perform step 2. When the handshake fails with `NOT_PAIRED`, the error contains `error.details.requestId` that identifies the pairing request on the backend.

The backend currently has a `/pair-device` endpoint that should be renamed to `/pairing-requests` to better reflect that it handles pairing request submissions (rather than directly pairing a device).

## Goals / Non-Goals

**Goals:**
- Rename `/pair-device` endpoint to `/pairing-requests` for clearer semantics
- Extract `requestId` from `NOT_PAIRED` error in the handshake error handler
- Submit the request ID to the `/pairing-requests` POST endpoint
- Display pairing submission status to the user
- Handle pairing success, failure, and pending states

**Non-Goals:**
- Modifying the WebSocket handshake protocol or error format
- Changing the backend endpoint implementation logic (only the name/route)
- Implementing automatic retry logic for failed pairing requests
- Supporting multiple concurrent pairing requests

## Decisions

### Decision 1: Rename endpoint from /pair-device to /pairing-requests

**Rationale**: The current name `/pair-device` is ambiguous - it suggests directly pairing a device, but the endpoint actually receives a pairing request ID to complete a pairing request. The name `/pairing-requests` better reflects that this endpoint handles pairing request submissions.

**Alternatives considered**:
- Keep `/pair-device`: Less clear semantics, implies direct device pairing
- Use `/pairing-requests/:id`: RESTful but unnecessary since we're POSTing the ID in the body
- Use `/pair`: Too generic, doesn't convey what's being submitted

**Decision**: Rename the endpoint route to `/pairing-requests`, rename the handler to `PairingRequestsHandler`, and update all references in tests and documentation.

### Decision 2: Handle NOT_PAIRED error specifically in App.jsx

**Rationale**: The pairing submission is UI-specific logic that depends on the error type. The handshake module should remain focused on the WebSocket protocol, while the UI layer handles the pairing workflow.

**Alternatives considered**:
- Handling in handshake.ts: Would mix protocol concerns with UI workflow concerns
- Creating a separate pairing service: Over-engineering for a simple POST request

**Decision**: Catch the handshake error in App.jsx, check for `NOT_PAIRED` error code, and extract the `requestId` from the error object.

### Decision 3: Use fetch API for /pairing-requests POST request

**Rationale**: The `/pairing-requests` endpoint is a simple HTTP POST. Using fetch keeps dependencies minimal and is consistent with standard web APIs.

**Alternatives considered**:
- Adding axios or another HTTP library: Unnecessary dependency for a single POST request
- Reusing WebSocket connection: Wrong protocol for this operation

**Decision**: Use `fetch('/pairing-requests', { method: 'POST', body: JSON.stringify({ requestId }) })` directly in App.jsx.

### Decision 4: Update progress stepper to show pairing status

**Rationale**: The UI already has a progress stepper with two steps. The second step ("Pair device with OpenClaw") should reflect the pairing submission status.

**Implementation**: Update the second ProgressStep variant based on pairing submission state (pending, success, error).

## Risks / Trade-offs

**[Risk]** POST to /pairing-requests might fail due to network issues or backend errors
→ **Mitigation**: Display error message to user, allow manual retry by refreshing the page. The pairing request ID persists on the backend, so retry is safe.

**[Risk]** User might close the page before pairing completes
→ **Mitigation**: Acceptable for MVP. The backend maintains the pairing request state, and the device can attempt handshake again later.

**[Trade-off]** Tight coupling between handshake error handling and pairing submission
→ **Acceptable**: This is the intended workflow. The coupling is explicit and well-defined by the protocol.

**[Risk]** Error object structure might vary (requestId might not always be present)
→ **Mitigation**: Check for `error.details?.requestId` existence before attempting to use it. If missing, show generic error message.
