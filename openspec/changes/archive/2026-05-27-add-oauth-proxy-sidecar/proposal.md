## Why

The device-pairing application currently exposes its APIs and UI without any authentication or authorization. Anyone with network access to the pod can submit pairing requests or access the UI. On OpenShift, the standard approach is to add an `oauth-proxy` sidecar that authenticates users via the cluster's built-in OAuth server and authorizes them using a SubjectAccessReview (SAR) check — in this case, verifying the user can "get pods" in the deployment's namespace.

## What Changes

- Add an `oauth-proxy` sidecar container to the Deployment that terminates TLS and authenticates users via OpenShift OAuth
- Configure a SAR check (`--openshift-sar`) to authorize only users who can "get pods" in the current namespace
- Add a Service to expose the oauth-proxy port (443) instead of the application port (8080) directly
- Add a Route with TLS reencrypt termination pointing to the oauth-proxy Service
- Create a ServiceAccount with the `oauth-redirectreference` annotation required by the OAuth proxy
- Add a TLS-serving-cert secret (via OpenShift's `service.beta.openshift.io/serving-cert-secret-name` annotation on the Service)
- **BREAKING**: Direct unauthenticated access to the application on port 8080 is no longer possible from outside the pod

## Capabilities

### New Capabilities

- `oauth-proxy-sidecar`: OAuth proxy sidecar container configuration, SAR authorization, and session cookie settings
- `oauth-service-account`: ServiceAccount with oauth-redirectreference annotation
- `oauth-tls-route`: Service and Route configuration for TLS-protected access to the oauth proxy

### Modified Capabilities

_None — existing application code and specs are unchanged. The oauth proxy sits in front of the app at the infrastructure level._

## Impact

- `deploy/kubernetes/deployment.yaml`: Add oauth-proxy sidecar container, volume mounts for TLS cert and session secret
- `deploy/kubernetes/service.yaml`: New file — Service exposing oauth-proxy port 443 with serving-cert annotation
- `deploy/kubernetes/route.yaml`: New file — Route with reencrypt TLS termination
- `deploy/kubernetes/serviceaccount.yaml`: New file — ServiceAccount with oauth redirect annotation
- No application code changes required
- New dependency on `registry.redhat.io/openshift4/ose-oauth-proxy` image
