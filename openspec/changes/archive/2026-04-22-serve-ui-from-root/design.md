## Context

The claw-device-pairing server uses Echo v5 and currently serves API endpoints (`/pair-device`, `/health`). The UI is a Patternfly-based React application that builds to static files. Currently, developers run the UI separately via a development server (port 5173/5174), requiring CORS middleware and two separate processes.

## Goals / Non-Goals

**Goals:**
- Serve the built UI from the backend's root path `/`
- Serve all static assets (JS, CSS, images) with correct MIME types
- Support client-side routing for the SPA (fallback to index.html for non-API routes)
- Eliminate need for separate UI development server in production

**Non-Goals:**
- Server-side rendering or dynamic HTML generation
- Serving unbundled/development UI files (only production builds)
- Modifying existing API routes or behavior

## Decisions

### Use Echo's Static middleware for file serving
Echo provides built-in static file serving that handles MIME types and efficient file serving. We'll use `e.Static("/", "ui/dist")` to serve files from the UI build directory.

**Alternatives considered:**
- Custom file server: More code to maintain, reinvents Echo's existing functionality
- Separate nginx/caddy frontend: Adds deployment complexity for a simple use case

### Serve UI before registering API routes
Registration order matters in Echo. By serving static files first and using specific paths for API routes, we ensure API endpoints take precedence over static file lookups.

**Alternatives considered:**
- API routes first: Would cause conflicts if UI contains files matching API route names
- Route groups: Unnecessary complexity for current simple routing structure

### Use SPA fallback for client-side routing
For paths that don't match static files or API routes, serve index.html to support React Router client-side routing.

**Implementation**: Use `e.File("/*", "ui/dist/index.html")` as a catch-all after API routes.

## Risks / Trade-offs

**[Risk]** UI build directory might not exist → **Mitigation**: Server should log clear error if ui/dist is missing at startup

**[Risk]** API route conflicts with UI route names → **Mitigation**: All API routes use explicit prefixes (`/pair-device`, `/health`), unlikely to conflict with typical UI asset names

**[Trade-off]** CORS middleware becomes unnecessary in production (UI served from same origin) but must remain for development when UI runs separately
