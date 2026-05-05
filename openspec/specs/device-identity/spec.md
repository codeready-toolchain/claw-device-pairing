# device-identity Specification

## Purpose
TBD - created by archiving change implement-device-handshake. Update Purpose after archive.
## Requirements
### Requirement: Ed25519 Keypair Generation
The system SHALL generate Ed25519 keypairs using the @noble/ed25519 library with cryptographically secure random private keys.

#### Scenario: Generate new keypair
- **WHEN** generateIdentity() is called
- **THEN** a new Ed25519 private key is generated using utils.randomSecretKey()

#### Scenario: Derive public key
- **WHEN** a private key is generated
- **THEN** the corresponding public key is derived using getPublicKeyAsync()

### Requirement: Device ID Fingerprinting
The device ID SHALL be the SHA-256 hash of the Ed25519 public key, encoded as lowercase hexadecimal.

#### Scenario: Fingerprint public key
- **WHEN** fingerprintPublicKey() is called with a public key
- **THEN** it returns the SHA-256 hash of the public key as lowercase hex string

#### Scenario: Device ID matches fingerprint
- **WHEN** a device identity is generated
- **THEN** the deviceId field equals the SHA-256 fingerprint of the publicKey

### Requirement: Base64url Encoding
Cryptographic material SHALL be encoded using base64url format (RFC 4648) without padding.

#### Scenario: Encode public key
- **WHEN** a public key is generated
- **THEN** it is encoded as base64url with '+' replaced by '-', '/' replaced by '_', and trailing '=' removed

#### Scenario: Encode private key
- **WHEN** a private key is generated
- **THEN** it is encoded as base64url with the same substitutions

#### Scenario: Decode base64url
- **WHEN** base64UrlDecode() is called with encoded data
- **THEN** it reverses the encoding and returns the original Uint8Array

### Requirement: Persistent Storage
Device identity SHALL be stored in browser localStorage using the key "openclaw-device-identity-v1" with a versioned schema and loaded on subsequent page loads.

#### Scenario: Store identity on generation
- **WHEN** a new identity is generated
- **THEN** it is stored in localStorage with version:1, deviceId, publicKey, privateKey, and createdAtMs fields

#### Scenario: Load existing identity
- **WHEN** loadOrCreateDeviceIdentity() is called and valid stored identity exists
- **THEN** the stored identity is loaded and returned

#### Scenario: Validate stored deviceId
- **WHEN** loading a stored identity
- **THEN** the deviceId is re-derived from the publicKey and corrected if mismatched

#### Scenario: Regenerate on corruption
- **WHEN** stored identity is invalid or corrupted
- **THEN** a new identity is generated and stored

#### Scenario: No forced regeneration
- **WHEN** loadOrCreateDeviceIdentity() is called
- **THEN** the function checks localStorage first and only generates a new identity if no valid stored identity exists

### Requirement: Device Payload Signing
The system SHALL sign device authentication payloads using Ed25519 signature with the device's private key.

#### Scenario: Sign payload string
- **WHEN** signDevicePayload() is called with a privateKey and payload string
- **THEN** the payload is UTF-8 encoded, signed with Ed25519, and returned as base64url

#### Scenario: Signature is deterministic
- **WHEN** the same payload and private key are signed multiple times
- **THEN** the same signature is produced each time

### Requirement: Type Safety
Device identity SHALL be represented by strongly-typed TypeScript interfaces.

#### Scenario: DeviceIdentity type
- **WHEN** a DeviceIdentity is created
- **THEN** it has deviceId, publicKey, and privateKey string fields

#### Scenario: StoredIdentity type
- **WHEN** a StoredIdentity is stored
- **THEN** it has version, deviceId, publicKey, privateKey, and createdAtMs fields

