## ADDED Requirements

### Requirement: OAuth Proxy Sidecar Container
The Deployment SHALL include an `oauth-proxy` sidecar container that authenticates and authorizes all incoming requests before forwarding them to the application container.

#### Scenario: Sidecar container is defined
- **WHEN** the Deployment manifest is examined
- **THEN** a second container named `oauth-proxy` is defined using the `registry.redhat.io/openshift4/ose-oauth-proxy` image

#### Scenario: Proxy listens on HTTPS port 8443
- **WHEN** the oauth-proxy container starts
- **THEN** it listens on port 8443 with TLS enabled using the service-serving-cert

#### Scenario: Proxy forwards to application on localhost
- **WHEN** an authenticated request is received
- **THEN** the oauth-proxy forwards it to `http://localhost:8080`

### Requirement: SAR Authorization Check
The oauth-proxy SHALL perform a SubjectAccessReview check to verify the user can "get pods" in the deployment's namespace.

#### Scenario: SAR check configuration
- **WHEN** the oauth-proxy container args are examined
- **THEN** the `--openshift-sar` flag is set to check the "get" verb on "pods" resource in the pod's namespace

#### Scenario: Authorized user can access the application
- **WHEN** a user who can "get pods" in the namespace accesses the Route
- **THEN** they are authenticated via OAuth and forwarded to the application

#### Scenario: Unauthorized user is denied
- **WHEN** a user who cannot "get pods" in the namespace attempts to access the Route
- **THEN** they receive a 403 Forbidden response after OAuth authentication

### Requirement: Exec-Based Health Probes
The `server` container SHALL use exec-based liveness and readiness probes that call the `/health` endpoint locally, bypassing the oauth-proxy entirely.

#### Scenario: Liveness probe uses exec command
- **WHEN** the `server` container definition is examined
- **THEN** a `livenessProbe` is configured with an `exec` command that runs `wget -q --spider http://localhost:8080/health`

#### Scenario: Readiness probe uses exec command
- **WHEN** the `server` container definition is examined
- **THEN** a `readinessProbe` is configured with an `exec` command that runs `wget -q --spider http://localhost:8080/health`

#### Scenario: No auth bypass on oauth-proxy
- **WHEN** the oauth-proxy container args are examined
- **THEN** no `--skip-auth-regex` flag is present

### Requirement: TLS Certificate Volume
The oauth-proxy container SHALL mount the service-serving-cert TLS secret for HTTPS termination.

#### Scenario: TLS secret volume is defined
- **WHEN** the Deployment manifest is examined
- **THEN** a volume of type `secret` is defined referencing the TLS serving cert secret name

#### Scenario: TLS secret is mounted in oauth-proxy
- **WHEN** the oauth-proxy container definition is examined
- **THEN** the TLS secret volume is mounted at `/etc/tls/private`

### Requirement: Session Secret Volume
The oauth-proxy container SHALL mount a session secret for encrypting OAuth session cookies.

#### Scenario: Session secret volume is defined
- **WHEN** the Deployment manifest is examined
- **THEN** a volume of type `secret` is defined referencing the session secret

#### Scenario: Session secret is mounted in oauth-proxy
- **WHEN** the oauth-proxy container definition is examined
- **THEN** the session secret volume is mounted at `/etc/proxy/secrets`

### Requirement: Cookie Configuration
The oauth-proxy SHALL use a named, secure cookie with a configurable expiry for session management.

#### Scenario: Cookie name is set
- **WHEN** the oauth-proxy container args are examined
- **THEN** the `--cookie-name` flag is set to `_oauth_proxy`

#### Scenario: Cookie is secure
- **WHEN** the oauth-proxy container args are examined
- **THEN** the `--cookie-secure=true` flag is set

#### Scenario: Cookie secret from file
- **WHEN** the oauth-proxy container args are examined
- **THEN** the `--cookie-secret-file` flag points to the mounted session secret
