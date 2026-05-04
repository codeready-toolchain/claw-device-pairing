## Why

The current device pairing UI shows a simple "Pairing device..." message without any visibility into the pairing process. Users need to see progress through different stages of device pairing to understand what's happening and where they are in the flow.

## What Changes

- Add a Patternfly Stepper component to the device pairing card that displays two sequential steps:
  - Generate device id
  - Pair device with OpenClaw
- Update the UI to visually track progress through the pairing stages

## Capabilities

### New Capabilities
- `progress-stepper`: UI component for showing multi-step progress in the device pairing flow using Patternfly Stepper

### Modified Capabilities
- `pairing-ui`: Enhance the existing UI to include visual step tracking and progress indication

## Impact

- Modifies `ui/src/App.jsx` to include Patternfly Stepper component
- Updates the device pairing card to show progress through pairing stages
- No API or backend changes required
- Relies on existing Patternfly React Core library (already a dependency)
