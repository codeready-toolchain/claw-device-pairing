## Context

The current device pairing UI (ui/src/App.jsx) displays a simple Patternfly Card with "Pairing device..." text. The user flow involves two sequential steps (generate device ID, then pair with OpenClaw), but this progression is invisible to the user. The project already uses Patternfly React Core v6.4.1 which includes a Stepper component designed for multi-step workflows.

## Goals / Non-Goals

**Goals:**
- Display device pairing progress using Patternfly's Stepper component
- Show two sequential steps: "Generate device id" and "Pair device with OpenClaw"
- Integrate stepper into the existing Card layout without breaking current functionality

**Non-Goals:**
- Implementing actual device ID generation or pairing logic (backend work)
- Adding step validation or error handling (future enhancement)
- Making steps interactive or allowing users to jump between steps

## Decisions

### Use Patternfly Stepper component
**Decision**: Use Patternfly's Stepper component with vertical orientation inside the Card body.

**Rationale**: The Stepper component is purpose-built for showing progress through sequential steps. Vertical orientation works well within a card and is mobile-friendly. Alternative considered: horizontal stepper, but vertical is more compact for a 2-step flow.

**Alternatives considered**:
- Custom progress indicator: More work, less consistent with Patternfly patterns
- ProgressBar: Doesn't convey individual step labels as clearly

### State management for current step
**Decision**: Use React useState hook to track current step index (0-based).

**Rationale**: Simple local state is sufficient for this UI-only change. No need for complex state management (Redux, Context) when state isn't shared across components.

**Alternatives considered**:
- Derive step from props/API: Premature - backend doesn't track step state yet

### Step progression approach
**Decision**: Steps remain static in this change. No automatic progression or step completion logic.

**Rationale**: This change adds visual structure only. Actual progression logic requires backend integration (future work).

## Risks / Trade-offs

**Risk**: Stepper shows steps but doesn't reflect actual pairing state → **Mitigation**: Document this as a visual-only change in tasks; backend integration is a separate concern

**Trade-off**: Vertical stepper takes more vertical space than the simple text → Acceptable since it provides better UX and the card can expand
