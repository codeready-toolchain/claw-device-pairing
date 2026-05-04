## 1. Update Component Imports

- [x] 1.1 Import Stepper and Step components from @patternfly/react-core in App.jsx
- [x] 1.2 Import useState hook from React

## 2. Add State Management

- [x] 2.1 Add useState hook to track current step index (initialize to 0)

## 3. Update Card Body Layout

- [x] 3.1 Remove the simple "Pairing device..." text from CardBody
- [x] 3.2 Add Stepper component to CardBody with isVertical={true}

## 4. Define Stepper Steps

- [x] 4.1 Create Step component for "Generate device id" with id="step-1" and titleId="step-1-title"
- [x] 4.2 Create Step component for "Pair device with OpenClaw" with id="step-2" and titleId="step-2-title"
- [x] 4.3 Ensure steps are ordered correctly within the Stepper component

## 5. Verification

- [x] 5.1 Verify the UI renders with the Stepper visible in the Card
- [x] 5.2 Verify both steps display with correct titles in the correct order
