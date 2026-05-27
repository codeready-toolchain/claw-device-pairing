## 1. Replace ProgressStepper with Spinner + Status Label

- [x] 1.1 Update imports in `ui/src/App.jsx`: remove `ProgressStepper`, `ProgressStep`, `Button`; add `Spinner`
- [x] 1.2 Replace the `<ProgressStepper>` block and `<Button>` with a `<Spinner>` and a status label element that derives its text from `handshakeStatus`, `pairingStatus`, and `approvalStatus` state
- [x] 1.3 Implement status label logic: "Generating device id..." during handshake, "Pairing device with OpenClaw..." during pairing, "Redirecting to OpenClaw..." on approval, error message on error

## 2. Auto-Redirect on Pairing Approval

- [x] 2.1 Add a `useEffect` that watches `approvalStatus` and calls `navigateToOpenClaw()` when it becomes `'approved'`
- [x] 2.2 Remove the `<Button>` onClick handler (now unused as a UI element) but keep the `navigateToOpenClaw` function for the auto-redirect

## 3. Verify

- [x] 3.1 Run `npm run build` in `ui/` to confirm the production build succeeds with no errors
