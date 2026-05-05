## MODIFIED Requirements

### Requirement: Connection Parameters
The handshake SHALL use appropriate connection parameters for device pairing, including optional authentication token from URL fragment.

#### Scenario: Client ID is set
- **WHEN** performHandshake is called
- **THEN** clientId is set to "device-pairing-ui"

#### Scenario: Role is operator
- **WHEN** performHandshake is called
- **THEN** role is set to "operator"

#### Scenario: Scopes are provided
- **WHEN** performHandshake is called
- **THEN** scopes array is set appropriately (may be empty for pairing flow)

#### Scenario: Client version is set
- **WHEN** performHandshake is called
- **THEN** clientVersion is set to a valid version string

#### Scenario: Token extracted from URL fragment without prefix
- **WHEN** the URL fragment is "#token=abc123"
- **THEN** the token value "abc123" is passed to performHandshake (without the "token=" prefix)

#### Scenario: Token extracted from URL fragment with hash only
- **WHEN** the URL fragment is "#abc123"
- **THEN** the token value "abc123" is passed to performHandshake

#### Scenario: Token extracted from URL fragment with slash
- **WHEN** the URL fragment is "#/token=abc123"
- **THEN** the token value "abc123" is passed to performHandshake (without the "token=" prefix)

#### Scenario: No token when fragment is empty
- **WHEN** the URL has no fragment
- **THEN** performHandshake is called with undefined token value
