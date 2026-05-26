## Context

The `performHandshake` function in `handshake.ts` resolves successfully when the server responds with `ok: true`, which means the device is already paired. Currently, `App.jsx` only handles this as a basic handshake success — it sets `handshakeStatus` to `'success'` but leaves `pairingStatus` at `'idle'` and `approvalStatus` at `'idle'`. This means the "Pair device" step stays in a neutral state and the "Go to OpenClaw" button remains disabled, even though no further action is needed.

The handshake error path (`NOT_PAIRED`) correctly triggers the pairing submission and polling flow, but the success path is incomplete.

## Goals / Non-Goals

**Goals:**
- When `performHandshake` resolves (device already paired), immediately mark both progress steps as complete and enable the navigation button.

**Non-Goals:**
- Changing the `performHandshake` function or the server-side protocol.
- Modifying the pairing submission or polling flows — those remain unchanged for the `NOT_PAIRED` case.

## Decisions

**Decision: Update state inline in the existing handshake success path**

In the first `useEffect` of `App.jsx`, the `try` block already calls `setHandshakeStatus('success')` when the handshake resolves. We will add `setPairingStatus('success')` and `setApprovalStatus('approved')` at the same point. These three state updates will batch in the same React render cycle, so the UI will jump directly to the fully-completed state.

Alternative considered: introducing an `alreadyPaired` boolean state and conditioning the second `useEffect` on it. This was rejected because it adds unnecessary state — the existing `pairingStatus` and `approvalStatus` are sufficient to drive the UI.

## Risks / Trade-offs

- **[Risk] State coupling** — Setting `pairingStatus` to `'success'` in the handshake path means the pairing submission `useEffect` could theoretically fire if `pairingRequestId` were somehow set in the same render. → Mitigated: the success path never sets `pairingRequestId`, so the submission effect's guard (`if (!pairingRequestId) return`) prevents any side effects.
