## MODIFIED Requirements

### Requirement: Pairing Status Display
The card body SHALL display a progress stepper showing the stages of device pairing.

#### Scenario: Stepper replaces simple message
- **WHEN** the application renders
- **THEN** the card body displays a Stepper component instead of just "Pairing device..." text

#### Scenario: Loading spinner is removed
- **WHEN** the card body is displayed
- **THEN** no Spinner component is shown (replaced by the Stepper)
