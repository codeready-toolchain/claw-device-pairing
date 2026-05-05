## Why

The device identity code currently has a `forceRegenerate = true` flag that bypasses localStorage lookup and always generates a new device identity, even when a valid one exists. This was added as a temporary workaround for `@noble/ed25519` v3.x compatibility but causes devices to get new IDs on every page load, breaking device persistence and session continuity.

## What Changes

- Remove the `forceRegenerate` flag from `loadOrCreateDeviceIdentity()` in `ui/src/device-identity.ts`
- Enable the localStorage lookup logic that was temporarily disabled
- Allow devices to reuse their stored identity across page loads and sessions

## Capabilities

### New Capabilities
<!-- None - this is removing a temporary workaround -->

### Modified Capabilities
- `device-identity`: Restore persistent device identity behavior by removing the force regeneration flag

## Impact

- `ui/src/device-identity.ts` - remove `forceRegenerate` flag and conditional block
- Device persistence - devices will now maintain the same device ID across page loads
- No breaking changes - this restores the originally intended behavior
- User experience - improves by maintaining stable device identity
