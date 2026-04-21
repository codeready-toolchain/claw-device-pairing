## Why

The claw-device-pairing project has separate UI and backend components that need to be packaged together for deployment. A multi-stage Containerfile enables building both components and combining them into a single production-ready container image, following the proven pattern from claw-signup.

## What Changes

- Create Containerfile with multi-stage build process
- Stage 1: Build UI with Node 20 Alpine, producing production dist bundle
- Stage 2: Build Go backend binary with static compilation
- Stage 3: Create minimal Alpine runtime image with both UI static files and backend binary
- Configure container to run as non-root user (UID 1000)
- Set up proper ENTRYPOINT for server execution

## Capabilities

### New Capabilities
- `container-build`: Multi-stage Containerfile for building and packaging UI and backend into single deployable image

### Modified Capabilities
<!-- No existing capabilities are being modified -->

## Impact

- New Containerfile at project root
- Enables containerized deployment of complete application
- No changes to existing source code or build processes
- Container image can be built with standard `podman build` or `docker build`
- Runtime container will serve UI static files and API endpoints from single process
