## Why

The backend server and UI currently run separately, requiring users to access the pairing interface through a development server. Serving the UI from the backend's root endpoint simplifies deployment and provides a single entry point for the application.

## What Changes

- Add a GET handler for the root path `/` that serves the built UI files
- Configure static file serving in the Echo server to serve UI assets (JS, CSS, images)
- Ensure UI is served before API routes to avoid conflicts

## Capabilities

### New Capabilities
- `serve-static-ui`: Serve the built Patternfly UI from the backend server's root endpoint and handle static assets

### Modified Capabilities
<!-- No existing spec requirements are changing -->

## Impact

- **cmd/main.go**: Add static file serving middleware and root handler
- **Build process**: May need to ensure UI is built before backend deployment
- **Development workflow**: Developers can access UI directly through backend server
