## ADDED Requirements

### Requirement: HTTP Server Initialization
The system SHALL initialize an HTTP server using Echo v5 framework with configurable port.

#### Scenario: Server starts on default port
- **WHEN** the server starts without port flag
- **THEN** the server listens on port 8080

#### Scenario: Server starts on custom port
- **WHEN** the server starts with `--port 9000` flag
- **THEN** the server listens on port 9000

#### Scenario: Invalid port number
- **WHEN** the server starts with port < 1 or > 65535
- **THEN** the server logs an error and exits with code 1

### Requirement: CLI Command Structure
The system SHALL provide a Cobra-based CLI with a `serve` subcommand to start the HTTP server.

#### Scenario: Serve command exists
- **WHEN** user runs `claw-device-pairing serve`
- **THEN** the HTTP server starts

#### Scenario: Port flag is available
- **WHEN** user runs `claw-device-pairing serve --help`
- **THEN** the help text shows `--port` flag with description

### Requirement: Graceful Shutdown
The system SHALL handle SIGINT and SIGTERM signals for graceful shutdown.

#### Scenario: Server receives interrupt signal
- **WHEN** server receives SIGINT or SIGTERM
- **THEN** server completes in-flight requests and shuts down within 10 seconds

#### Scenario: Shutdown timeout exceeded
- **WHEN** shutdown takes longer than 10 seconds
- **THEN** server logs error and exits

### Requirement: Structured Logging
The system SHALL use slog for structured JSON logging.

#### Scenario: Server startup logged
- **WHEN** server starts successfully
- **THEN** log entry contains "server started" message with port number

#### Scenario: Errors are logged
- **WHEN** an error occurs
- **THEN** log entry contains error details with appropriate severity level

### Requirement: CORS Middleware
The system SHALL enable CORS middleware only in development mode when ENV=development.

#### Scenario: CORS enabled in development
- **WHEN** ENV environment variable is "development"
- **THEN** CORS middleware allows origins http://localhost:5173 and http://localhost:5174

#### Scenario: CORS disabled in production
- **WHEN** ENV environment variable is not "development"
- **THEN** CORS middleware is not registered

### Requirement: Health Check Endpoint
The system SHALL provide a health check endpoint at GET /health.

#### Scenario: Health check returns OK
- **WHEN** GET request is made to /health
- **THEN** response status is 200 OK and body contains {"status":"ok"}

### Requirement: Request Logging Middleware
The system SHALL log all incoming HTTP requests using Echo's request logger middleware.

#### Scenario: Requests are logged
- **WHEN** any HTTP request is received
- **THEN** request details are logged including method, path, and status code
