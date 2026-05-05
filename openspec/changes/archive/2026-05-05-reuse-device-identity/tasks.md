## 1. Remove Force Regeneration Flag

- [x] 1.1 Read current implementation in ui/src/device-identity.ts
- [x] 1.2 Locate the `forceRegenerate` flag declaration (around line 64-65)
- [x] 1.3 Remove the `const forceRegenerate = true;` line and its comment
- [x] 1.4 Remove the `if (!forceRegenerate)` conditional wrapper (line 67)
- [x] 1.5 Remove the closing brace of the conditional block (line 101)
- [x] 1.6 Verify the localStorage lookup code is now unconditional

## 2. Test Device Identity Persistence

- [x] 2.1 Clear localStorage in the browser
- [x] 2.2 Load the application and verify a new device identity is generated
- [x] 2.3 Check localStorage contains the device identity under key "openclaw-device-identity-v1"
- [x] 2.4 Reload the page and verify the same device ID is reused (check console logs)
- [x] 2.5 Verify no new device identity generation console message appears on reload

## 3. Test Identity Validation

- [x] 3.1 Manually corrupt the stored device identity in localStorage (e.g., change deviceId)
- [x] 3.2 Reload the page and verify deviceId is re-derived from publicKey
- [x] 3.3 Verify the corrected identity is saved back to localStorage
- [x] 3.4 Test with completely invalid JSON in localStorage and verify graceful regeneration

## 4. Verify WebSocket Handshake

- [x] 4.1 Load the application and complete the initial handshake
- [x] 4.2 Note the device ID from the handshake
- [x] 4.3 Reload the page and verify the same device ID is used in the new handshake
- [x] 4.4 Confirm persistent device identity across multiple page loads
