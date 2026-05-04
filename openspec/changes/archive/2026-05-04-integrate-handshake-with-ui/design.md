## Context

The UI currently displays a ProgressStepper with two steps but has no logic to drive progression. The handshake module (ui/src/handshake.ts) implements the WebSocket authentication protocol but isn't connected to the React UI. We need to wire them together so the handshake runs on page load and updates the UI state.

## Goals / Non-Goals

**Goals:**
- Trigger handshake automatically when App component mounts using React useEffect
- Track handshake state (idle, loading, success, error) in React useState
- Update ProgressStepper to show first step as complete when handshake succeeds
- Display error messages to user if handshake fails
- Configure gateway URL from environment variable with sensible default

**Non-Goals:**
- Implementing retry logic for failed handshakes (future enhancement)
- Implementing the second pairing step logic (separate change)
- Persisting handshake result beyond page reload (localStorage not needed yet)
- Supporting multiple gateway URLs or connection switching

## Decisions

### Use React useEffect for handshake trigger
**Decision**: Call performHandshake in a useEffect hook with empty dependency array to trigger on mount.

**Rationale**: useEffect with empty deps runs once on mount, which is exactly when we want to initiate the handshake. This is idiomatic React.

**Alternatives considered**:
- Call in event handler: Wrong UX, handshake should be automatic
- Call directly in render: Violates React rules, causes infinite loops

### Track handshake state with useState
**Decision**: Use useState to track: handshakeStatus ('idle' | 'loading' | 'success' | 'error'), handshakeError (string | null), and deviceToken (string | null).

**Rationale**: React state triggers re-renders when handshake status changes, allowing UI to update. Separate fields for status, error, and token keep state clear.

**Alternatives considered**:
- Single state object: More complex updates, harder to read
- useReducer: Overkill for this simple state machine

### Mark step complete using Patternfly ProgressStep variant
**Decision**: Set ProgressStep variant to "success" for step 1 when handshakeStatus === 'success'.

**Rationale**: Patternfly ProgressStep supports variant prop for visual states. "success" shows a checkmark indicating completion.

**Alternatives considered**:
- Custom CSS styling: More work, inconsistent with Patternfly patterns
- Different component: ProgressStepper is already in use

### Gateway URL from environment variable
**Decision**: Use import.meta.env.VITE_GATEWAY_URL with fallback to ws://localhost:8080/ws.

**Rationale**: Vite exposes env vars prefixed with VITE_. This allows deployment-time configuration without code changes. Localhost default works for development.

**Alternatives considered**:
- Hardcoded URL: Not flexible for different environments
- Runtime configuration API: Over-engineered for simple URL config

### Display error in CardBody
**Decision**: Show error message in CardBody when handshakeStatus === 'error', replacing the ProgressStepper.

**Rationale**: Errors should be visible and actionable. Replacing stepper makes it clear something went wrong.

**Alternatives considered**:
- Toast notification: Might be dismissed before user sees it
- Alert component above stepper: Takes up space even when no error

## Risks / Trade-offs

**Risk**: Handshake fails silently if gateway unreachable → **Mitigation**: Display clear error message with error code

**Risk**: User refreshes page, handshake runs again → **Mitigation**: Acceptable for now, future work can add token persistence

**Risk**: Slow handshake causes long loading state → **Mitigation**: Timeouts in handshake.ts (10s challenge, 30s response) prevent indefinite waiting

**Trade-off**: Hardcoding role='operator' and scopes=[] → Acceptable since this is specifically a device pairing flow for operators

**Risk**: Environment variable misconfiguration in production → **Mitigation**: Document required VITE_GATEWAY_URL in README or deployment guide
