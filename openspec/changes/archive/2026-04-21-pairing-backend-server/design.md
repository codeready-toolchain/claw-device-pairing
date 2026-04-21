## Context

The claw-device-pairing project has a React UI but no backend server. The claw-signup project provides a proven reference implementation using Echo v5, Cobra CLI, and structured logging. Reusing this architecture ensures consistency and reduces learning curve for developers working across Claw projects.

Current state:
- UI exists at `ui/` with React + Vite
- No backend server or API endpoints
- claw-signup reference at `../claw-signup` with established patterns

Constraints:
- Must use same libraries as claw-signup for ecosystem consistency
- Single POST endpoint initially (extensible for future endpoints)
- No database or persistence layer in this change

## Goals / Non-Goals

**Goals:**
- HTTP server with `/pair-device` POST endpoint accepting JSON
- Project structure matching claw-signup (cmd/internal layout)
- Graceful shutdown and structured logging
- Input validation for device ID field
- Proper HTTP status codes and error responses

**Non-Goals:**
- Database integration or persistence layer
- Authentication/authorization (future work)
- WebSocket support for real-time pairing status
- Integration with Kubernetes or external services
- Production deployment configuration (containers, manifests)

## Decisions

### 1. Web Framework
**Decision**: Use Echo v5 (same as claw-signup)
**Rationale**: Proven in claw-signup, excellent performance, clean middleware support, type-safe context
**Alternatives considered**:
- net/http stdlib: More verbose, missing middleware ecosystem
- Gin: Different API surface, no benefit over Echo for this use case
- Chi: Good alternative, but Echo already validated in claw-signup

### 2. Project Structure
**Decision**: Mirror claw-signup layout
```
claw-device-pairing/
├── cmd/
│   └── main.go              (entry point, CLI commands)
├── internal/
│   ├── handlers/            (HTTP handlers)
│   │   ├── pairing.go
│   │   └── pairing_test.go
│   ├── models/              (request/response types)
│   │   └── pairing.go
│   └── logger/              (slog initialization)
│       └── logger.go
├── go.mod
├── go.sum
└── Makefile
```
**Rationale**: Developers familiar with claw-signup can navigate immediately, established patterns for testing and organization
**Alternatives considered**:
- Flat structure: Harder to maintain as project grows
- pkg/ instead of internal/: internal/ signals non-reusable packages, appropriate here

### 3. Request/Response Model
**Decision**: Define models in `internal/models/pairing.go`
```go
type PairDeviceRequest struct {
    ID string `json:"id" validate:"required"`
}

type PairDeviceResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

type ErrorResponse struct {
    Error  string `json:"error"`
    Status string `json:"status"`
}
```
**Rationale**: Centralized type definitions, easy to validate, clear contract
**Validation**: Use Echo's c.Bind() for JSON parsing, custom validation for ID field

### 4. Handler Pattern
**Decision**: Use handler struct with constructor (same as claw-signup)
```go
type PairingHandler struct {
    // Future: add dependencies like database client
}

func NewPairingHandler() *PairingHandler {
    return &PairingHandler{}
}

func (h *PairingHandler) HandlePairDevice(c echo.Context) error {
    // Parse, validate, process
}
```
**Rationale**: Extensible for dependency injection, testable, follows claw-signup pattern
**Alternatives considered**:
- Plain functions: Harder to inject dependencies later
- Global state: Not testable, causes issues in concurrent tests

### 5. Error Handling
**Decision**: Return JSON error responses with appropriate HTTP status codes
- 400 Bad Request: Invalid JSON or missing/empty ID field
- 500 Internal Server Error: Unexpected server errors
- 200 OK: Successful pairing request accepted

**Rationale**: Clear contract for frontend, follows REST conventions
**Logging**: Use slog for structured logging (all errors logged before returning)

### 6. CLI Structure
**Decision**: Use Cobra with `serve` subcommand (matching claw-signup)
```bash
claw-device-pairing serve --port 8080
```
**Rationale**: Consistent with claw-signup, extensible for future commands
**Flags**: `--port` for configuring listen port (default 8080)

### 7. CORS Configuration
**Decision**: Enable CORS middleware in development mode only (ENV=development)
```go
AllowOrigins: []string{"http://localhost:5173"}  // Vite dev server
```
**Rationale**: UI runs on different port in dev, production serves static files from same origin
**Security**: Only enabled when ENV env var is set to development

## Risks / Trade-offs

**[Risk]** No authentication on `/pair-device` endpoint
→ **Mitigation**: Acceptable for initial implementation, document as future work, endpoint can be extended with auth middleware

**[Risk]** No database means pairing requests are not persisted
→ **Mitigation**: Intentional for this phase, handler logs request and returns success (future change adds persistence)

**[Trade-off]** Duplicating logger package from claw-signup vs. extracting to shared library
→ **Decision**: Duplicate for now (single file, ~20 lines), extract if 3+ projects need it

**[Trade-off]** Echo v5 vs newer frameworks
→ **Decision**: Echo v5 is proven, stable, and matches claw-signup. Consistency > novelty.

**[Risk]** No integration tests in initial implementation
→ **Mitigation**: Include unit tests for handler with mocked Echo context, add integration tests in follow-up
