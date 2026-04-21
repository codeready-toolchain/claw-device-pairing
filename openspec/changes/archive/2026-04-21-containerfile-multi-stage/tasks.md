## 1. Create Containerfile

- [x] 1.1 Create Containerfile at project root
- [x] 1.2 Add comment header explaining multi-stage build approach

## 2. UI Build Stage

- [x] 2.1 Define Stage 1 with FROM node:20-alpine AS ui-builder and --platform=${BUILDPLATFORM}
- [x] 2.2 Set WORKDIR to /app/ui
- [x] 2.3 Copy ui/package*.json to working directory
- [x] 2.4 Run npm ci to install dependencies
- [x] 2.5 Copy ui/ directory to working directory
- [x] 2.6 Run npm run build to create production bundle

## 3. Go Build Stage

- [x] 3.1 Define Stage 2 with FROM mirror.gcr.io/library/golang:1.25 AS server-builder and --platform=${BUILDPLATFORM}
- [x] 3.2 Declare ARG for TARGETOS
- [x] 3.3 Declare ARG for TARGETARCH
- [x] 3.4 Set WORKDIR to /app
- [x] 3.5 Copy go.mod and go.sum to working directory
- [x] 3.6 Run go mod download to cache dependencies
- [x] 3.7 Copy cmd/ directory to working directory
- [x] 3.8 Copy internal/ directory to working directory
- [x] 3.9 Run go build with CGO_ENABLED=0, GOOS=${TARGETOS:-linux}, GOARCH=${TARGETARCH} to ./bin/claw-device-pairing ./cmd

## 4. Runtime Stage

- [x] 4.1 Define Stage 3 with FROM alpine:latest and --platform=${TARGETPLATFORM}
- [x] 4.2 Set WORKDIR to /app
- [x] 4.3 Copy binary from server-builder stage: /app/bin/claw-device-pairing to /app/claw-device-pairing
- [x] 4.4 Copy UI files from ui-builder stage: /app/ui/dist to /app/ui/dist
- [x] 4.5 Set USER to 1000:1000
- [x] 4.6 Add EXPOSE 8080
- [x] 4.7 Set ENTRYPOINT to ["/app/claw-device-pairing", "serve"]

## 5. Verification

- [x] 5.1 Build container image: podman build -t claw-device-pairing:latest .
- [x] 5.2 Verify image builds successfully without errors
- [x] 5.3 Check image size is reasonable (<100MB for runtime image)
- [x] 5.4 Run container: podman run -p 8080:8080 claw-device-pairing:latest
- [x] 5.5 Verify server starts and logs "server started" message
- [x] 5.6 Test health endpoint: curl http://localhost:8080/health
- [x] 5.7 Verify container runs as UID 1000 (check with podman exec)
- [x] 5.8 Test cross-platform build with --platform=linux/arm64
- [x] 5.9 Verify graceful shutdown with podman stop
