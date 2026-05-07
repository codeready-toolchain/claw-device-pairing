## Context

The `HandlePairDevice` method currently processes incoming pairing requests via WebSocket, validating and responding to the client. This change introduces Kubernetes CR creation to enable declarative tracking and operator-driven reconciliation of pairing requests. The ClawDevicePairingRequest CRD may already exist or needs to be created separately.

## Goals / Non-Goals

**Goals:**
- Create ClawDevicePairingRequest CR for each pairing request received
- Name CR using incoming requestId for traceability
- Populate CR spec with RequestID and Selector fields
- Retrieve current namespace via environment variable
- Retrieve claw.sandbox.redhat.com/instance label from pod metadata
- Initialize Kubernetes client for CR creation

**Non-Goals:**
- CR reconciliation logic (handled by separate controller)
- Modifying existing pairing flow beyond adding CR creation and error handling
- Retry logic for CR creation (fail fast, let user retry)
- CRD definition or installation

## Decisions

**1. Namespace from Environment Variable**
Use `NAMESPACE` environment variable set in deployment via downward API.
- **Why**: Standard Kubernetes pattern, simple configuration, no runtime API calls needed
- **Alternative**: Query Kubernetes API for pod namespace → adds latency and requires additional RBAC permissions

**2. Instance Label from Pod Metadata**
Retrieve `claw.sandbox.redhat.com/instance` label value via downward API environment variable.
- **Why**: Avoids pod metadata API calls, available at startup
- **Alternative**: Read `/var/run/secrets/kubernetes.io/podinfo/labels` → requires volume mount
- **Alternative**: Query Kubernetes API → adds latency and RBAC requirements

**3. CR Naming**
Use `requestId` directly as CR name (after validation for DNS-1123 compliance).
- **Why**: Direct correlation between request and CR, no additional mapping needed
- **Alternative**: Generate UUID → loses direct traceability to request

**4. Kubernetes Client Initialization**
Initialize in-cluster Kubernetes client at server startup (not per-request).
- **Why**: Client reuse, connection pooling, avoid initialization overhead per request
- **Alternative**: Lazy initialization on first request → delays first pairing

**5. CR Creation Failure Handling**
Fail the pairing request and return error to UI when CR creation fails.
- **Why**: CR creation is essential for tracking; pairing without CR defeats the purpose of declarative management
- **Alternative**: Log but continue → allows pairing without tracking, creates inconsistent state
- **Implementation**: Return succinct user-facing error message ("Something wrong happened, could not pair the device") while logging detailed error server-side

## Risks / Trade-offs

**[Risk: Missing environment variables in development]**
→ Mitigation: Provide sensible defaults (`default` namespace) and clear error messages when instance label is missing

**[Risk: RequestId not DNS-1123 compliant]**
→ Mitigation: Validate and sanitize requestId before using as CR name (lowercase, alphanumeric with hyphens, max 63 chars)

**[Risk: Instance label not present on pod]**
→ Mitigation: Log warning and create CR without selector, or use empty selector value

**[Risk: CR already exists with same requestId]**
→ Mitigation: Handle AlreadyExists error gracefully, optionally update existing CR

**[Trade-off: CR creation adds latency to pairing flow]**
→ Accept minor latency increase (<50ms typical) as acceptable for tracking benefit
