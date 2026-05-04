## 1. Dependencies and Setup

- [x] 1.1 Add @noble/ed25519 dependency to ui/package.json
- [x] 1.2 Create ui/src/device-identity.ts module file
- [x] 1.3 Create ui/src/handshake.ts module file

## 2. Base64url Encoding Utilities

- [x] 2.1 Implement base64UrlEncode() function to encode Uint8Array to base64url without padding
- [x] 2.2 Implement base64UrlDecode() function to decode base64url strings back to Uint8Array
- [x] 2.3 Implement bytesToHex() function to convert Uint8Array to lowercase hex string

## 3. Device Identity Core Functions

- [x] 3.1 Implement fingerprintPublicKey() to compute SHA-256 hash of public key as hex
- [x] 3.2 Implement generateIdentity() to create Ed25519 keypair and derive deviceId
- [x] 3.3 Define DeviceIdentity and StoredIdentity TypeScript types

## 4. Device Identity Storage

- [x] 4.1 Implement loadOrCreateDeviceIdentity() to load from localStorage or generate new identity
- [x] 4.2 Add validation logic to verify stored deviceId matches public key fingerprint
- [x] 4.3 Add error handling for corrupted localStorage data with fallback to regeneration

## 5. Device Payload Signing

- [x] 5.1 Implement signDevicePayload() to sign UTF-8 encoded strings with Ed25519 private key
- [x] 5.2 Ensure signature output is base64url encoded

## 6. Signature Payload Construction

- [x] 6.1 Implement buildSignaturePayload() to construct v3 format payload string
- [x] 6.2 Include all required fields: device, client, role, scopes, token, nonce, platform, deviceFamily, ts
- [x] 6.3 Implement scopes array serialization as comma-delimited string

## 7. Connect Request Construction

- [x] 7.1 Implement buildConnectRequest() to construct the connect request frame
- [x] 7.2 Set protocol version to minProtocol:3, maxProtocol:3
- [x] 7.3 Populate client object with id, version, platform, mode
- [x] 7.4 Populate role and scopes fields
- [x] 7.5 Populate device object with id, publicKey, signature, signedAt, nonce

## 8. Challenge Handling

- [x] 8.1 Implement waitForChallenge() to receive and parse connect.challenge event
- [x] 8.2 Extract nonce and ts from challenge payload
- [x] 8.3 Add 10-second timeout for challenge reception

## 9. Connect Response Handling

- [x] 9.1 Implement waitForConnectResponse() to receive connect response
- [x] 9.2 Parse hello-ok payload to extract protocol, server info, and device token
- [x] 9.3 Parse error responses to extract error codes and messages
- [x] 9.4 Add 30-second timeout for connect response

## 10. Handshake Orchestration

- [x] 10.1 Implement performHandshake() function to orchestrate the full handshake flow
- [x] 10.2 Sequence: establish WebSocket, wait for challenge, sign payload, send connect, wait for response
- [x] 10.3 Return device token and connection info on success
- [x] 10.4 Throw descriptive errors on failure with error codes

## 11. Integration and Verification

- [x] 11.1 Test device identity generation creates valid Ed25519 keypairs
- [x] 11.2 Verify deviceId matches SHA-256 fingerprint of public key
- [x] 11.3 Test localStorage persistence and retrieval
- [x] 11.4 Verify signature payload matches v3 format specification
- [x] 11.5 Test complete handshake flow end-to-end (may require mock gateway or integration test environment)
