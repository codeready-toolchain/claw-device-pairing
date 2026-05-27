## Context

The device pairing UI currently uses a PatternFly `ProgressStepper` with two steps ("Generate device id" and "Pair device with OpenClaw") and a "Go to OpenClaw" button that the user must click after pairing succeeds. The flow is straightforward and the stepper adds visual weight without much benefit. After pairing approval, requiring a manual click delays the user unnecessarily.

## Goals / Non-Goals

**Goals:**
- Replace the ProgressStepper with a centered Spinner and a text label that reflects the current step
- Automatically redirect to the OpenClaw URL once pairing is approved, eliminating the manual button click
- Keep the same state machine and handshake/pairing logic unchanged

**Non-Goals:**
- Changing the handshake or pairing protocol
- Modifying backend endpoints or pairing approval flow
- Adding error recovery or retry mechanisms (existing error handling remains as-is)

## Decisions

### Replace ProgressStepper with Spinner + label

**Decision**: Use a single PatternFly `Spinner` component with a text element below it showing the current step label.

**Rationale**: The two-step flow is too simple to warrant a stepper. A spinner with descriptive text communicates progress just as effectively with less visual clutter.

**Alternatives considered**:
- Keep ProgressStepper but simplify styling -- still heavier than needed for two steps
- Use a custom loading animation -- unnecessary when PatternFly provides `Spinner`

**Step labels by state**:
- During handshake (loading): "Generating device id..."
- Handshake done, pairing in progress (pending/progressing): "Pairing device with OpenClaw..."
- Error: Show error message text, hide spinner
- Success: "Redirecting to OpenClaw..." (briefly visible before redirect)

### Auto-redirect on approval

**Decision**: Use a `useEffect` that watches `approvalStatus`. When it becomes `'approved'`, call `navigateToOpenClaw()` automatically.

**Rationale**: Eliminates a click that adds no value. The user's intent after successful pairing is always to proceed to OpenClaw.

**Alternatives considered**:
- Redirect immediately in the polling callback -- less clean, mixes concerns
- Add a short delay before redirect -- unnecessary complexity, the spinner label provides transition context

## Risks / Trade-offs

- [No cancel opportunity] Users cannot abort the redirect once pairing succeeds. This is acceptable because proceeding to OpenClaw is the only valid next action.
- [Brief flash of redirect label] The "Redirecting to OpenClaw..." label may appear only briefly before navigation. This is fine -- it confirms the transition is happening.
