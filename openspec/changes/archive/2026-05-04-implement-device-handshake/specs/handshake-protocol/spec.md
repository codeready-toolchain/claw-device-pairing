## ADDED Requirements

### Requirement: WebSocket Connection
The system SHALL establish a WebSocket connection to the OpenClaw gateway using text frames with JSON payloads.

#### Scenario: Connect to gateway
- **WHEN** initiating a connection
- **THEN** a WebSocket is opened to the gateway URL using the ws or wss protocol

#### Scenario: JSON message format
- **WHEN** sending or receiving messages
- **THEN** all messages are JSON-encoded text frames

### Requirement: Challenge Reception
The system SHALL receive and parse the connect.challenge event from the gateway before sending the connect request.

#### Scenario: Receive challenge event
- **WHEN** WebSocket connection is established
- **THEN** the first message from the gateway is a connect.challenge event

#### Scenario: Parse challenge payload
- **WHEN** connect.challenge is received
- **THEN** it contains type:"event", event:"connect.challenge", and payload with nonce and ts fields

#### Scenario: Extract challenge nonce
- **WHEN** processing the challenge
- **THEN** the nonce value is extracted for use in signature generation

### Requirement: Connect Request Construction
The system SHALL construct a connect request with all required protocol fields including client identity, role, scopes, and device authentication.

#### Scenario: Build connect frame
- **WHEN** constructing the connect request
- **THEN** it has type:"req", a unique id, method:"connect", and params object

#### Scenario: Include protocol version
- **WHEN** params are populated
- **THEN** minProtocol and maxProtocol are both set to 3

#### Scenario: Include client info
- **WHEN** params are populated
- **THEN** client object includes id, version, platform, and mode fields

#### Scenario: Include role and scopes
- **WHEN** params are populated
- **THEN** role is set appropriately (e.g., "operator" or "node") and scopes array is provided

#### Scenario: Include device identity
- **WHEN** params are populated
- **THEN** device object includes id, publicKey, signature, signedAt, and nonce fields

### Requirement: Signature Payload Construction
The system SHALL construct the v3 signature payload binding all critical authentication parameters.

#### Scenario: Build v3 payload
- **WHEN** creating the signature payload
- **THEN** it follows the format: device=${deviceId}:client=${clientId}:role=${role}:scopes=${scopes}:token=${token}:nonce=${nonce}:platform=${platform}:deviceFamily=${deviceFamily}:ts=${signedAt}

#### Scenario: Include challenge nonce
- **WHEN** building the signature payload
- **THEN** the nonce from connect.challenge is included in the payload

#### Scenario: Include timestamp
- **WHEN** building the signature payload
- **THEN** signedAt is set to the current Unix timestamp in milliseconds

#### Scenario: Scopes serialization
- **WHEN** scopes array is included in signature
- **THEN** scopes are joined with comma delimiter (e.g., "operator.read,operator.write")

### Requirement: Device Authentication Signing
The system SHALL sign the signature payload with the device's Ed25519 private key and include the signature in the connect request.

#### Scenario: Sign authentication payload
- **WHEN** the signature payload is constructed
- **THEN** it is signed using signDevicePayload() with the device private key

#### Scenario: Include signature in device object
- **WHEN** the connect request is built
- **THEN** device.signature contains the base64url-encoded signature

#### Scenario: Include matching nonce
- **WHEN** the connect request is built
- **THEN** device.nonce matches the nonce from connect.challenge

### Requirement: Connect Response Handling
The system SHALL receive and parse the connect response to determine authentication success or failure.

#### Scenario: Receive hello-ok response
- **WHEN** connect succeeds
- **THEN** the response has type:"res", ok:true, and payload.type:"hello-ok"

#### Scenario: Extract protocol version
- **WHEN** hello-ok is received
- **THEN** payload.protocol indicates the negotiated protocol version

#### Scenario: Extract server info
- **WHEN** hello-ok is received
- **THEN** payload.server contains version and connId fields

#### Scenario: Extract device token
- **WHEN** hello-ok includes auth
- **THEN** payload.auth.deviceToken contains the issued token for future connections

#### Scenario: Handle connect failure
- **WHEN** connect fails
- **THEN** the response has ok:false and error object with code and message

### Requirement: Handshake Timeout
The system SHALL enforce timeouts on handshake steps to prevent indefinite waiting.

#### Scenario: Challenge timeout
- **WHEN** waiting for connect.challenge
- **THEN** a timeout (10 seconds) is enforced and connection fails if exceeded

#### Scenario: Response timeout
- **WHEN** waiting for connect response
- **THEN** a timeout (30 seconds) is enforced and connection fails if exceeded

### Requirement: Error Handling
The system SHALL handle and report authentication errors with appropriate error codes and messages.

#### Scenario: Device auth errors
- **WHEN** device authentication fails
- **THEN** error.details.code indicates the specific failure (e.g., DEVICE_AUTH_NONCE_MISMATCH, DEVICE_AUTH_SIGNATURE_INVALID)

#### Scenario: Token auth errors
- **WHEN** token authentication fails
- **THEN** error includes appropriate AUTH_TOKEN_* error codes

#### Scenario: Protocol version mismatch
- **WHEN** protocol versions don't overlap
- **THEN** error indicates protocol version incompatibility
