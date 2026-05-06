## Context

The device pairing workflow currently has a gap: after submitting a pairing request to `/pairing-requests`, the UI has no mechanism to determine when the backend has approved the pairing. The backend processes pairing asynchronously (likely requiring administrator approval), but the UI provides no feedback or next step.

Users need to know:
1. When their pairing request has been approved
2. How to proceed to OpenClaw once approved

The current UI state after pairing submission:
- Progress stepper shows "Pair device with OpenClaw" step as complete
- No indication of approval status
- No way to navigate to OpenClaw
- User must manually navigate or refresh

## Goals / Non-Goals

**Goals:**
- Poll backend for pairing status after submission
- Display "Go to OpenClaw" button for clear next action
- Enable button only when pairing is approved
- Navigate to OpenClaw preserving authentication token
- Handle polling timeout gracefully
- Provide clear visual feedback during polling

**Non-Goals:**
- WebSocket-based real-time updates (polling is simpler for MVP)
- Backend notification/webhook system
- Handling pairing rejection (assume approval or timeout)
- Persisting polling state across page refreshes
- Supporting multiple simultaneous pairing requests

## Decisions

### Decision 1: Poll with 1-second interval, 30-second timeout

**Rationale**: Pairing approval is likely a manual administrative action that takes seconds to minutes. 1-second polling provides responsive feedback without excessive server load. 30-second timeout prevents infinite polling if approval never comes.

**Alternatives considered**:
- Exponential backoff: Adds complexity without clear benefit for short-lived polling
- WebSocket: Over-engineering for this use case; polling is simpler
- 500ms interval: Too aggressive for a manual approval process

**Decision**: Poll every 1 second for up to 30 seconds. Use `setInterval` with cleanup on unmount.

### Decision 2: GET `/pairing-requests/:id` returns 202 while pending, 200 when approved

**Rationale**: HTTP status codes naturally express polling state. `202 Accepted` means "request received but not yet processed". `200 OK` with approval details means "pairing complete".

**Alternatives considered**:
- Always return 200 with status field in body: Less RESTful, requires parsing body to know state
- 404 for pending: Semantically incorrect (resource exists but is pending)
- 204 No Content for pending: Loses ability to include progress metadata

**Decision**: Return `202 No Content` while pending, `200 OK` with pairing details (device token, etc.) when approved. Include response body with `{"status":"pending"}` or `{"status":"approved","deviceToken":"..."}`.

### Decision 3: Navigate to root with token fragment, removing `/integration/device-pairing`

**Rationale**: The pairing UI is served at `/integration/device-pairing`, but OpenClaw main UI is at the root. Authentication token must be preserved in the URL fragment to maintain session.

**Implementation**:
1. Get current `window.location.href`
2. Parse and extract `#token=...` fragment
3. Construct new URL: `${protocol}//${host}${fragment}`
4. Use `window.location.href = newUrl` for navigation

**Alternatives considered**:
- Redirect to specific path like `/dashboard`: Assumes OpenClaw structure we shouldn't hard-code
- Use `window.location.replace()`: Prevents back button, which might confuse users
- Post message to parent frame: Only works if embedded; we're a standalone page

**Decision**: Navigate to `${protocol}//${host}#token=...`, letting OpenClaw handle routing from its root.

### Decision 4: Store polling state in React state, cleanup on unmount

**Rationale**: Polling is a component-level concern tied to the pairing request. React state naturally manages the lifecycle with `useEffect` cleanup.

**Implementation**:
- Start polling when `pairingRequestId` is set
- Store interval ID for cleanup
- Clear interval when component unmounts or polling completes
- Update button enable state based on approval status

**Alternatives considered**:
- Global polling manager: Over-engineering for single-use case
- Web Workers: Unnecessary for 1-second interval

**Decision**: Use `useEffect` with interval cleanup and button state derived from polling result.

## Risks / Trade-offs

**[Risk]** Polling continues even if user navigates away (tab closed, browser crash)
→ **Mitigation**: Acceptable for MVP. Polling stops after 30s timeout. Backend should handle abandoned requests gracefully.

**[Risk]** 30-second timeout might be too short for slow approval processes
→ **Mitigation**: Timeout is configurable in code. Can extend if needed based on user feedback. User can refresh page to retry.

**[Risk]** Network issues during polling cause silent failures
→ **Mitigation**: Show error message if polling request fails. Stop polling on network error and display retry option.

**[Risk]** Backend returns 200 before pairing is fully complete (race condition)
→ **Mitigation**: Backend must only return 200 when pairing is fully committed. This is a backend contract enforced by specs.

**[Trade-off]** Polling vs WebSocket adds latency and server load
→ **Acceptable**: 1-second latency is imperceptible for human-initiated approval. 30 requests per pairing is minimal load. Can upgrade to WebSocket later if needed.

**[Risk]** Token fragment might be missing or malformed
→ **Mitigation**: Check for token presence before navigation. Fall back to root without fragment if missing. Log warning for debugging.
