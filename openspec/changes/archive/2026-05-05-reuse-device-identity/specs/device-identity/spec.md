## MODIFIED Requirements

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
