## Context

The device pairing UI is built with Vite and currently uses absolute paths for asset loading. The build output assumes deployment at the root path (`/`), which limits deployment flexibility. The Go server may serve the UI from different paths (e.g., `/pair-device/`), so the UI needs to support deployment at any base path.

Vite provides a `base` configuration option that controls how asset paths are generated in the build output. By default, it's set to `/`, which generates absolute paths.

## Goals / Non-Goals

**Goals:**
- Configure Vite to generate relative asset paths in the production build
- Support deployment of the UI at any base path without rebuilding
- Maintain existing development server functionality

**Non-Goals:**
- Modify runtime application code (this is a build-time configuration change only)
- Change how the Go server serves the UI (server-side routing is separate)
- Support multiple base paths simultaneously (single deployment = single base path)

## Decisions

### Decision 1: Use Vite `base: './'` configuration

**Rationale**: Vite's `base` option with value `'./'` generates relative paths for all assets (JS, CSS, images) in the production build. This is the standard Vite approach for supporting flexible deployment paths.

**Alternatives considered**:
- Runtime base path detection: Would require JavaScript to detect and rewrite paths at runtime, adding complexity and potential for errors
- Environment variable configuration: Would require rebuilding for different deployment paths, defeating the purpose of flexible deployment

**Decision**: Set `base: './'` in `vite.config.js`. This is the simplest and most maintainable approach.

### Decision 2: Apply to production builds only

**Rationale**: Development server works fine with default configuration. The `base` setting only affects the production build output.

**Implementation**: The `base: './'` setting is active for all builds, but its effect is primarily visible in production builds where asset paths are bundled and referenced.

## Risks / Trade-offs

**[Risk]** Relative paths may behave differently if the application uses client-side routing with history mode
→ **Mitigation**: This application is a simple single-page UI without client-side routing, so this is not a concern. If routing is added later, verify asset paths work correctly.

**[Risk]** Development and production builds could diverge in behavior
→ **Mitigation**: Vite's dev server handles relative paths correctly, and we'll verify the production build works as expected before deployment.

**[Trade-off]** All assets must be served from the same directory structure
→ **Acceptable**: The Go server already serves all UI assets from a consistent base path, so this constraint is already met.
