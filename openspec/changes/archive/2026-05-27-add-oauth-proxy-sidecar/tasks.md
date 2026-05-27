## 1. ServiceAccount and Secrets

- [x] 1.1 Create `deploy/kubernetes/serviceaccount.yaml` with a `claw-device-pairing` ServiceAccount annotated with `serviceaccounts.openshift.io/oauth-redirectreference.primary` pointing to the Route
- [x] 1.2 Create `deploy/kubernetes/session-secret.yaml` with a Secret named `claw-device-pairing-proxy` containing a `session_secret` key with a random base64-encoded value

## 2. Service and Route

- [x] 2.1 Create `deploy/kubernetes/service.yaml` with a Service named `claw-device-pairing` exposing port 443 targeting port 8443 (`oauth-proxy`), annotated with `service.beta.openshift.io/serving-cert-secret-name: claw-device-pairing-tls`
- [x] 2.2 Create `deploy/kubernetes/route.yaml` with a Route named `claw-device-pairing` using reencrypt TLS termination, targeting the Service on port `oauth-proxy` with `insecureEdgeTerminationPolicy: Redirect`

## 3. Deployment Update

- [x] 3.1 Add `serviceAccountName: claw-device-pairing` to the Deployment pod spec
- [x] 3.2 Add volumes to the Deployment: `tls` (secret: `claw-device-pairing-tls`) and `proxy-secret` (secret: `claw-device-pairing-proxy`)
- [x] 3.3 Add exec-based liveness and readiness probes to the `server` container using `wget -q --spider http://localhost:8080/health`
- [x] 3.4 Add the `oauth-proxy` sidecar container to the Deployment with the `ose-oauth-proxy` image, port 8443, volume mounts for TLS (`/etc/tls/private`) and session secret (`/etc/proxy/secrets`), and args for `--upstream=http://localhost:8080`, `--openshift-sar`, `--cookie-secret-file`, `--tls-cert`, `--tls-key`, `--cookie-secure=true` (no `--skip-auth-regex`)
