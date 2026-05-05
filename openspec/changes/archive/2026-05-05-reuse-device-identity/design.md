## Context

The `loadOrCreateDeviceIdentity()` function in `ui/src/device-identity.ts` was designed to check localStorage for an existing device identity before generating a new one. However, a temporary workaround was added with a `forceRegenerate = true` flag to bypass this logic during the migration to `@noble/ed25519` v3.x.

Current implementation (lines 64-65):
```typescript
const forceRegenerate = true;  // Force regeneration for @noble/ed25519 v3.x compatibility

if (!forceRegenerate) {
  // localStorage lookup code that never executes
}
```

This causes a new device identity to be generated on every page load, defeating the purpose of persistent storage and causing devices to lose their identity.

## Goals / Non-Goals

**Goals:**
- Restore persistent device identity behavior
- Remove the temporary `forceRegenerate` workaround
- Enable localStorage lookup to reuse existing device identities
- Maintain stable device IDs across page loads and sessions

**Non-Goals:**
- Changing the localStorage schema or storage key
- Modifying the device identity generation logic
- Adding new identity validation rules beyond what's already implemented
- Supporting multiple device identities per browser

## Decisions

### Decision 1: Simply remove the forceRegenerate flag and conditional

**Rationale**: The localStorage lookup logic is already implemented and tested. The `forceRegenerate` flag was explicitly marked as temporary with a comment "Remove this after confirming it works". Since the `@noble/ed25519` v3.x migration is complete and working, we can simply remove the workaround.

**Alternatives considered**:
- Setting `forceRegenerate = false` instead of removing it: Leaves dead code in the codebase
- Adding a configuration option to control regeneration: Unnecessary complexity for what was meant to be a temporary workaround

**Decision**: Remove lines 64-66 entirely (the `forceRegenerate` flag declaration and the `if (!forceRegenerate)` wrapper around the localStorage logic).

### Decision 2: Keep existing validation and fallback logic

**Rationale**: The current code already handles:
- Missing or corrupt localStorage data (falls through to regeneration)
- Schema version validation
- DeviceId fingerprint validation and correction
- Safe error handling with try/catch

No changes needed to this logic - it's already robust.

## Risks / Trade-offs

**[Risk]** Existing stored identities might have been generated with an incompatible version of the crypto library
→ **Mitigation**: The validation logic already re-derives the deviceId from the public key and corrects mismatches. Invalid identities will trigger regeneration via the catch block.

**[Risk]** Users might have multiple device identities stored from when regeneration was forced
→ **Mitigation**: Only the most recent identity (stored under the v1 key) will be loaded. Previous identities are effectively orphaned but harmless.

**[Trade-off]** Removes the ability to force regeneration for debugging
→ **Acceptable**: Developers can manually clear localStorage if needed. The force regeneration was meant as a temporary migration tool, not a permanent debugging feature.
