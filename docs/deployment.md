# Deployment Guide

## Environment Variables

The claw-device-pairing server requires the following environment variables to be configured in the deployment:

### Required Environment Variables

Both environment variables are **required** and must be set. The server will fail to start if either is missing.

| Variable | Description | Source |
|----------|-------------|--------|
| `NAMESPACE` | The Kubernetes namespace where the server is running | Downward API (`metadata.namespace`) |
| `CLAW_INSTANCE` | The value of the `claw.sandbox.redhat.com/instance` label for the current Claw instance | Downward API (`metadata.labels`) |

**Important:** The server validates these environment variables at startup and will panic with a clear error message if either is missing.

### Configuration in Deployment Manifest

The environment variables are configured using the Kubernetes Downward API in the deployment manifest:

```yaml
env:
- name: NAMESPACE
  valueFrom:
    fieldRef:
      fieldPath: metadata.namespace
- name: CLAW_INSTANCE
  valueFrom:
    fieldRef:
      fieldPath: metadata.labels['claw.sandbox.redhat.com/instance']
```

### Deployment Example

See `deploy/kubernetes/deployment.yaml` for a complete deployment example.

## ClawDevicePairingRequest Custom Resource

The server creates `ClawDevicePairingRequest` custom resources when it receives pairing requests. These CRs are used by the Claw operator to track and reconcile device pairing requests.

### CR Spec

```yaml
apiVersion: claw.sandbox.redhat.com/v1alpha1
kind: ClawDevicePairingRequest
metadata:
  name: <sanitized-request-id>
  namespace: <namespace>
spec:
  requestID: <original-request-id>
  selector:
    matchLabels:
      claw.sandbox.redhat.com/instance: <instance-value>
```

### CR Naming

The CR name is derived from the pairing request ID by:
1. Converting to lowercase
2. Replacing invalid characters with hyphens
3. Removing leading/trailing hyphens
4. Truncating to 63 characters maximum

## RBAC Requirements

The service account used by the server must have permissions to create `ClawDevicePairingRequest` resources:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: claw-device-pairing
rules:
- apiGroups: ["claw.sandbox.redhat.com"]
  resources: ["clawdevicepairingrequests"]
  verbs: ["create", "get", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: claw-device-pairing
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: claw-device-pairing
subjects:
- kind: ServiceAccount
  name: claw-device-pairing
```

## Local Development

The server is designed to run inside a Kubernetes cluster and requires in-cluster configuration:
- **You must set both `NAMESPACE` and `CLAW_INSTANCE` environment variables** - the server will not start without them
- **The server requires in-cluster service account credentials** - it will fail to start if these are not available
- The server uses `rest.InClusterConfig()` which requires `KUBERNETES_SERVICE_HOST` and `KUBERNETES_SERVICE_PORT` environment variables
- Service account credentials must be mounted at `/var/run/secrets/kubernetes.io/serviceaccount/`

**Note:** The server cannot run in local development mode outside of a Kubernetes cluster. For local testing, you must deploy it to a cluster or use a tool like `kubectl port-forward` to access the deployed service.

### Running in a Development Cluster

Deploy to a development cluster with the required environment variables:

```yaml
env:
- name: NAMESPACE
  valueFrom:
    fieldRef:
      fieldPath: metadata.namespace
- name: CLAW_INSTANCE
  valueFrom:
    fieldRef:
      fieldPath: metadata.labels['claw.sandbox.redhat.com/instance']
```
