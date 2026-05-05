## Why

The device pairing UI is currently built with absolute paths for asset loading, which assumes deployment at the root path (`/`). To support flexible deployment scenarios where the application may be served from a subpath (e.g., `/pair-device/`, `/ui/`, or any custom path), the build configuration needs to generate relative asset references.

## What Changes

- Configure Vite to use relative paths for all asset imports (JS, CSS, images)
- Update `vite.config.js` with `base: './'` to generate relative paths in the build output
- Ensure the production build uses relative references for all static assets

## Capabilities

### New Capabilities
- `vite-relative-paths`: Build configuration for Vite to generate relative asset paths supporting flexible deployment

### Modified Capabilities
- `pairing-ui`: Build output now supports deployment at any base path, not just root

## Impact

- Vite configuration file (`ui/vite.config.js`)
- Build output structure and asset references
- Deployment flexibility - application can now be served from any base path
- No runtime code changes required - this is purely a build-time configuration change
