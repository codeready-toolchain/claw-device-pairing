## 1. Project Initialization

- [x] 1.1 Initialize Go module with `go mod init github.com/xcoulon/claw-device-pairing`
- [x] 1.2 Create `cmd/` directory at project root
- [x] 1.3 Create `internal/handlers/` directory for HTTP handlers
- [x] 1.4 Create `internal/models/` directory for request/response types
- [x] 1.5 Create `internal/logger/` directory for logging setup

## 2. Dependencies

- [x] 2.1 Add Echo v5 dependency: `go get github.com/labstack/echo/v5`
- [x] 2.2 Add Cobra dependency: `go get github.com/spf13/cobra`
- [x] 2.3 Run `go mod tidy` to resolve all dependencies

## 3. Models Package

- [x] 3.1 Create `internal/models/pairing.go` file
- [x] 3.2 Define `PairDeviceRequest` struct with `ID string` field and JSON tag
- [x] 3.3 Define `PairDeviceResponse` struct with `Status` and `Message` fields
- [x] 3.4 Define `ErrorResponse` struct with `Error` and `Status` fields
- [x] 3.5 Add JSON struct tags to all model fields

## 4. Logger Package

- [x] 4.1 Create `internal/logger/logger.go` file
- [x] 4.2 Implement `Init()` function to configure slog with JSON handler
- [x] 4.3 Set log level to Info for production, Debug for development
- [x] 4.4 Export logger initialization for use in main.go

## 5. Pairing Handler

- [x] 5.1 Create `internal/handlers/pairing.go` file
- [x] 5.2 Define `PairingHandler` struct
- [x] 5.3 Implement `NewPairingHandler()` constructor function
- [x] 5.4 Implement `HandlePairDevice(c echo.Context) error` method
- [x] 5.5 Add JSON request parsing with `c.Bind()` in handler
- [x] 5.6 Add validation for non-empty ID field (trim whitespace)
- [x] 5.7 Return 400 Bad Request for invalid JSON or empty ID
- [x] 5.8 Return 200 OK with success response for valid requests
- [x] 5.9 Log all errors and successful requests with slog

## 6. Handler Tests

- [x] 6.1 Create `internal/handlers/pairing_test.go` file
- [x] 6.2 Write test for valid pairing request (200 OK)
- [x] 6.3 Write test for invalid JSON (400 Bad Request)
- [x] 6.4 Write test for missing ID field (400 Bad Request)
- [x] 6.5 Write test for empty ID field (400 Bad Request)
- [x] 6.6 Write test for whitespace-only ID (400 Bad Request)

## 7. Main Server Setup

- [x] 7.1 Create `cmd/main.go` file
- [x] 7.2 Define root Cobra command for CLI
- [x] 7.3 Define `serve` subcommand with `--port` flag (default 8080)
- [x] 7.4 Add port validation (1-65535 range)
- [x] 7.5 Initialize logger in serve command
- [x] 7.6 Create Echo instance
- [x] 7.7 Add CORS middleware with ENV=development check
- [x] 7.8 Add request logging middleware
- [x] 7.9 Register `/pair-device` POST route with handler
- [x] 7.10 Register `/health` GET route returning {"status":"ok"}
- [x] 7.11 Create http.Server with graceful shutdown support
- [x] 7.12 Add signal handling for SIGINT and SIGTERM
- [x] 7.13 Implement graceful shutdown with 10-second timeout
- [x] 7.14 Add main() function to execute root command

## 8. Build and Development Tooling

- [x] 8.1 Create Makefile with `build` target
- [x] 8.2 Add `test` target to Makefile running `go test ./...`
- [x] 8.3 Add `run` target to Makefile running `go run cmd/main.go serve`
- [x] 8.4 Add `lint` target with `go vet ./...`
- [x] 8.5 Add `.gitignore` entries for `bin/` and `tmp/` directories

## 9. Verification

- [x] 9.1 Run `make build` and verify binary compiles successfully
- [x] 9.2 Run `make test` and verify all tests pass
- [x] 9.3 Start server with `make run` or `./bin/claw-device-pairing serve`
- [x] 9.4 Test health endpoint: `curl http://localhost:8080/health`
- [x] 9.5 Test pair-device with valid request: `curl -X POST -H "Content-Type: application/json" -d '{"id":"test-device"}' http://localhost:8080/pair-device`
- [x] 9.6 Test pair-device with empty id: `curl -X POST -H "Content-Type: application/json" -d '{"id":""}' http://localhost:8080/pair-device`
- [x] 9.7 Test pair-device with invalid JSON: `curl -X POST -H "Content-Type: application/json" -d '{invalid}' http://localhost:8080/pair-device`
- [x] 9.8 Verify graceful shutdown with Ctrl+C (SIGINT)
- [x] 9.9 Verify CORS headers present in development mode (ENV=development)
