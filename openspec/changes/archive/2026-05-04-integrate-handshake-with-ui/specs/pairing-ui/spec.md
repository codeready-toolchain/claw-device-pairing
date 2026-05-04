## ADDED Requirements

### Requirement: Dynamic Content Based on Handshake State
The card body SHALL display different content based on the current handshake state.

#### Scenario: Loading state content
- **WHEN** the handshake is in progress (loading state)
- **THEN** a loading message or spinner is displayed in the card body

#### Scenario: Success state content
- **WHEN** the handshake succeeds
- **THEN** the ProgressStepper is displayed with the first step marked as complete

#### Scenario: Error state content
- **WHEN** the handshake fails
- **THEN** an error message is displayed instead of the ProgressStepper

### Requirement: State-Driven Rendering
The application SHALL use React state to control UI rendering based on handshake lifecycle.

#### Scenario: UseEffect for handshake
- **WHEN** the App component mounts
- **THEN** a useEffect hook triggers the handshake

#### Scenario: State updates trigger re-renders
- **WHEN** handshake state changes
- **THEN** the UI re-renders to reflect the new state
