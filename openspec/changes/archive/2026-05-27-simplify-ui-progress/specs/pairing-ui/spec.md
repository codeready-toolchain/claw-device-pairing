## MODIFIED Requirements

### Requirement: Pairing Status Display
The card body SHALL display a Spinner with a status label showing the current stage of device pairing.

#### Scenario: Spinner replaces stepper
- **WHEN** the application renders
- **THEN** the card body displays a Spinner component with a status label instead of a ProgressStepper

#### Scenario: No navigation button
- **WHEN** the card body is displayed
- **THEN** no "Go to OpenClaw" button is shown

### Requirement: Navigation Button Component
**This requirement is removed -- see navigation-to-openclaw delta spec.**

#### Scenario: Button is no longer rendered
- **WHEN** the application renders
- **THEN** no Button component is displayed in the card body
