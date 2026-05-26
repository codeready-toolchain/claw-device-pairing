## Context

The application currently has an inline `/health` handler in `cmd/main.go` that returns `{"status": "ok"}`. This needs to be enhanced to include version metadata (`commit_hash`, `build_time`) for OpenShift liveness/readiness probes and operational visibility. The version variables are already set via `ldflags` at build time in `internal/version/version.go`.

## Goals / Non-Goals

**Goals:**
- Provide a `/health` endpoint that returns commit hash and build time alongside status
- Follow existing handler patterns (dedicated handler file in `internal/handlers/`)
- Keep the endpoint lightweight with no external dependencies (no K8s client, no DB)

**Non-Goals:**
- Deep health checks (database connectivity, downstream service checks)
- Separate liveness vs readiness probe endpoints
- Metrics or Prometheus integration

## Decisions

**1. Dedicated handler file vs inline handler**
Create `internal/handlers/health.go` with a standalone handler function rather than keeping it inline in `cmd/main.go`. This follows the project pattern where handlers live in `internal/handlers/` and keeps `cmd/main.go` focused on wiring.

**2. Standalone function vs struct-based handler**
Use a standalone `HandleHealth` function (not attached to a struct) since the health handler has no dependencies — it only reads package-level variables from `internal/version`. This contrasts with `PairingRequestsHandler` which needs a K8s client injected.

**3. Response structure**
Return a flat JSON object with three fields: `status`, `commit_hash`, `build_time`. Use `snake_case` keys to match the existing API convention (see `models.ErrorResponse`).

## Risks / Trade-offs

- [Minimal risk] The handler reads package-level vars that are set once at build time — no concurrency or state concerns.
