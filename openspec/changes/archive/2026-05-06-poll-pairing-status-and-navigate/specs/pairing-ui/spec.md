## ADDED Requirements

### Requirement: Navigation Button Component
The card body SHALL include a "Go to OpenClaw" button below the progress stepper.

#### Scenario: Button is rendered
- **WHEN** the application renders
- **THEN** a Button component with text "Go to OpenClaw" is displayed below the ProgressStepper

#### Scenario: Button uses Patternfly component
- **WHEN** the navigation button is rendered
- **THEN** it uses the Patternfly Button component from `@patternfly/react-core`

#### Scenario: Button disabled state is controlled
- **WHEN** the button is rendered
- **THEN** its disabled state is controlled by React state based on pairing approval status
