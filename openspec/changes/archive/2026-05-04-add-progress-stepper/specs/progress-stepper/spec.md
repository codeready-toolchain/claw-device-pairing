## ADDED Requirements

### Requirement: Stepper Component Display
The card body SHALL display a Patternfly Stepper component showing the device pairing workflow steps.

#### Scenario: Stepper is rendered in card body
- **WHEN** the application loads
- **THEN** a Patternfly Stepper component is visible inside the Card body

#### Scenario: Stepper uses vertical orientation
- **WHEN** the Stepper is displayed
- **THEN** the Stepper component has `isVertical={true}` property

### Requirement: Step Definitions
The Stepper SHALL display exactly two steps in sequential order representing the device pairing flow.

#### Scenario: First step shows generate device ID
- **WHEN** the Stepper is rendered
- **THEN** the first step displays title "Generate device id"

#### Scenario: Second step shows pair with OpenClaw
- **WHEN** the Stepper is rendered
- **THEN** the second step displays title "Pair device with OpenClaw"

#### Scenario: Steps are ordered correctly
- **WHEN** viewing the Stepper
- **THEN** "Generate device id" appears before "Pair device with OpenClaw"

### Requirement: Step State Tracking
The application SHALL track the current step using React state.

#### Scenario: Current step is stored in state
- **WHEN** the component initializes
- **THEN** a state variable tracks the current step index (0-based)

#### Scenario: Initial step is first step
- **WHEN** the application first loads
- **THEN** the current step index is 0 (Generate device id)

### Requirement: Patternfly Step Components
Each step SHALL be rendered using Patternfly's Step component with appropriate properties.

#### Scenario: Steps use Patternfly Step component
- **WHEN** the Stepper is rendered
- **THEN** each step uses the Patternfly `<Step>` component

#### Scenario: Steps have unique IDs
- **WHEN** steps are defined
- **THEN** each step has a unique `id` property

#### Scenario: Step titles are set
- **WHEN** each Step component is rendered
- **THEN** it has a `titleId` property matching the step's unique identifier
