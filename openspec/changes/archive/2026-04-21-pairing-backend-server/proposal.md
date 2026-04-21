## Why

The device pairing UI needs a backend HTTP server to process pairing requests. A Go server using the same libraries and structure as claw-signup ensures consistency across the Claw ecosystem and enables maintainers to work across both projects easily.

## What Changes

- Create Go backend server using Echo v5 web framework
- Implement `/pair-device` POST endpoint accepting JSON with device ID
- Use cmd/internal project layout matching claw-signup structure
- Add Cobra CLI for server commands
- Include graceful shutdown and structured logging with slog
- Set up development and build tooling (Makefile, go.mod)

## Capabilities

### New Capabilities
- `http-server`: HTTP server infrastructure with Echo v5, middleware, and graceful shutdown
- `pair-device-endpoint`: POST endpoint to process device pairing requests with JSON payload validation

### Modified Capabilities
<!-- No existing capabilities are being modified -->

## Impact

- New Go module at project root
- New dependencies: Echo v5, Cobra, standard library packages
- New `cmd/` directory with main.go entry point
- New `internal/` directory with handlers and models packages
- No impact on existing UI code
- Server will run independently and can be containerized
