## Why

The device pairing UI needs to authenticate with the OpenClaw gateway using the WebSocket handshake protocol. This requires implementing device identity generation (Ed25519 keypairs) and the challenge-response authentication flow described in the gateway protocol specification.

## What Changes

- Implement device identity generation using Ed25519 keypairs with the exact same logic as OpenClaw's `generateIdentity()` function
- Implement WebSocket handshake protocol that handles `connect.challenge` events and responds with signed `connect` requests
- Add device authentication signing logic that creates signatures binding the server nonce, device ID, and other connection parameters
- Store and retrieve device identity from browser localStorage
- Add utility functions for base64url encoding/decoding and cryptographic operations

## Capabilities

### New Capabilities
- `device-identity`: Ed25519 keypair generation, device ID fingerprinting, and persistent storage using the exact implementation from OpenClaw
- `handshake-protocol`: WebSocket challenge-response authentication flow including connect.challenge handling and signed connect requests

### Modified Capabilities

## Impact

- New TypeScript modules in `ui/src/` for device identity and handshake protocol
- New dependency: `@noble/ed25519` for Ed25519 cryptographic operations
- Browser localStorage will store device identity keypairs
- WebSocket connection logic will be added to handle the gateway handshake sequence
