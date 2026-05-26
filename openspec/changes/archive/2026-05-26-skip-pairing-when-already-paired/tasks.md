## 1. Handle Already-Paired Device in Handshake Success Path

- [x] 1.1 In `ui/src/App.jsx`, update the handshake success path (after `setHandshakeStatus('success')` in the `try` block) to also call `setPairingStatus('success')` and `setApprovalStatus('approved')`
- [x] 1.2 Verify that the pairing submission `useEffect` does not trigger when `pairingRequestId` is `null` (existing guard is sufficient — no code change needed, just verify)

## 2. Testing

- [x] 2.1 Manually test the already-paired path: confirm both progress steps show success and the "Go to OpenClaw" button is enabled immediately
- [x] 2.2 Manually test the not-paired path: confirm the existing pairing submission and polling flow still works as before
