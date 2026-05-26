## Why

The application needs a `/health` endpoint that reports build version information (commit hash and build time) so it can be used as liveness and readiness probes when deployed on OpenShift. A basic `/health` handler already exists inline in `cmd/main.go` but only returns `{"status": "ok"}` — it needs to be moved to a proper handler and enhanced with version metadata.

## What Changes

- Move the `/health` handler from an inline function in `cmd/main.go` to a dedicated handler in `internal/handlers/`
- Return a JSON response with `commit_hash` and `build_time` fields sourced from `internal/version/version.go`
- Keep the `200 OK` status code for successful responses

## Capabilities

### New Capabilities
- `health-endpoint`: Health check handler returning version info (commit hash, build time) for Kubernetes liveness/readiness probes

### Modified Capabilities

## Impact

- `cmd/main.go`: Replace inline `/health` handler with call to new handler
- `internal/handlers/`: New health handler file
- `internal/version/version.go`: Read-only dependency (no changes needed)
