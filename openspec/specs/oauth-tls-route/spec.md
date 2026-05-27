## Purpose

Service and Route configuration for TLS-protected access to the oauth proxy.

## Requirements

### Requirement: Service for OAuth Proxy
A Kubernetes Service SHALL expose the oauth-proxy sidecar's HTTPS port.

#### Scenario: Service is defined
- **WHEN** the Kubernetes manifests are examined
- **THEN** a Service named `claw-device-pairing` exists in `deploy/kubernetes/service.yaml`

#### Scenario: Service targets port 8443
- **WHEN** the Service spec is examined
- **THEN** the Service targets port 8443 named `oauth-proxy` on pods matching `app: claw-device-pairing`

#### Scenario: Service exposes port 443
- **WHEN** the Service spec is examined
- **THEN** the Service exposes port 443 externally

#### Scenario: Serving cert annotation is set
- **WHEN** the Service metadata is examined
- **THEN** the `service.beta.openshift.io/serving-cert-secret-name` annotation is set to `claw-device-pairing-tls`

### Requirement: Route with TLS Reencrypt
An OpenShift Route SHALL expose the Service externally with reencrypt TLS termination.

#### Scenario: Route is defined
- **WHEN** the Kubernetes manifests are examined
- **THEN** a Route named `claw-device-pairing` exists in `deploy/kubernetes/route.yaml`

#### Scenario: Route uses reencrypt TLS
- **WHEN** the Route spec is examined
- **THEN** the `tls.termination` is set to `Reencrypt`

#### Scenario: Route targets the Service
- **WHEN** the Route spec is examined
- **THEN** the Route targets the `claw-device-pairing` Service on port `oauth-proxy`

#### Scenario: Route uses insecure redirect
- **WHEN** the Route spec is examined
- **THEN** the `tls.insecureEdgeTerminationPolicy` is set to `Redirect`
