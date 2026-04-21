## Context

The claw-device-pairing project consists of two components: a React UI (built with Vite) and a Go backend server. Currently, both can be built separately, but there's no unified way to package them for deployment. The claw-signup project has established a successful multi-stage Containerfile pattern that should be replicated here for consistency.

Current state:
- UI builds to `ui/dist/` directory with production-optimized static files
- Go backend compiles to `bin/claw-device-pairing` binary
- No containerization or deployment packaging exists
- Reference implementation available at `../claw-signup/Containerfile`

Constraints:
- Must use multi-stage build to minimize final image size
- Backend currently doesn't serve UI static files, but Containerfile should prepare for this
- Should follow claw-signup's proven pattern for maintainability

## Goals / Non-Goals

**Goals:**
- Create reproducible, multi-stage container build
- Minimize final image size using Alpine base
- Package both UI and backend in single deployable image
- Run container as non-root user for security
- Match claw-signup's Containerfile structure for consistency

**Non-Goals:**
- Modifying backend to serve UI static files (separate change if needed)
- Creating Kubernetes manifests or deployment configs
- Adding health checks or monitoring beyond what exists
- Optimizing build caching strategies (use standard patterns)

## Decisions

### 1. Multi-Stage Build Strategy
**Decision**: Use 3-stage build: UI builder → Go builder → Runtime
**Rationale**: Separates build environments, minimizes final image, proven pattern from claw-signup
**Stages**:
1. Node 20 Alpine for UI build (npm ci + npm run build)
2. Go 1.25 for backend build (CGO_ENABLED=0 for static binary)
3. Alpine latest for runtime (minimal footprint)

**Alternatives considered**:
- Single stage with all tools: Results in >1GB image with unnecessary build tools
- Distroless base: More complex, Alpine well-understood in organization

### 2. Base Image Versions
**Decision**: 
- UI builder: `node:20-alpine`
- Go builder: `mirror.gcr.io/library/golang:1.25`
- Runtime: `alpine:latest`

**Rationale**: Matches claw-signup exactly, Node 20 is LTS, Go 1.25 matches go.mod, Alpine for small size
**Alternatives considered**:
- Ubuntu/Debian bases: Larger size, unnecessary dependencies
- Specific Alpine version tag: `latest` is acceptable for runtime, no version-specific features needed

### 3. Cross-Platform Build Support
**Decision**: Use `--platform=${BUILDPLATFORM}` for builders, `--platform=${TARGETPLATFORM}` for runtime
**Rationale**: Enables building ARM64 images on AMD64 machines (and vice versa), important for M1/M2 Macs
**Build args**: `TARGETOS`, `TARGETARCH` passed to Go build for cross-compilation

### 4. Static Binary Compilation
**Decision**: Build Go binary with `CGO_ENABLED=0`
**Rationale**: Produces fully static binary with no libc dependencies, works on Alpine without glibc
**Build command**:
```bash
CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -o bin/claw-device-pairing ./cmd
```

### 5. File Locations in Runtime Image
**Decision**:
- Binary: `/app/claw-device-pairing`
- UI files: `/app/ui/dist/`
- Working directory: `/app`

**Rationale**: Matches claw-signup, backend can reference ui/dist/ when static file serving is added
**Alternatives considered**:
- Serving UI from /var/www: Less consistent with claw-signup, harder to reference

### 6. User and Permissions
**Decision**: Run as UID/GID 1000:1000 (non-root)
**Rationale**: Security best practice, matches claw-signup, no elevated privileges needed
**Implementation**: `USER 1000:1000` before ENTRYPOINT

### 7. Port Exposure
**Decision**: `EXPOSE 8080`
**Rationale**: Matches default port in cmd/main.go, documents intent for orchestration
**Note**: Not enforced, just metadata for users/tools

### 8. Entrypoint Configuration
**Decision**: `ENTRYPOINT ["/app/claw-device-pairing", "serve"]`
**Rationale**: Uses exec form (no shell), starts server by default, allows --port override via CMD/args
**Flexibility**: Users can override port with `podman run ... --port 9000`

## Risks / Trade-offs

**[Risk]** Backend doesn't currently serve UI static files
→ **Mitigation**: Containerfile prepares structure for future, UI files copied to expected location

**[Risk]** Node modules not cached between builds (slower rebuilds)
→ **Mitigation**: Acceptable for now, can add caching in CI/CD pipeline later, not a Containerfile concern

**[Risk]** Alpine compatibility issues with compiled Go binary
→ **Mitigation**: CGO_ENABLED=0 produces static binary, no libc dependency, well-tested pattern

**[Trade-off]** Using `alpine:latest` vs pinned version
→ **Decision**: Use `latest` for simplicity, matches claw-signup, security updates applied automatically

**[Trade-off]** Single large build stage vs separate stages
→ **Decision**: Separate stages reduce final image from ~1.5GB to ~50MB, worth the complexity

**[Risk]** Build may fail if ui/dist or cmd/ directories don't exist
→ **Mitigation**: Document build prerequisites, container build will fail fast with clear error
