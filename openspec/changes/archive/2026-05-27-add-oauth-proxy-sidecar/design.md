## Context

The claw-device-pairing application runs as a single-container Deployment on OpenShift, serving a React UI and REST APIs on port 8080 with no authentication. There is currently no Service, Route, or Ingress defined — the Deployment exists in isolation. All endpoints (pairing requests, health, static UI) are open to anyone with network access.

OpenShift provides a built-in OAuth server and the `ose-oauth-proxy` image, which is the standard mechanism for adding authentication to internal web applications without modifying application code.

## Goals / Non-Goals

**Goals:**
- Authenticate users via OpenShift OAuth before they can access the device-pairing UI or APIs
- Authorize access by verifying the user can "get pods" in the deployment's namespace (SAR check)
- Expose the application via a TLS-protected Route
- Keep the application code unchanged — all auth happens at the infrastructure layer

**Non-Goals:**
- Fine-grained per-endpoint authorization (e.g., different permissions for pairing vs. health)
- Application-level session management or token handling
- Supporting non-OpenShift clusters (vanilla Kubernetes with oauth2-proxy)
- High availability or horizontal scaling

## Decisions

### Use `ose-oauth-proxy` as a sidecar container

**Decision**: Add the OpenShift oauth-proxy as a sidecar in the same Pod, proxying to the application container on localhost:8080.

**Rationale**: This is OpenShift's recommended pattern. The proxy handles OAuth login, session cookies, and SAR authorization checks. The application doesn't need code changes — it sees authenticated requests forwarded from the proxy.

**Alternatives considered**:
- Application-level OAuth middleware in Go — requires significant code changes (token validation, session store, redirect flow). More control but more maintenance.
- Standalone oauth-proxy Deployment — adds network hop and separate scaling. Sidecar shares the Pod lifecycle and communicates over localhost.

### SAR check: "get pods" in the current namespace

**Decision**: Use `--openshift-sar='{"resource":"pods","verb":"get","namespace":"$(NAMESPACE)"}'` to gate access.

**Rationale**: "get pods" is a common low-privilege check that confirms the user has some access to the namespace. It's a well-established convention for internal tooling on OpenShift.

**Alternatives considered**:
- Custom RBAC verb on a CRD — tighter but requires creating/managing a ClusterRole. Overhead not justified for this use case.
- Cluster-admin check — too restrictive, would lock out most developers.

### TLS with service-serving-cert

**Decision**: Use OpenShift's `service.beta.openshift.io/serving-cert-secret-name` annotation on the Service to automatically provision a TLS certificate. The Route uses `reencrypt` termination.

**Rationale**: Automatic cert provisioning, no manual cert management. The reencrypt Route terminates external TLS at the router and re-encrypts to the oauth-proxy's HTTPS port.

**Alternatives considered**:
- Edge termination with HTTP backend — would expose unencrypted traffic inside the cluster between router and pod. Less secure.
- Manual cert management — operational burden for no benefit.

### Session secret via a static Secret resource

**Decision**: Create a Secret containing a random cookie-secret value that the oauth-proxy uses to encrypt session cookies. This Secret is defined as a Kubernetes manifest.

**Rationale**: The cookie secret must be stable across pod restarts to avoid invalidating user sessions. A pre-created Secret is the simplest approach.

**Alternatives considered**:
- Generating cookie-secret at pod startup — sessions would be lost on every restart.
- Storing in a ConfigMap — secrets should not be in ConfigMaps.

### Exec-based health probes on the server container

**Decision**: Use exec-based liveness and readiness probes on the `server` container that run `wget -q --spider http://localhost:8080/health`, bypassing the oauth-proxy entirely.

**Rationale**: The probe runs inside the `server` container and hits the application's HTTP port directly on localhost. This avoids punching a hole in the auth layer with `--skip-auth-regex` and is more secure — the `/health` endpoint is never exposed unauthenticated through the proxy.

**Alternatives considered**:
- `--skip-auth-regex='^/health$'` on the oauth-proxy — works but creates an unauthenticated path through the proxy, which could be abused for information disclosure.
- HTTP probe targeting the oauth-proxy port — requires the kubelet to authenticate, which it cannot do.

## Risks / Trade-offs

- [WebSocket compatibility] The oauth-proxy may need `--pass-host-header` and WebSocket-specific flags if the handshake protocol uses WebSocket connections through the proxy → Verify WebSocket passthrough works; add `--proxy-websockets=true` flag if needed.
- [NAMESPACE env var in SAR] The `$(NAMESPACE)` in the SAR JSON is expanded by Kubernetes env var substitution at container start → Verify the env var is correctly interpolated. If not, the SAR check will fail and deny all access.
- [Session cookie expiry] Default session TTL may be too short for long pairing workflows → Set `--cookie-expire` to a reasonable value (e.g., 24h).
- [First-time redirect loop] If the user's browser has strict cookie policies, the OAuth redirect loop may fail silently → Document that third-party cookies must be allowed for the Route domain.

## Migration Plan

1. Deploy the ServiceAccount and Secret first (no-op for existing pods)
2. Deploy the updated Deployment with the oauth-proxy sidecar, Service, and Route
3. Verify: access the Route URL, confirm OAuth login redirects, confirm SAR check works
4. Rollback: revert the Deployment to remove the sidecar and delete the Service/Route
