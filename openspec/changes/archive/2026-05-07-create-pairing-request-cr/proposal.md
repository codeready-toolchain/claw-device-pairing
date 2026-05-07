## Why

Enable declarative device pairing management by creating ClawDevicePairingRequest CRs from pairing handler. This allows Kubernetes-native tracking and reconciliation of pairing requests through Custom Resources.

## What Changes

- Create `ClawDevicePairingRequest` CR when `HandlePairDevice` receives a PairingRequest
- Name the CR using the incoming `requestId`
- Add namespace environment variable to deployment configuration
- Populate CR spec with `RequestID` and `Selector` fields
- Implement mechanism to retrieve current `claw.sandbox.redhat.com/instance` label for selector
- Add Kubernetes client support for CR creation in pairing handler

## Capabilities

### New Capabilities

- `pairing-request-cr`: Creation and management of ClawDevicePairingRequest custom resources from the pairing handler, including namespace detection and instance label retrieval

### Modified Capabilities

<!-- No existing capabilities are being modified -->

## Impact

- `HandlePairDevice` method in backend server
- Deployment manifests (add namespace environment variable)
- New ClawDevicePairingRequest CRD or existing CRD spec
- Kubernetes client initialization and CR creation logic
- Pod metadata access for instance label retrieval
