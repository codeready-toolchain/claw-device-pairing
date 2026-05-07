# Multi-stage build for claw-device-pairing
# Stage 1: Build UI with Node
# Stage 2: Build Go server binary
# Stage 3: Create minimal runtime image

# Stage 1: Build UI
FROM --platform=${BUILDPLATFORM} node:20-alpine AS ui-builder

WORKDIR /app/ui

# Copy package files
COPY ui/package*.json ./

# Install dependencies
RUN npm ci

# Copy UI source
COPY ui/ ./

# Build UI
RUN npm run build

# Stage 2: Build Go server
FROM --platform=${BUILDPLATFORM} mirror.gcr.io/library/golang:1.25 AS server-builder
ARG TARGETOS
ARG TARGETARCH
ARG COMMIT_HASH=unknown
ARG BUILD_TIME=unknown

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Build server binary with version information
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build \
    -ldflags "-X 'github.com/xcoulon/claw-device-pairing/internal/version.CommitHash=${COMMIT_HASH}' \
              -X 'github.com/xcoulon/claw-device-pairing/internal/version.BuildTime=${BUILD_TIME}'" \
    -o bin/claw-device-pairing ./cmd

# Stage 3: Runtime
FROM --platform=${TARGETPLATFORM} alpine:latest

WORKDIR /app

# Copy server binary from builder
COPY --from=server-builder /app/bin/claw-device-pairing /app/claw-device-pairing

# Copy UI static files from UI builder
COPY --from=ui-builder /app/ui/dist /app/ui/dist

USER 1000:1000

# Expose port
EXPOSE 8080

# Run server
ENTRYPOINT ["/app/claw-device-pairing", "serve"]
