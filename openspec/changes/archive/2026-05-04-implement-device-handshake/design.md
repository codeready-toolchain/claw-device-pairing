## Context

The device pairing UI needs to authenticate with an OpenClaw gateway using the WebSocket protocol documented in `openclaw/docs/gateway/protocol.md`. The gateway requires all clients to provide device identity (Ed25519 keypair) and sign a challenge nonce during the handshake. The existing OpenClaw codebase (`openclaw/ui/src/ui/device-identity.ts`) provides a reference implementation that we must replicate exactly to ensure compatibility.

## Goals / Non-Goals

**Goals:**
- Implement Ed25519 device identity generation that produces identical output to OpenClaw's `generateIdentity()` function
- Implement WebSocket handshake protocol that handles `connect.challenge` and sends signed `connect` requests
- Store device identity persistently in browser localStorage using the same schema as OpenClaw
- Support the v3 signature payload format that binds nonce, device ID, platform, and other connection parameters

**Non-Goals:**
- Implementing full WebSocket client with all RPC methods (only handshake is in scope)
- Token refresh or rotation logic (out of scope for initial pairing)
- Supporting legacy v2 signature format
- Implementing all gateway protocol features beyond authentication

## Decisions

### Use @noble/ed25519 library for cryptographic operations
**Decision**: Use the `@noble/ed25519` library for Ed25519 keypair generation and signing.

**Rationale**: OpenClaw uses this exact library. Using the same library ensures binary-compatible signatures and eliminates compatibility issues.

**Alternatives considered**:
- Web Crypto API: Doesn't support Ed25519 directly in all browsers
- tweetnacl: Different library could produce incompatible key formats

### Replicate OpenClaw's device identity schema exactly
**Decision**: Use identical localStorage key (`openclaw-device-identity-v1`) and StoredIdentity schema with version, deviceId, publicKey, privateKey, and createdAtMs fields.

**Rationale**: If users have already paired an OpenClaw device from this browser, reusing the same identity prevents duplicate pairing requests. Exact schema compatibility is critical.

**Alternatives considered**:
- Custom schema: Would require new pairing approval for same browser/device
- Different storage key: Would duplicate device identities unnecessarily

### Implement v3 signature payload format
**Decision**: Sign the v3 payload format: `device=${deviceId}:client=${clientId}:role=${role}:scopes=${scopes}:token=${token}:nonce=${nonce}:platform=${platform}:deviceFamily=${deviceFamily}:ts=${signedAt}`

**Rationale**: The protocol docs specify v3 as the preferred format with stronger binding including platform and deviceFamily. v2 is accepted for compatibility but v3 is the target.

**Alternatives considered**:
- v2 format: Less secure, legacy compatibility only

### Use base64url encoding for keys and signatures
**Decision**: Encode all cryptographic material (public keys, private keys, signatures) using base64url (RFC 4648).

**Rationale**: OpenClaw uses base64url throughout. It's URL-safe and the protocol expects this encoding.

**Alternatives considered**:
- Hex encoding: Less compact, not used by OpenClaw
- Standard base64: Not URL-safe, requires escaping

### Device ID as SHA-256 fingerprint of public key
**Decision**: Derive deviceId by computing SHA-256 hash of the Ed25519 public key and encoding as lowercase hex.

**Rationale**: This is exactly how OpenClaw derives device IDs. The fingerprint provides stable, globally unique device identification.

**Alternatives considered**:
- Random UUID: Not deterministic from keypair, breaks recovery scenarios
- Hash of private key: Insecure

## Risks / Trade-offs

**Risk**: localStorage can be cleared by users → **Mitigation**: Document that clearing browser data requires re-pairing; this is expected behavior

**Risk**: Timing attacks on signature verification → **Mitigation**: Signing happens client-side, timing attacks are gateway concern

**Risk**: WebSocket connection failures during handshake → **Mitigation**: Implement timeout and error handling for challenge wait and connect response

**Trade-off**: Exact OpenClaw compatibility limits flexibility → Acceptable since compatibility is the primary goal

**Risk**: Platform/deviceFamily values need to match gateway expectations → **Mitigation**: Use standard platform values from navigator.platform or hardcoded sensible defaults
