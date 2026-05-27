## ADDED Requirements

### Requirement: Dedicated ServiceAccount
A ServiceAccount SHALL be created for the claw-device-pairing Deployment with the OAuth redirect reference annotation.

#### Scenario: ServiceAccount is defined
- **WHEN** the Kubernetes manifests are examined
- **THEN** a ServiceAccount named `claw-device-pairing` exists in `deploy/kubernetes/serviceaccount.yaml`

#### Scenario: OAuth redirect annotation is set
- **WHEN** the ServiceAccount metadata is examined
- **THEN** the `serviceaccounts.openshift.io/oauth-redirectreference.primary` annotation is set with a JSON object referencing the Route name

#### Scenario: Deployment uses the ServiceAccount
- **WHEN** the Deployment manifest is examined
- **THEN** the `spec.template.spec.serviceAccountName` is set to `claw-device-pairing`

### Requirement: Session Cookie Secret
A Kubernetes Secret SHALL be created to store the cookie encryption key used by the oauth-proxy.

#### Scenario: Secret manifest exists
- **WHEN** the Kubernetes manifests are examined
- **THEN** a Secret named `claw-device-pairing-proxy` exists in `deploy/kubernetes/session-secret.yaml`

#### Scenario: Secret contains session data
- **WHEN** the Secret data is examined
- **THEN** it contains a `session_secret` key with a base64-encoded random value
