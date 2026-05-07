# ClawDevicePairingRequest Lifecycle

## Overview

The `ClawDevicePairingRequest` custom resource is used to track device pairing requests in a Kubernetes-native way. This document describes the lifecycle of these resources and how they integrate with the Claw operator.

## CR Creation

When the claw-device-pairing server receives a pairing request via the `/pairing-requests` endpoint:

1. **Request Validation**: The server validates the incoming request ID
2. **CR Name Sanitization**: The request ID is sanitized to comply with DNS-1123 naming requirements
3. **CR Creation**: A `ClawDevicePairingRequest` CR is created with:
   - **Name**: Sanitized request ID
   - **Namespace**: Current pod namespace (from `NAMESPACE` env var)
   - **Spec.RequestID**: Original request ID (preserved)
   - **Spec.Selector**: Label selector for the Claw instance

4. **Error Handling**: If CR creation fails, the pairing request is rejected with a user-friendly error

## CR Spec Fields

### `requestID`
The original pairing request ID as submitted by the client. This field preserves the exact value from the API request, even if the CR name had to be sanitized.

### `selector`
A Kubernetes label selector that identifies which Claw instance this pairing request is for. The selector is constructed from the `claw.sandbox.redhat.com/instance` label value.

**Example:**
```yaml
selector:
  matchLabels:
    claw.sandbox.redhat.com/instance: "prod-instance-1"
```

If the `CLAW_INSTANCE` environment variable is not set, this field will have an empty matchLabels map.

## Operator Integration

The Claw operator is expected to:

1. **Watch** for `ClawDevicePairingRequest` resources in its namespace
2. **Filter** CRs using the selector to identify requests for its instance
3. **Process** pairing requests by:
   - Validating the device
   - Creating or updating device registrations
   - Updating the CR status to reflect approval/rejection
4. **Clean up** processed CRs based on retention policies

## CR Status (Operator-Managed)

The operator is responsible for updating the CR status. Expected status fields include:

```yaml
status:
  phase: "Pending" | "Approved" | "Rejected"
  message: "Human-readable status message"
  approvedAt: "2026-05-06T10:00:00Z"
  approvedBy: "operator-admin-user"
```

**Note**: The pairing server creates CRs but does not update their status. Status management is the responsibility of the Claw operator.

## CR Retention

The operator should implement retention policies for completed pairing requests:

- **Approved requests**: Retain for audit trail (configurable TTL)
- **Rejected requests**: Retain for troubleshooting (shorter TTL)
- **Abandoned requests**: Clean up after timeout period

## AlreadyExists Handling

If a CR with the same name already exists:
- The server logs this as informational (not an error)
- The pairing request is considered successfully submitted
- The operator will process the existing CR

This behavior supports:
- Retry scenarios
- Idempotent pairing operations
- Recovery from transient failures

## Multi-Instance Deployments

In deployments with multiple Claw instances in the same namespace:

1. Each instance's pods have a unique `claw.sandbox.redhat.com/instance` label
2. Each pairing CR includes a selector matching that instance
3. Each operator watches CRs but only processes those matching its selector
4. This prevents cross-instance pairing conflicts

## Troubleshooting

See [troubleshooting.md](troubleshooting.md) for common CR-related issues and solutions.
