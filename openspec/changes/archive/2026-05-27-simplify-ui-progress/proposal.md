## Why

The current ProgressStepper UI is visually heavy for a simple two-step flow. A lightweight Spinner with a status label provides a cleaner, more focused experience. Additionally, requiring users to click "Go to OpenClaw" after pairing adds unnecessary friction -- the app should auto-redirect once pairing succeeds.

## What Changes

- Replace the PatternFly `ProgressStepper` component with a `Spinner` and a text label showing the current step
- Remove the "Go to OpenClaw" `Button` -- navigate automatically to the OpenClaw URL when pairing is approved
- Remove unused PatternFly imports (`ProgressStepper`, `ProgressStep`, `Button`)

## Capabilities

### New Capabilities

- `spinner-status`: Replace the ProgressStepper with a Spinner and a status label that reflects the current pairing step

### Modified Capabilities

- `navigation-to-openclaw`: Navigation becomes automatic on pairing approval instead of requiring a button click
- `pairing-ui`: Card body displays a Spinner + label instead of a ProgressStepper, and the "Go to OpenClaw" button is removed
- `progress-stepper`: This capability is superseded -- the ProgressStepper component is removed entirely

## Impact

- `ui/src/App.jsx`: Main component rewritten to use Spinner + label, auto-redirect logic replaces button click handler
- PatternFly imports: `ProgressStepper`, `ProgressStep`, `Button` removed; `Spinner` added
- No backend changes required
- No new dependencies (Spinner is already in `@patternfly/react-core`)
