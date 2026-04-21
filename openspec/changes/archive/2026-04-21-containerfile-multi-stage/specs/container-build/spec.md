## ADDED Requirements

### Requirement: Multi-Stage Build Structure
The Containerfile SHALL use a multi-stage build with three distinct stages: UI builder, Go builder, and runtime.

#### Scenario: Three stages are defined
- **WHEN** the Containerfile is parsed
- **THEN** it contains exactly three FROM statements with appropriate aliases

#### Scenario: Stages are named correctly
- **WHEN** examining the Containerfile
- **THEN** stages are named `ui-builder`, `server-builder`, and the final runtime stage

### Requirement: UI Build Stage
The UI build stage SHALL use Node 20 Alpine as the base image and produce production-ready static files.

#### Scenario: Node 20 Alpine base image
- **WHEN** the UI builder stage is defined
- **THEN** it uses `node:20-alpine` as the base image

#### Scenario: Dependencies are installed
- **WHEN** UI build executes
- **THEN** package.json and package-lock.json are copied and `npm ci` is run

#### Scenario: UI is built
- **WHEN** UI source is copied
- **THEN** `npm run build` produces output in dist/ directory

#### Scenario: Build platform is specified
- **WHEN** UI builder FROM statement is defined
- **THEN** it includes `--platform=${BUILDPLATFORM}` for cross-platform builds

### Requirement: Go Build Stage
The Go build stage SHALL compile a static Go binary for the target platform.

#### Scenario: Go 1.25 base image
- **WHEN** the server builder stage is defined
- **THEN** it uses `mirror.gcr.io/library/golang:1.25` as the base image

#### Scenario: Build arguments are accepted
- **WHEN** the server builder stage is defined
- **THEN** it declares ARG for TARGETOS and TARGETARCH

#### Scenario: Dependencies are cached
- **WHEN** Go build executes
- **THEN** go.mod and go.sum are copied before source code and `go mod download` is run

#### Scenario: Static binary is produced
- **WHEN** Go build command executes
- **THEN** CGO_ENABLED=0 is set and GOOS/GOARCH are configured for target platform

#### Scenario: Binary is placed correctly
- **WHEN** Go build completes
- **THEN** binary is output to `bin/claw-device-pairing`

### Requirement: Runtime Stage
The runtime stage SHALL create a minimal image with both the Go binary and UI static files.

#### Scenario: Alpine runtime base
- **WHEN** the runtime stage is defined
- **THEN** it uses `alpine:latest` as the base image

#### Scenario: Target platform is specified
- **WHEN** runtime FROM statement is defined
- **THEN** it includes `--platform=${TARGETPLATFORM}` for correct architecture

#### Scenario: Binary is copied from builder
- **WHEN** runtime image is constructed
- **THEN** binary is copied from server-builder stage to `/app/claw-device-pairing`

#### Scenario: UI files are copied from builder
- **WHEN** runtime image is constructed
- **THEN** UI dist directory is copied from ui-builder stage to `/app/ui/dist`

### Requirement: Security Configuration
The runtime container SHALL run as a non-root user for security.

#### Scenario: Non-root user is set
- **WHEN** runtime stage is configured
- **THEN** USER directive sets UID and GID to 1000:1000

#### Scenario: User is set before entrypoint
- **WHEN** Containerfile is parsed
- **THEN** USER directive appears before ENTRYPOINT

### Requirement: Port Configuration
The Containerfile SHALL expose port 8080 for the HTTP server.

#### Scenario: Port 8080 is exposed
- **WHEN** examining the Containerfile
- **THEN** EXPOSE directive declares port 8080

### Requirement: Entrypoint Configuration
The runtime container SHALL be configured to run the server with the serve command.

#### Scenario: Entrypoint uses exec form
- **WHEN** ENTRYPOINT is defined
- **THEN** it uses JSON array form for direct execution without shell

#### Scenario: Server is started with serve command
- **WHEN** container starts
- **THEN** ENTRYPOINT executes `/app/claw-device-pairing serve`

### Requirement: Working Directory
The runtime container SHALL use /app as the working directory.

#### Scenario: Working directory is set
- **WHEN** runtime stage is configured
- **THEN** WORKDIR directive sets working directory to /app

### Requirement: Build Optimization
The Containerfile SHALL optimize build caching by copying dependency files before source code.

#### Scenario: UI dependencies copied separately
- **WHEN** UI builder stage executes
- **THEN** package files are copied before source code to cache npm install

#### Scenario: Go dependencies copied separately
- **WHEN** server builder stage executes
- **THEN** go.mod and go.sum are copied before source code to cache module download
